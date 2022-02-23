/*

 Simple simulation of the game of life from recent interview.
 Rules are implemented in doStep() in a switch block.

 Each step we analyze only the positions surrounding live dots. All other dots won't become live according to the rules.
 We sum up the effect from all living dots to their neighbouring positions in the temporary map (map[[2]int]Affect).
 Affect is a special structure for counting all effects and storing the current state of the dot at this position.
 After looping over all neighbours of live dots, we analyze the temporary map and apply rules to all affected positions,
 creating the result of the step.

*/

package main
import (
    "fmt"
)

type Affect struct {
    count int
    live bool
}

// entrypoint
// input data defined here
func main() {

    var coordinates = []int{0, 0, 0, 1, 1, 0, 1, 1, -1, -1}

    printDots(coordinates)
    solution(coordinates, 5)

}

// Implements steps loop only
func solution (coordinates []int, steps int) []int {

    // steps loop
    for step := 0; step < steps; step++ {
        coordinates = doStep(coordinates)
        printDots(coordinates)
    }

    return coordinates
}

// Creates new affects map, call addAffects for each live dot, applies game rules
func doStep (coordinates []int) (result []int) {

    var affects map[[2]int]Affect
    affects = make(map[[2]int]Affect)

    // add 1 affect for each cell around currently live dots
    for i := 0; i < len(coordinates) - 1; i+=2 {
        addAffects(coordinates[i], coordinates[i + 1], affects)
    }

    // apply rule to all affected dots
    for coord, affect := range affects {
        switch {
            case affect.count < 3 || affect.count > 5:
                // if it was alive it dies, if it was dead, it doesn't born again, so do nothing
            case affect.count == 3 && !affect.live:
                // borns
                result = append(result, coord[0], coord[1])
            case (affect.count == 4 || affect.count == 5) && affect.live:
                // continues to live
                result = append(result, coord[0], coord[1])
        }
    }

    return
}

// Top up effects counter for surrounding dots
func addAffects(x, y int, affects map[[2]int]Affect) {
    for xd := -1; xd <= 1; xd++ {
        for yd := -1; yd <= 1; yd++ {
            if xd == 0 && yd == 0 {
                continue;
            }
            coord := [2]int{x + xd, y + yd}
            if affect, ok := affects[coord]; ok {
                affect.count++
                if xd == 0 && yd == 0 {
                    affect.live = true
                }
                affects[coord] = affect
            } else {
                affects[coord] = Affect{1, xd == 0 && yd == 0}
            }
        }
    }
}

// Useful visualization for the small 11x11 square
func printDots(res []int) {
    var size = 11
    var dots []bool

    half := size >> 1
    dots = make([]bool, size * size)
    for i := 0; i < len(res) - 1; i += 2 {
        y := res[i + 1] + half
        x := res[i] + half
        if x < 0 || y < 0 || x >= size || y >= size {
            continue
        }
        dots[y * size + x] = true
    }

    for y := 0; y < size; y++ {
        for x := 0; x < size; x++ {
            if dots[y * size + x] {
                fmt.Print("#")
            } else {
                fmt.Print(".")
            }
        }
        fmt.Print("\r\n")
    }

    fmt.Print("\r\n\r\n")
}
