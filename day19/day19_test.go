package day19

import (
	aoc "aoc"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

/*
Part 2 bruteforce won't work well since it's 4000^4=2.56e+14 combinations without filtering
Notes and thought process for debugging why rewrite was working for p1 but not p2

x: [0,4000] m: [0,4000] a: [0,4000] s: [0,4000]

// in{s<1351:px,qqz}
px -> x: [0,4000] m: [0,4000] a: [0,4000] s: [0,1350]
qqz -> x: [0,4000] m: [0,4000] a: [0,4000] s: [1351,4000]

// px{a<2006:qkq,m>2090:A,rfg}
qkq -> x: [0,4000] m: [0,4000] a: [0,2005] s: [0,1350]
A -> x: [0,4000] m: [2091,4000] a: [2006,4000] s: [0,1350] // ACCEPTED
rfg ->  x: [0,4000] m: [0,2090] a: [2006,4000] s: [0,1350]

// qqz{s>2770:qs,m<1801:hdj,R}
qs -> x: [0,4000] m: [0,4000] a: [0,4000] s: [2771,4000]
hdj -> x: [0,4000] m: [0,1800] a: [0,4000] s: [1351,2770]
R -> x: [0,4000] m: [1801,4000] a: [0,4000] s: [1351,2770] // REJECTED

// qkq{x<1416:A,crn}
A -> x: [0,1415] m: [0,4000] a: [0,2005] s: [0,1350] // ACCEPTED
crn -> x: [1416,4000] m: [0,4000] a: [0,2005] s: [0,1350]

// crn{x>2662:A,R}
A -> x: [2663,4000] m: [0,4000] a: [0,2005] s: [0,1350] // ACCEPTED
R -> x: [1416,2662] m: [0,4000] a: [0,2005] s: [0,1350] // REJECTED

// rfg{s<537:gd,x>2440:R,A}
gd -> x: [0,4000] m: [0,2090] a: [2006,4000] s: [0,536] // REJECTED (gd is all fails)
R -> x: [2441,4000] m: [0,2090] a: [2006,4000] s: [537,1350] // REJECTED
A -> x: [0,2440] m: [0,2090] a: [2006,4000] s: [537,1350] // ACCEPTED

// gd{a>3333:R,R}
R x: [0,4000] m: [0,2090] a: [2006,3333] s: [0,536] // REJECTED
R x: [0,4000] m: [0,2090] a: [3334,4000] s: [0,536] // REJECTED

qs{s>3448:A,lnx}
A -> x: [0,4000] m: [0,4000] a: [0,4000] s: [3449,4000] // ACCEPTED
lnx -> x: [0,4000] m: [0,4000] a: [0,4000] s: [2771,3448]

hdj{m>838:A,pv}
A -> x: [0,4000] m: [839,1800] a: [0,4000] s: [1351,2770] // ACCEPTED
pv -> x: [0,4000] m: [0,838] a: [0,4000] s: [1351,2770]

lnx{m>1548:A,A}
A -> x: [0,4000] m: [1549,4000] a: [0,4000] s: [2771,3448] // ACCEPTED
A -> x: [0,4000] m: [0,1548] a: [0,4000] s: [2771,3448] // ACCEPTED

pv{a>1716:R,A}
R -> x: [0,4000] m: [0,838] a: [1717,4000] s: [1351,2770] // REJECTED
A -> x: [0,4000] m: [0,838] a: [0,1716] s: [1351,2770] // ACCEPTED

A
A -> x: [0,1415] m: [0,4000] a: [0,2005] s: [0,1350] // ACCEPTED
A -> x: [2663,4000] m: [0,4000] a: [0,2005] s: [0,1350] // ACCEPTED
A -> x: [0,4000] m: [2091,4000] a: [2006,4000] s: [0,1350] // ACCEPTED
A -> x: [0,2440] m: [0,2090] a: [2006,4000] s: [537,1350] // ACCEPTED
A -> x: [0,4000] m: [0,4000] a: [0,4000] s: [3449,4000] // ACCEPTED
A -> x: [0,4000] m: [1549,4000] a: [0,4000] s: [2771,3448] // ACCEPTED
A -> x: [0,4000] m: [0,1548] a: [0,4000] s: [2771,3448] // ACCEPTED
A -> x: [0,4000] m: [839,1800] a: [0,4000] s: [1351,2770] // ACCEPTED
A -> x: [0,4000] m: [0,838] a: [0,1716] s: [1351,2770] // ACCEPTED

R
x: [1416,2662] m: [0,4000] a: [0,2005] s: [0,1350] // x<=2662 F
x: [2441,4000] m: [0,2090] a: [2006,4000] s: [537,1350] // x>2440
x: [0,4000] m: [1801,4000] a: [0,4000] s: [1351,2770] // s<=2770 m>=1801
x: [0,4000] m: [0,838] a: [1717,4000] s: [1351,2770] // a>1716
R x: [0,4000] m: [0,2090] a: [2006,3333] s: [0,536]
R x: [0,4000] m: [0,2090] a: [3334,4000] s: [0,536]

sum += rangeX * rangeM * rangeA * rangeS

Output:  // The solved parts match ranges.... the math doesn't math up
Solved p {x:{min:0 max:1415} m:{min:0 max:4000} a:{min:0 max:2005} s:{min:0 max:1350} ruleIdx:0 state:A}
Solved p {x:{min:2663 max:4000} m:{min:0 max:4000} a:{min:0 max:2005} s:{min:0 max:1350} ruleIdx:0 state:A}
Solved p {x:{min:0 max:4000} m:{min:2091 max:4000} a:{min:2006 max:4000} s:{min:0 max:1350} ruleIdx:0 state:A}
Solved p {x:{min:0 max:2440} m:{min:0 max:2090} a:{min:2006 max:4000} s:{min:537 max:1350} ruleIdx:0 state:A}
Solved p {x:{min:0 max:4000} m:{min:0 max:4000} a:{min:0 max:4000} s:{min:3449 max:4000} ruleIdx:0 state:A}
Solved p {x:{min:0 max:4000} m:{min:1549 max:4000} a:{min:0 max:4000} s:{min:2771 max:3448} ruleIdx:0 state:A}
Solved p {x:{min:0 max:4000} m:{min:0 max:1548} a:{min:0 max:4000} s:{min:2771 max:3448} ruleIdx:0 state:A}
Solved p {x:{min:0 max:4000} m:{min:839 max:1800} a:{min:0 max:4000} s:{min:1351 max:2770} ruleIdx:0 state:A}
Solved p {x:{min:0 max:4000} m:{min:0 max:838} a:{min:0 max:1716} s:{min:1351 max:2770} ruleIdx:0 state:A}
Res: 167578630387434


.... my ranges were off..  Was 0-4000, instead of 1-4000...
*/

var example = `px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}`

func TestPart1Example(t *testing.T) {
	output := Process(strings.Split(example, "\n"), 1)
	assert.Equal(t, 19114, output)
}

func TestPart1(t *testing.T) {
	input, _ := aoc.GetInputData(19)
	output := Process(input, 1)

	assert.Equal(t, 446517, output)
}

func TestPart2Example(t *testing.T) {
	output := Process(strings.Split(example, "\n"), 2)
	assert.Equal(t, 167409079868000, output)
}
func TestPart2(t *testing.T) {
	input, _ := aoc.GetInputData(19)
	output := Process(input, 2)

	assert.Equal(t, 130090458884662, output)
}
