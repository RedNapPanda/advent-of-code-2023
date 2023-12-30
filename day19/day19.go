package day19

import (
	"fmt"
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
	var success, failure, ii intRange
	var sb, fb bool
	if i.min <= v {
		ii = intRange{i.min, min(i.max, v)}
		if ge {
			failure = ii
			fb = true
		} else {
			success = ii
			sb = true
		}
	}
	if i.max >= v {
		ii = intRange{max(i.min, v), i.max}
		if ge {
			success = ii
			sb = true
		} else {
			failure = ii
			fb = true
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

	for _, p := range parts {
		_ = f(&workflows, []part{p}, &accepted, 0)
	}

	// fmt.Printf("%+v\n", parts)
	// fmt.Printf("%+v\n", accepted)

	sum := 0
	if version == 1 {
		for _, p := range accepted {
			// fmt.Printf("Accepted p %+v\n", p)
			sum += p.x.min + p.m.min + p.a.min + p.s.min
		}
	} else {

	}

	return sum
}

func f(workflows *map[string]workflow, ps []part, acc *[]part, sum int) int {
	if len(ps) == 0 {
		return 0
	}
	p := ps[0]
	ps = ps[1:]
	if p.state == "R" {
		return 0
	}
	if p.state == "A" {
		// fmt.Printf("Solved p %+v\n", p)
		*acc = append(*acc, p)
		return (p.x.max - p.x.min) *
			(p.m.max - p.m.min) *
			(p.a.max - p.a.min) *
			(p.s.max - p.s.min)
	}
	wk := (*workflows)[p.state]
	r := wk.rules[p.ruleIdx]
	v := p.value(r.c)
	success, sb, fail, fb := v.split(r.val, r.ge)
	if sb {
		ps = append(ps, buildPart(r.c, p, success, 0, r.res))
	}
	if fb {
		if p.ruleIdx+1 >= len(wk.rules) {
			if wk.fail == "R" {
				fmt.Printf("wk.fail p %+v, r %c %d %t\n", p, r.c, r.val, r.ge)
			}
			ps = append(ps, buildPart(r.c, p, fail, 0, wk.fail))
		} else {
			if p.state == "R" {
				fmt.Printf("p.state p %+v, r %c %d %t\n", p, r.c, r.val, r.ge)
			}
			ps = append(ps, buildPart(r.c, p, fail, p.ruleIdx+1, p.state))
		}
	}
	return sum + f(workflows, ps, acc, sum)
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
