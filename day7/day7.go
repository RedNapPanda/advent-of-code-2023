package day6

import (
	"sort"
	"strconv"
	"strings"
)

var cardMap = map[byte]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func getCard(
	char byte,
	joker bool,
) int {
	if joker && char == 'J' {
		return -1
	}
	return cardMap[char]
}

type hand struct {
	handStr string
	cards   map[byte]int
	bet     int
}

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func (h *hand) handType() int {
	three := false
	pairCount := 0
	for _, v := range h.cards {
		switch v {
		case 5:
			return FiveOfAKind
		case 4:
			return FourOfAKind
		case 3:
			three = true
		case 2:
			pairCount++
		}
	}
	if three {
		if pairCount > 0 {
			return FullHouse
		}
		return ThreeOfAKind
	}
	switch pairCount {
	case 2:
		return TwoPair
	case 1:
		return OnePair
	default:
		return HighCard
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

func sortHands(
	hands []hand,
	withJoker bool,
) {
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
