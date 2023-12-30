package day19

import (
	"strconv"
	"strings"
)

type workflow struct {
	rules []rule
	fail  string
}

type part struct {
	x, m, a, s intRange
	ruleIdx    int
	state      string
}

type intRange struct {
	min, max int
}

func (i intRange) split(v int, ge bool) (intRange, bool, intRange, bool) {
	var success, failure intRange
	var sb, fb bool
	if ge {
		if i.min <= v {
			failure = intRange{i.min, min(i.max, v)}
			fb = true
		}
		if i.max > v {
			success = intRange{max(i.min, v+1), i.max}
			sb = true
		}
	} else {
		if i.max >= v {
			failure = intRange{max(i.min, v), i.max}
			fb = true
		}
		if i.min < v {
			success = intRange{i.min, min(i.max, v-1)}
			sb = true
		}
	}
	return success, sb, failure, fb
}

func (p part) value(c byte) intRange {
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
	return intRange{}
}

type rule struct {
	ge  bool
	c   byte
	val int
	res string
}

func Process(lines []string, version int) int {
	parts, workflows := parse(lines)
	var accepted []part

	sum := 0
	if version == 1 {
		for _, p := range parts {
			_ = f(&workflows, p, &accepted)
		}
	} else {
		sum = f(&workflows, part{intRange{1, 4000}, intRange{1, 4000}, intRange{1, 4000}, intRange{1, 4000}, 0, "in"}, &accepted)
	}

	if version == 1 {
		for _, p := range accepted {
			sum += p.x.min + p.m.min + p.a.min + p.s.min
		}
	}
	return sum
}

func f(workflows *map[string]workflow, p part, acc *[]part) int {
	sum := 0
	wk := (*workflows)[p.state]
	if p.state == "R" {
		return 0
	}
	if p.state == "A" {
		*acc = append(*acc, p)
		return (p.x.max - p.x.min + 1) *
			(p.m.max - p.m.min + 1) *
			(p.a.max - p.a.min + 1) *
			(p.s.max - p.s.min + 1)
	}
	r := wk.rules[p.ruleIdx]
	v := p.value(r.c)
	success, sb, fail, fb := v.split(r.val, r.ge)
	if sb {
		sum += f(workflows, buildPart(r.c, p, success, 0, r.res), acc)
	}
	if fb {
		if p.ruleIdx+1 >= len(wk.rules) {
			sum += f(workflows, buildPart(r.c, p, fail, 0, wk.fail), acc)
		} else {
			sum += f(workflows, buildPart(r.c, p, fail, p.ruleIdx+1, p.state), acc)
		}
	}
	return sum
}

func buildPart(c byte, old part, new intRange, ruleIdx int, state string) part {
	switch c {
	case 'x':
		return part{new, old.m, old.a, old.s, ruleIdx, state}
	case 'm':
		return part{old.x, new, old.a, old.s, ruleIdx, state}
	case 'a':
		return part{old.x, old.m, new, old.s, ruleIdx, state}
	case 's':
		return part{old.x, old.m, old.a, new, ruleIdx, state}
	}
	return part{}
}

func parse(lines []string) ([]part, map[string]workflow) {
	v := make(map[string]workflow)
	var parts []part

	toInt := func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	}

	toIntRange := func(s string) intRange {
		i, _ := strconv.Atoi(s)
		return intRange{i, i}
	}

	for _, l := range lines {
		if l == "" {
			continue
		}
		if l[0] == '{' {
			s := strings.Split(l[1:len(l)-1], ",")
			parts = append(parts, part{
				x:     toIntRange(s[0][2:]),
				m:     toIntRange(s[1][2:]),
				a:     toIntRange(s[2][2:]),
				s:     toIntRange(s[3][2:]),
				state: "in",
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
				r = append(r, rule{ge != '<', c, toInt(split[n][2:colonIdx]), split[n][colonIdx+1:]})
			}
			v[key] = workflow{rules: r, fail: split[len(split)-1]}
		}
	}

	return parts, v
}
