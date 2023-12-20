package day6

import (
	"fmt"
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

type Hand struct {
	handStr string
	cards   map[byte]int
	bet     int
}

/*
Ranks
5 of a kind:    6
4 of a kind:    5
Full house:     4
3 of a kind:    3
2 pair:         2
1 pair:         1
high card:      0
*/
func (h *Hand) rank() int {
	three := false
	pairCount := 0
	for _, v := range h.cards {
		switch v {
		case 5:
			return 6
		case 4:
			return 5
		case 3:
			three = true
		case 2:
			pairCount++
		}
	}
	if three {
		return pairCount + 3
	}
	return pairCount
}

func Part1(lines []string) int {
	rankMap := make(map[int][]Hand)
	for _, line := range lines {
		handMap := make(map[byte]int)
		split := strings.Split(line, " ")
		bet, _ := strconv.Atoi(split[1])
		for _, b := range []byte(split[0]) {
			if value, ok := handMap[b]; ok {
				handMap[b] = value + 1
			} else {
				handMap[b] = 1
			}
		}
		hand := Hand{split[0], handMap, bet}
		rank := hand.rank()

		if ranks, ok := rankMap[rank]; ok {
			rankMap[rank] = append(ranks, hand)
		} else {
			rankMap[rank] = []Hand{hand}
		}
	}

	sum := 0
	rank := 1
	for i := 0; i < 7; i++ {
		if _, ok := rankMap[i]; !ok {
			continue
		}
		hands := rankMap[i]
		sort.Slice(hands, func(i, j int) bool {
			for idx, iChar := range []byte(hands[i].handStr) {
				jChar := hands[j].handStr[idx]
				if iChar == jChar {
					continue
				}
				return cardMap[iChar] < cardMap[jChar]
			}
			return true
		})
		for j := 0; j < len(hands); j++ {

		}
		for _, hand := range hands {
			sum += rank * hand.bet
			rank++
		}
	}
	return sum
}

func Part2(lines []string) int {
	rankMap := make(map[int][]Hand)
	for _, line := range lines {
		handMap := make(map[byte]int)
		split := strings.Split(line, " ")
		bet, _ := strconv.Atoi(split[1])
		for _, b := range []byte(split[0]) {
			if value, ok := handMap[b]; ok {
				handMap[b] = value + 1
			} else {
				handMap[b] = 1
			}
		}
		hand := Hand{split[0], handMap, bet}
		rank := hand.rank()
		// fmt.Printf("cards: rank(%d) | %+v\n", rank, hand)

		if ranks, ok := rankMap[rank]; ok {
			rankMap[rank] = append(ranks, hand)
		} else {
			rankMap[rank] = []Hand{hand}
		}
	}

	sum := 0
	rank := 1
	for i := 0; i < 7; i++ {
		if _, ok := rankMap[i]; !ok {
			continue
		}
		hands := rankMap[i]
		fmt.Printf("rank(%d) hands: %d\n", i, len(hands))
		sort.Slice(hands, func(i, j int) bool {
			for idx, iChar := range []byte(hands[i].handStr) {
				jChar := hands[j].handStr[idx]
				if iChar == jChar {
					continue
				}
				return cardMap[iChar] < cardMap[jChar]
			}
			return true
		})
		fmt.Printf("Hands: %+v\n", hands)
		for j := 0; j < len(hands); j++ {

		}
		for _, hand := range hands {
			fmt.Printf("sum %d | value %d | rank %d | hand %+v\n", sum, rank*hand.bet, rank, hand)
			sum += rank * hand.bet
			rank++
		}
	}
	return sum
}
