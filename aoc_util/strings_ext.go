package aoc_util

func Repeat(orig, sep string, times int) string {
	result := orig
	for i := 1; i < times; i++ {
		result += sep + orig
	}
	return result
}
