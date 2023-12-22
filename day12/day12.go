package day12

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	result := 0
	for _, line := range lines {
		unfixedRecord, groups := parseLine(line)
		chars := []byte(unfixedRecord)
		chars = append(chars, '.')
		charLen := len(chars)
		groups = append(groups, 0)
		slices.Reverse(groups)

		l := make([]string, charLen)
		for i, c := range chars {
			l[i] = string(c)
		}
		l = append(l, "_")
		dp := make([][]int, len(groups))
		for i := 0; i < len(dp); i++ {
			dp[i] = make([]int, charLen)
		}
		for i := charLen - 1; i >= 0 && chars[i] != '#'; i-- {
			dp[0][i] = 1
		}

		broken := func(slice []byte, x, y, group int) int {
			if match(slice, group) {
				return dp[x-1][min(y+group+1, charLen-1)]
			}
			return 0
		}

		for x := 1; x < len(groups); x++ {
			group := groups[x]
			for y := charLen - 2; y >= 0; y-- {
				slice := chars[y:]
				switch chars[y] {
				case '.':
					dp[x][y] = dp[x][y+1]
				case '#':
					dp[x][y] = broken(slice, x, y, group)
				case '?':
					value := broken(slice, x, y, group)
					dp[x][y] = dp[x][y+1] + value
				}
			}
		}
		fmt.Printf("%+v\n", line)
		for _, d := range dp {
			fmt.Printf("%+v\n", d)
		}
		/*
		   			dp[charIndex][j] = how many arrangements when matching chars[charIndex..n-1] with springs[j..m-1]
		   			dp[charIndex][j] = match(chars[charIndex]):
		   				 # => if(s[j] == T) dp[charIndex+1][j+1] (you have to match one damaged spring)
		   				      else 0
		   				 . => if(s[j] == F) dp[charIndex+1][j+1] + dp[charIndex+1][j]  (you can either match the whole gap or match more operational inside the same gap)
		   				      else 0
		   				 ? => you can replace it to . or # so it's the sum of both

		               dyn[i][j] =
		               match string[j] with
		                   | '.' -> dyn[i][j + 1]
		                   | '#' -> if can_match string[j..n] current_group
		                       then dyn[i - 1][j + current_group + 1]
		                       else 0
		                   | '?' -> case '.' + case '#'

		   			[? ? ? . # # # _]: 1,1,3
		   			[0 0 0 0 0 0 0 1]: 0
		   			[0 0 0 0 1 0 0 0]: 3
		   			[0 0 1 0 0 0 0 0]: 1
		   			[1 0 0 0 0 0 0 0]: 1

		   			[. ? ? . . ? ? . . . ? # # . _]: 1,1,3
		   			[0 0 0 0 0 0 0 0 0 0 0 0 0 1 1]: 0
		   			[1 1 1 1 1 1 1 1 1 1 1 0 0 0 0]: 3
		   			[4 4 3 2 2 2 1 0 0 0 0 0 0 0 0]: 1
		   			[4 4 2 0 0 0 0 0 0 0 0 0 0 0 0]: 1

		   			[? # ? # ? # ? # ? # ? # ? # ? _]: 1,3,1,6
		   			[0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 1]: 0
		   			[0 0 0 0 0 0 0 0 2 1 0 0 0 0 0 0]: 6
		   			[0 0 0 0 0 0 1 1 0 0 0 0 0 0 0 0]: 1
		   			[0 0 1 1 0 0 0 0 0 0 0 0 0 0 0 0]: 3
		   			[1 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0]: 1

		*/

		result += dp[len(dp)-1][0]
	}

	return result
}

func match(slice []byte, group int) bool {
	result := len(slice) > group
	for _, c := range slice[:min(len(slice), group)] {
		if !(c == '?' || c == '#') {
			result = false
		}
	}

	return result && (slice[group] == '.' || slice[group] == '?')
}

func parseLine(line string) (string, []int) {
	split := strings.Split(line, " ")
	groupStrs := strings.Split(split[1], ",")
	groups := make([]int, len(groupStrs))
	for i, s := range groupStrs {
		digit, _ := strconv.Atoi(s)
		groups[i] = digit
	}
	return split[0], groups
}

func dfs(original []byte, record []byte, groups []int) int {
	if len(groups) == 0 {
		return 1
	}
	prevResult := 0
	current, nextGroups := groups[0], groups[1:]

	sum := 0
	for _, g := range nextGroups {
		sum += g
	}
	for i := 0; i <= len(record)-sum-len(nextGroups)-current; i++ {
		b := contains(record[:i], '#')
		if b {
			break
		}
		next := i + current
		if next <= len(record) && !contains(record[i:next], '.') && (next >= len(record) || record[next] != '#') {
			var bytes []byte
			if next+1 < len(record) {
				bytes = record[next+1:]
			}
			val := dfs(original, bytes, nextGroups)
			prevResult += val
		}

	}
	return prevResult
}

func contains(bs []byte, b byte) bool {
	for _, b1 := range bs {
		if b1 == b {
			return true
		}
	}
	return false
}
