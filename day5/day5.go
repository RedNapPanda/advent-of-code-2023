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

func (m mapRange) intersect(o mapRange) (bool, mapRange, []mapRange) {
	if o.start >= m.end || m.start >= o.end {
		return false, m, nil
	}
	if m.start >= o.start && m.end <= o.end {
		return true, mapRange{m.start, m.end, o.shift}, nil
	}
	ms := max(m.start, o.start)
	me := min(m.end, o.end)

	var rem []mapRange
	if m.start < o.start {
		rem = append(rem, mapRange{m.start, o.start, 0})
	}

	if m.end > o.end {
		rem = append(rem, mapRange{o.end, m.end, 0})
	}

	return true, mapRange{ms, me, o.shift}, rem
}

type step struct {
	mapRange
	step int
}

func Part1(lines []string) int {
	lowest := math.MaxInt
	mapDatas := parseMappings(lines, 1)
	var steps []step
	for s := 0; s < len(mapDatas[0]); s++ {
		steps = append(steps, step{mapDatas[0][s], 1})
	}
	var completed []mapRange
stepJmp:
	for len(steps) > 0 {
		s := steps[0]
		steps = steps[1:]
		if s.step == len(mapKeys)-1 {
			completed = append(completed, s.mapRange)
			if s.start < lowest {
				lowest = s.start
			}
			continue
		}

		for _, mp := range mapDatas[s.step] {
			b, i, rm := s.mapRange.intersect(mp)
			if !b {
				continue
			}
			if len(rm) != 0 {
				for _, r := range rm {
					steps = append(steps, step{r, s.step})
				}
			}
			i.start += i.shift
			i.end += i.shift
			steps = append(steps, step{i, s.step + 1})
			continue stepJmp
		}
		s.step += 1
		steps = append(steps, s)
	}

	return lowest
}

// TODO: Finish optimizing this. Getting distract with BFS/DFS atm, which isn't related
func Part2(lines []string) int {
	lowest := math.MaxInt
	mapDatas := parseMappings(lines, 2)
	var steps []step
	for s := 0; s < len(mapDatas[0]); s++ {
		steps = append(steps, step{mapDatas[0][s], 1})
	}
	var completed []mapRange
stepJmp:
	for len(steps) > 0 {
		s := steps[0]
		steps = steps[1:]
		if s.step == len(mapKeys) {
			completed = append(completed, s.mapRange)
			if s.start < lowest {
				lowest = s.start
			}
			continue
		}

		for _, mp := range mapDatas[s.step] {
			b, i, rm := s.mapRange.intersect(mp)
			if !b {
				continue
			}
			if len(rm) != 0 {
				for _, r := range rm {
					steps = append(steps, step{r, s.step})
				}
			}
			i.start += i.shift
			i.end += i.shift
			steps = append(steps, step{i, s.step + 1})
			continue stepJmp
		}
		s.step += 1
		steps = append(steps, s)
	}

	return lowest
}

func parseMappings(lines []string, part int) map[int][]mapRange {
	index := 0
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
				for i := 0; i < len(split); i += 2 {
					start, _ := strconv.Atoi(split[i])
					length, _ := strconv.Atoi(split[i+1])
					seeds = append(seeds, mapRange{start: start, end: start + length})
				}
			}
			val[0] = seeds
			continue
		}
		if index < len(mapKeys)-1 && line == mapKeys[index+1] {
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
		mapRanges = append(mapRanges, mapRange{src, src + length, dest - src})
	}
	val[index] = mapRanges

	for i, _ := range val {
		slices.SortFunc(val[i], func(a, b mapRange) int {
			return a.start - b.start
		})
	}

	return val
}
