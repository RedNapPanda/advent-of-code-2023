package day11

import "fmt"

func Process(lines []string) int {
    var matrix [][]byte
    inc, lineLength := 0, 0
    // parse rows and insert extra empty rows
    for x, line := range lines {
        if lineLength == 0 {
            lineLength = len(line)
        }
        matrix = append(matrix, make([]byte, len(line)))
        isEmpty := true
        for y, c := range []byte(line) {
            matrix[x+inc][y] = c
            if c == '#' {
                isEmpty = false
            }
        }
        if isEmpty {
            slice := matrix[x+inc]
            matrix = append(matrix, slice)
            inc++
        }
    }
    var emptyColumns []int
    // find empty columns
    for y := 0; y < len(matrix[0]); y++ {
        isEmpty := true
        for x := 0; x < len(matrix); x++ {
            c := matrix[x][y]
            if c == '#' {
                isEmpty = false
            }
        }
        if isEmpty {
            emptyColumns = append(emptyColumns, y)
        }
    }

    inc = 0
    for _, y := range emptyColumns {
        y = y + inc
        for x := 0; x < len(matrix); x++ {
            pre := matrix[x][:y+1]
            post := matrix[x][y:]
            list := append(pre, post...)
            if y == len(matrix[x]) {
                list = append(matrix[x], '.')
            }
            list[y] = '.'
            matrix[x] = list
        }
        inc++
    }
    fmt.Printf("%d\n", len(matrix))
    fmt.Printf("%+v\n", emptyColumns)

    return 0
}
