package day5

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

var mapKeys = []string{
	"seeds: ",
	"seed-to-soil map:",
	"soil-to-fertilizer map:",
	"fertilizer-to-water map:",
	"water-to-light map:",
	"light-to-temperature map:",
	"temperature-to-humidity map:",
	"humidity-to-location map:",
}

type intRange struct {
	start, end int
}

/*
f(x) = if input intersects src range -> shift intersection and create additional ranges to cover XOR

Use a deque to handle the additional ranges that don't intersect.  They could intersect with another mapping or just fallthrough.
[98,104) AND [98,100) -> [50,52) + [100,104)
[40,68) AND [50,98) -> [40,50) + [52,70)
[79,93) AND [50,98) -> [81,95)
[55,68) AND [50,98) -> [57,70)

[98,100) -> [50,52)
[50,98) -> [52,100)

break at first intersected range
*/
type mapRange struct {
	start, end, shift int
}

func Part1(lines []string) int {
	lowest := math.MaxInt
	val := 0
	mapDatas := parseMappings(lines, 1)
	for _, mapSeed := range mapDatas[0] {
		val = mapSeed.start
		for n := 1; n < len(mapKeys); n++ {
			for _, mr := range mapDatas[n] {
				if mr.shift != -1 && val >= mr.start && val < mr.end {
					val = val + (mr.shift - mr.start)
					break
				}
			}
		}
		if val < lowest {
			lowest = val
		}
	}
	return lowest
}

// TODO: Finish optimizing this. Getting distract with BFS/DFS atm, which isn't related
func Part2(lines []string) int {
	lowest := math.MaxInt
	val := 0
	mapDatas := parseMappings(lines, 2)
	mr := mapDatas[0]
	var intersected []mapRange
	for _, mapSeed := range mapDatas[0] {
		val = mapSeed.start
		for _, key := range mapKeys {
			for _, mr := range mapDatas[key] {
				if val >= mr.start && val < mr.end {
					val = val + (mr.shift - mr.start)
					break
				}
			}
		}
		if val < lowest {
			lowest = val
		}
	}
	return lowest
}

func parseMappings(lines []string, part int) map[int][]mapRange {
	index := 1
	val := make(map[int][]mapRange)
	var mapRanges []mapRange
	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, mapKeys[0]) {
			var seeds []mapRange
			if part == 1 {
				split := strings.Split(line[len(mapKeys[0]):], " ")
				for _, s := range split {
					v, _ := strconv.Atoi(s)
					seeds = append(seeds, mapRange{start: v, end: v})
				}
			} else if part == 2 {
				split := strings.Split(line[len(mapKeys[0]):], " ")
				seeds = make([]mapRange, len(split))
				for i := 0; i < len(split); i += 2 {
					start, _ := strconv.Atoi(split[i])
					end, _ := strconv.Atoi(split[i+1])
					seeds = append(seeds, mapRange{start: start, end: end})
				}
			}
			val[0] = seeds
			continue
		}
		if index < len(mapKeys) && line == mapKeys[index] {
			if len(mapRanges) != 0 {
				val[index] = mapRanges
			}
			index++
			mapRanges = []mapRange{}
			continue
		}
		split := strings.Split(line, " ")
		dest, _ := strconv.Atoi(split[0])
		src, _ := strconv.Atoi(split[1])
		length, _ := strconv.Atoi(split[2])
		mapRanges = append(mapRanges, mapRange{src, src + length, dest})
	}
	val[index] = mapRanges

	for _, v := range val {
		slices.SortFunc(v, func(a, b mapRange) int {
			return a.start - b.start
		})
	}

	return val
}
