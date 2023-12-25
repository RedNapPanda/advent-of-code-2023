package day5

import (
	"math"
	"strconv"
	"strings"
	"sync"
)

var seedPrefix = "seeds: "
var lineKeys = []string{
	"seed-to-soil map:",
	"soil-to-fertilizer map:",
	"fertilizer-to-water map:",
	"water-to-light map:",
	"light-to-temperature map:",
	"temperature-to-humidity map:",
	"humidity-to-location map:",
}

type seedRange struct {
	seed   int
	length int
	count  int
}

func (s *seedRange) next() bool {
	if s.count >= s.length {
		return false
	}
	s.count++
	s.seed++
	return true
}

type mapData struct {
	key      string
	mappings [][]int
}

func Part1(lines []string) int {
	var mapDatas []mapData
	var seeds []int
	index := 0
	lowest := math.MaxInt
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, seedPrefix) {
			seeds = parseSeeds(line)
			continue
		}
		if index < len(lineKeys) && line == lineKeys[index] {
			index++
			mapDatas = append(mapDatas, mapData{})
			continue
		}
		m := parseMapSlice(line)
		mapDatas[index-1].mappings = append(mapDatas[index-1].mappings, m)
	}

	for _, seed := range seeds {
		seed = applyMappings(mapDatas, seed)
		if seed < lowest {
			lowest = seed
		}
	}
	return lowest
}

func Part2(lines []string) int {
	var mapDatas []mapData
	var seedRanges []seedRange
	index := 0
	lowest := math.MaxInt
	var mutex sync.Mutex
	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, seedPrefix) {
			seedRanges = parseSeedRanges(line)
			continue
		}
		if index < len(lineKeys) && line == lineKeys[index] {
			index++
			mapDatas = append(mapDatas, mapData{})
			continue
		}
		m := parseMapSlice(line)
		mapDatas[index-1].mappings = append(mapDatas[index-1].mappings, m)
	}
	wg := sync.WaitGroup{}

	for _, sr := range seedRanges {
		wg.Add(1)
		go func(sr seedRange) {
			defer wg.Done()
			rangeLeast := math.MaxInt
			for ok := true; ok; ok = sr.next() {
				targetSeed := applyMappings(mapDatas, sr.seed)
				if targetSeed < rangeLeast {
					rangeLeast = targetSeed
				}
			}
			mutex.Lock()
			if rangeLeast < lowest {
				lowest = rangeLeast
			}
			mutex.Unlock()
		}(sr)
	}
	wg.Wait()
	return lowest
}

func applyMappings(mapDatas []mapData, seed int) int {
	for _, mapData := range mapDatas {
		for _, mapping := range mapData.mappings {
			v, mapped := mapValue(mapping, seed)
			seed = v
			if mapped {
				break
			}
		}
	}
	return seed
}

func mapValue(mapping []int, value int) (int, bool) {
	if mapping[1] <= value && value < mapping[1]+mapping[2] {
		return value - mapping[1] + mapping[0], true
	}
	return value, false
}

func parseSeeds(line string) []int {
	split := strings.Split(line[len(seedPrefix):], " ")
	var seeds []int
	for i := 0; i < len(split); i++ {
		seed, _ := strconv.Atoi(split[i])
		seeds = append(seeds, seed)
	}
	return seeds
}

func parseSeedRanges(line string) []seedRange {
	split := strings.Split(line[len(seedPrefix):], " ")
	var seeds []seedRange
	for i := 0; i < len(split); i += 2 {
		seed, _ := strconv.Atoi(split[i])
		length, _ := strconv.Atoi(split[i+1])
		seeds = append(seeds, seedRange{seed, length, 0})
	}
	return seeds
}

func parseMapSlice(line string) []int {
	split := strings.Split(line, " ")
	dest, _ := strconv.Atoi(split[0])
	source, _ := strconv.Atoi(split[1])
	length, _ := strconv.Atoi(split[2])

	return []int{dest, source, length}
}
