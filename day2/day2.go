package day2

import (
    "fmt"
    "regexp"
    "strconv"
    "strings"
)

type DiceSet struct {
    red   int
    green int
    blue  int
}

type Game struct {
    id   int
    sets []DiceSet
}

func Part1(lines []string, maxRed int, maxGreen int, maxBlue int) int {
    sum := 0
    for idx, line := range lines {
        game, err := parseGame(line)
        if err != nil {
            fmt.Printf("Error parsing line %d | error: %s | %s\n", idx, err, line)
            continue
        }
        if game.possible(maxRed, maxGreen, maxBlue) {
            sum += game.id
        }
    }
    return sum
}

func Part2(lines []string) int {
    sum := 0
    for idx, line := range lines {
        game, err := parseGame(line)
        if err != nil {
            fmt.Printf("Error parsing line %d | error: %s | %s\n", idx, err, line)
            continue
        }
        red, green, blue := minimumDiceSet(game)
        sum += red * green * blue
    }
    return sum
}

func (game *Game) possible(maxRed int, maxGreen int, maxBlue int) bool {
    for _, set := range game.sets {
        if maxRed < set.red ||
            maxGreen < set.green ||
            maxBlue < set.blue {
            return false
        }
    }
    return true
}

func parseGame(s string) (*Game, error) {
    gameRegex := regexp.MustCompile(`Game (\d+): (.+)`)
    matches := gameRegex.FindAllStringSubmatch(s, 1)

    id, _ := strconv.Atoi(matches[0][1])

    var sets []DiceSet
    setStrings := strings.Split(matches[0][2], ";")
    for _, setStr := range setStrings {
        red, green, blue := parseDiceSet(setStr)
        sets = append(sets, DiceSet{red, green, blue})
    }

    return &Game{
        id:   id,
        sets: sets,
    }, nil
}

func parseDiceSet(s string) (int, int, int) {
    diceStrings := strings.Split(s, ",")
    red, green, blue := 0, 0, 0
    for _, diceStr := range diceStrings {
        splitDice := strings.Split(strings.TrimSpace(diceStr), " ")
        count, _ := strconv.Atoi(splitDice[0])
        trimmed := strings.TrimSpace(splitDice[1])
        if trimmed == "red" {
            red = count
        } else if trimmed == "green" {
            green = count
        } else if trimmed == "blue" {
            blue = count
        }
    }
    return red, green, blue
}

func minimumDiceSet(game *Game) (int, int, int) {
    minRed, minGreen, minBlue := -1, -1, -1
    for _, set := range game.sets {
        minRed = max(minRed, set.red)
        minGreen = max(minGreen, set.green)
        minBlue = max(minBlue, set.blue)
    }
    return minRed, minGreen, minBlue
}
