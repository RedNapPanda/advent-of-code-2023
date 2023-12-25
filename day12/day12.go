package day12

import (
	"fmt"
	"strconv"
	"strings"
)

var recurseCache = make(map[string]int)

/*
RecursiveProcess

	???.### 1,1,3
	0 - group size 1

		'.' -> recurse(charIndex + 1, groupIndex)
		'#' -> if isBrokenGroup recurse(charIndex+groupSize+1, groupIndex+1) else 0
		'?' -> sum '.' + '#'
*/
func RecursiveProcess(lines []string) int {
	result := 0
	for _, line := range lines {
		recurseCache = make(map[string]int)
		unfixedRecord, groups := parseLine(line)
		chars := []byte(unfixedRecord)
		// Closing off record with a working spring to ensure the last group could match at the very end
		chars = append(chars, '.')

		result += recurse(chars, groups, 0, 0)
	}
	return result
}

func recurse(chars []byte, groups []int, charIndex, groupIndex int) int {
	charLen, groupLen := len(chars), len(groups)
	if charIndex == charLen {
		// end of springs, have we matched all groups?
		if groupIndex == groupLen {
			return 1
		}
		return 0
	}
	key := fmt.Sprintf("%d,%d", charIndex, groupIndex)
	if v, ok := recurseCache[key]; ok {
		return v
	}

	handleWorking := func() int {
		return recurse(chars, groups, charIndex+1, groupIndex)
	}
	handleBroken := func() int {
		if groupIndex == groupLen {
			return 0
		}
		endGroupIndex := charIndex + groups[groupIndex]
		if !isBrokenGroup(chars, charIndex, endGroupIndex) {
			return 0
		}
		if endGroupIndex == charLen {
			// end of springs, have we matched all groups?
			if groupIndex == groupLen {
				return 1
			}
			return 0
		}

		return recurse(chars, groups, endGroupIndex+1, groupIndex+1)
	}

	value := 0
	switch chars[charIndex] {
	case '.':
		value = handleWorking()
	case '#':
		value = handleBroken()
	case '?':
		value = handleWorking() + handleBroken()
	}
	recurseCache[key] = value
	return value
}

/*
Process

	Needed help with tabulation for this process.  Turns out its almost identical to the recursive strategy, just reversed in a sense

	Dpamic programming. f (pos, groups, len) = number of ways to:
	   parse the first pos positions
	   have groups groups of #
	   with the last group of # having length len

The transitions are:

	if the character is # or ?, we can continue the previous group or start a new group:
	f (pos + 1, groups + (len == 0), len + 1) += f (pos, groups, len)
	if the character is . or ?, and the length of the current group is zero or exactly what we need, we can proceed without a group:
	f (pos + 1, groups, 0) += f (pos, groups, len)

In the end, the answer is f (lastPos, numberOfGroups, 0). (Add a trailing . to the string for convenience to avoid cases.)

dp[y][charIndex]
dp[y][charIndex] =

	dp[x][y] =
	match string[y] with
	    | '.' -> dp[x][y + 1]
	    | '#' -> if can_match string[y : n] current_group
	        then dp[x - 1][y + current_group + 1]
	        else 0
	    | '?' -> case '.' + case '#'

x => group index
y => string index
append a '.' to the end of the record, ensures we have at least 1 working spring for the record
for every group x := groupLen-1;x >= 0; x--
for 2nd to last char (skip the added one) y := len(chars)-2
'.' -> we take the value of the string[r:] in the same row => dp[x][y + 1] - this is a working spring, therefore can't match against a group
'#' ->
'?' -> union of '.' and '#' since it could be either

backtrack
[? ? ? . # # # .]: 1,1,3
. -> only this is 1 since the next is broken
[0 0 0 0 0 0 0 1]: 0
#, ## wrong len
### -> 3
[0 0 0 0 1 0 0 0]: 3
?.### -> 1
[0 0 1 0 0 0 0 0]: 1,3
??.###
???.### -> 1
[1 0 0 0 0 0 0 0]: 1,1,3

[. ? ? . . ? ? . . . ? # # . _]: 1,1,3
[0 0 0 0 0 0 0 0 0 0 0 0 0 1 1]: 0
[1 1 1 1 1 1 1 1 1 1 1 0 0 0 0]: 1
[4 4 3 2 2 2 1 0 0 0 0 0 0 0 0]: 1,3
[4 4 2 0 0 0 0 0 0 0 0 0 0 0 0]: 1,1,3

[? # ? # ? # ? # ? # ? # ? # ? _]: 1,3,1,6
[0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 1]: 0
[0 0 0 0 0 0 0 0 2 1 0 0 0 0 0 0]: 6
[0 0 0 0 0 0 1 1 0 0 0 0 0 0 0 0]: 1,6
[0 0 1 1 0 0 0 0 0 0 0 0 0 0 0 0]: 3,1,6
[1 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0]: 1,3,1,6
*/
func Process(lines []string) int {
	var result = int(0)
	for _, line := range lines {
		unfixedRecord, groups := parseLine(line)
		chars := []byte(unfixedRecord)
		/*
			appending a working spring to ensure if the end of the record is a broken spring group, it gets closed off
			1,1,3 -> 1,1,3,0
			both match => ???.### -> #.#.###. (closing it off ensures that a group size of zero broken springs matches after
			base case group size 0
			ex: = # -> #. matches group 1
			without it, # doesn't match
		*/
		chars = append(chars, '.')
		groupLen := len(groups)
		charLen := len(chars)
		/*
			Instead of an entire dp matrix of [charLen, groupLen+1]
				The extra group is broken spring group of 0 to match with the appended working spring
				Does the same as filling previous RTL with 1s till '#'

			Since you only ever need either data from:
				current groupState: dp[y+1][y]
				prev groupState (next group in the group listing): dp[y - 1][y + current_group + 1]
		*/
		prev := make([]int, charLen)
		current := make([]int, charLen)
		/*
			Find the last broken spring.
			Anything beyond this for initial previous is a 1 (1 possible combination of 0 broken springs)
		*/
		for lastBrokenSpring := charLen - 1; lastBrokenSpring >= 0 && chars[lastBrokenSpring] != '#'; lastBrokenSpring-- {
			prev[lastBrokenSpring] = 1
		}

		/*
			handleBroken := func() int {
				if groupIndex == groupLen {
					return 0
				}
				endGroupIndex := charIndex + groups[groupIndex]
				if !isBrokenGroup(chars, charIndex, endGroupIndex) {
					return 0
				}
				if endGroupIndex == charLen {
					// end of springs, have we matched all groups?
					if groupIndex == groupLen {
						return 1
					}
					return 0
				}
				return recurse(chars, groups, endGroupIndex+1, groupIndex+1)
			}
		*/
		handleBroken := func(x, y int) int {
			if x == groupLen {
				return 0
			}
			brokenGroup := groups[x]
			endGroupIndex := y + brokenGroup
			if !isBrokenGroup(chars, y, endGroupIndex) {
				return 0
			}
			if endGroupIndex == charLen {
				// end of springs, have we matched all groups
				if x != groupLen {
					return 0
				}
			}
			return prev[min(y+brokenGroup+1, charLen-1)]
		}

		// handleBroken used for checking broken spring group matches: ['#', '?'] prefix substring chars[y:], group size to check
		// handleBroken := func(slice []byte, y, y, brokenGroup int) int {
		// 	valid := len(slice) > brokenGroup
		// 	// check slice[:brokenGroup] does not contain a working spring '.'
		// 	for y := 0; y < min(len(slice), brokenGroup); y++ {
		// 		if slice[y] == '.' {
		// 			valid = false
		// 			break
		// 		}
		// 	}
		// 	/*
		// 		brokenGroup = 3
		// 		### -> true
		// 		###. -> true
		// 		.### -> false
		// 		.###. -> false
		//
		// 		if char len > brokenGroup -> cannot match
		// 		does the beginning brokenGroup slice all possibly broken springs
		// 		next spring after broken group is not a broken spring [group len would then be brokenGroup + n]
		// 		  (safe via short circuit of len check earlier)
		// 	*/
		// 	if valid && slice[brokenGroup] != '#' {
		// 		return prev[min(y+brokenGroup+1, charLen-1)]
		// 	}
		// 	return 0
		// }

		for y := groupLen - 1; y >= 0; y-- {
			for x := charLen - 2; x >= 0; x-- { // compute matches for groups from end to beginning
				switch chars[x] {
				case '.':
					current[x] = current[x+1] // working so is the value of the f(substring chars[x+1:])
				case '#':
					current[x] = handleBroken(y, x) // see @handleBroken
				case '?':
					current[x] = current[x+1] + handleBroken(y, x) // working or broken, so sum of both states
				}
			}
			prev = current
			current = make([]int, charLen)
		}
		result += prev[0]
	}
	return result
}

func isBrokenGroup(slice []byte, start, end int) bool {
	sliceLen := len(slice)
	if end > sliceLen {
		return false
	}
	for x := start; x < end; x++ {
		if slice[x] == '.' {
			return false
		}
	}
	return sliceLen > end && slice[end] != '#'
}

func parseLine(line string) (string, []int) {
	split := strings.Split(line, " ")
	groupStrs := strings.Split(split[1], ",")
	groups := make([]int, len(groupStrs))
	for x, s := range groupStrs {
		digit, _ := strconv.Atoi(s)
		groups[x] = digit
	}
	return split[0], groups
}
