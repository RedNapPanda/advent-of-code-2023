package day6

import (
	"sort"
	"strconv"
	"strings"
)

var cardMap = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

func getCard(char byte, joker bool) int {
	if joker && char == 'J' {
		return 1
	}
	return cardMap[char]
}

type hand struct {
	handStr string
	cards   map[byte]int
	bet     int
}

const FIVE_OF_A_KIND = 6
const FOUR_OF_A_KIND = 5
const FULL_HOUSE = 4
const THREE_OF_A_KIND = 3
const TWO_PAIR = 2
const ONE_PAIR = 1
const HIGH_CARD = 0

func (h *hand) handType() int {
	three := false
	pairCount := 0
	for _, v := range h.cards {
		switch v {
		case 5:
			return FIVE_OF_A_KIND
		case 4:
			return FOUR_OF_A_KIND
		case 3:
			three = true
		case 2:
			pairCount++
		}
	}
	if three {
		if pairCount > 0 {
			return FULL_HOUSE
		}
		return THREE_OF_A_KIND
	}
	switch pairCount {
	case 2:
		return TWO_PAIR
	case 1:
		return ONE_PAIR
	default:
		return HIGH_CARD
	}
}

func Part1(lines []string) int {
	var hands []hand
	for _, line := range lines {
		split := strings.Split(line, " ")
		bet, _ := strconv.Atoi(split[1])
		handMap := make(map[byte]int)
		for _, b := range []byte(split[0]) {
			if _, ok := handMap[b]; ok {
				handMap[b]++
			} else {
				handMap[b] = 1
			}
		}
		hands = append(hands, hand{split[0], handMap, bet})
	}

	sortHands(hands, false)
	return calcWinnings(hands)
}

func Part2(lines []string) int {
	var hands []hand
	for _, line := range lines {
		split := strings.Split(line, " ")
		bet, _ := strconv.Atoi(split[1])
		jokers := strings.Count(split[0], "J")
		key, highest := byte(0), 0
		handMap := make(map[byte]int)
		for _, b := range []byte(split[0]) {
			if b == 'J' {
				continue
			}
			if _, ok := handMap[b]; ok {
				handMap[b]++
			} else {
				handMap[b] = 1
			}
			if highest <= handMap[b] {
				highest = handMap[b]
				key = b
			}
		}
		delete(handMap, key)
		handMap[key] = highest + jokers
		hands = append(hands, hand{split[0], handMap, bet})
	}

	sortHands(hands, true)
	return calcWinnings(hands)
}

func sortHands(hands []hand, withJoker bool) {
	sort.SliceStable(hands, func(i, j int) bool {
		hand1, hand2 := hands[i], hands[j]
		type1, type2 := hand1.handType(), hand2.handType()
		if type1 != type2 {
			return type1 < type2
		} else {
			for idx, char1 := range []byte(hand1.handStr) {
				char2 := hand2.handStr[idx]
				if char1 == char2 {
					continue
				}
				return getCard(char1, withJoker) < getCard(char2, withJoker)
			}
			return true
		}
	})
}

func calcWinnings(hands []hand) int {
	sum := 0
	for i := 1; i <= len(hands); i++ {
		sum += i * hands[i-1].bet
	}
	return sum
}
