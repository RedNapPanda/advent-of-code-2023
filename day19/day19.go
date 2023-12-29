package day19

import (
	"fmt"
	"strconv"
	"strings"
)

type part struct {
	x, m, a, s int
	state      byte
}

func (p part) value(c byte) int {
	switch c {
	case 'x':
		return p.x
	case 'm':
		return p.m
	case 'a':
		return p.a
	case 's':
		return p.s
	}
	return -1
}

type workflow struct {
	rules []rule
	res   string
}

func (w workflow) apply(p part) string {
	for _, r := range w.rules {
		v := p.value(r.c)
		if v == -1 {
			continue
		}
		v -= r.val
		if r.neg && v < 0 || !r.neg && v > 0 {
			return r.res
		}
	}
	return w.res
}

type rule struct {
	neg bool
	c   byte
	val int
	res string
}

func Process(lines []string, version int) int {
	parts, workflows := parse(lines)
	var accepted []part

	for _, p := range parts {
		w := workflows["in"]
		for p.state == 0 {
			res := w.apply(p)
			if res == "R" {
				break
			}
			if res == "A" {
				accepted = append(accepted, p)
				break
			}
			w = workflows[res]
		}
	}

	fmt.Printf("%+v\n", parts)
	fmt.Printf("%+v\n", accepted)

	sum := 0
	for _, p := range accepted {
		sum += p.x + p.m + p.a + p.s
	}

	return sum
}

func parse(lines []string) ([]part, map[string]workflow) {
	v := make(map[string]workflow)
	var parts []part

	toInt := func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	}

	for _, l := range lines {
		if l == "" {
			continue
		}
		if l[0] == '{' {
			s := strings.Split(l[1:len(l)-1], ",")
			parts = append(parts, part{
				x: toInt(s[0][2:]),
				m: toInt(s[1][2:]),
				a: toInt(s[2][2:]),
				s: toInt(s[3][2:]),
			})
		} else {
			fBrace := strings.Index(l, "{")
			key := l[:fBrace]
			split := strings.Split(l[fBrace+1:len(l)-1], ",")
			var r []rule
			for n := 0; n < len(split)-1; n++ {
				c := split[n][0]
				ge := split[n][1]
				colonIdx := strings.Index(split[n], ":")
				neg := false
				if ge == '<' {
					neg = true
				}
				r = append(r, rule{neg, c, toInt(split[n][2:colonIdx]), split[n][colonIdx+1:]})
			}
			v[key] = workflow{rules: r, res: split[len(split)-1]}
		}
	}

	return parts, v
}
