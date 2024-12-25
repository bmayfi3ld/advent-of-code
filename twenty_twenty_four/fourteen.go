package twentytwentyfour

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/bmayfi3ld/advent-of-code/pkg/cmd"
	wrapper "github.com/bmayfi3ld/advent-of-code/pkg/timer"
)

func init() {
	cmd.RegisterCommand("2024-14-1-test", func() error { return fourteenOne(fourteenTestInput, 11, 7) })
	cmd.RegisterCommand("2024-14-1-real", func() error { return fourteenOne(fourteenRealInput, 101, 103) })
	cmd.RegisterCommand("2024-14-2-real", func() error { return fourteenTwo(fourteenRealInput, 101, 103) })
	// cmd.RegisterCommand("2024-t-1-template", func() error { return fourteenOne(sixInput) })
}

func fourteenOne(input string, xSize, ySize int) error {
	fmt.Println("hello")
	defer wrapper.ProfileFunction("fourteenOne")()

	re := regexp.MustCompile(`p=(\d{1,3}),(\d{1,3}) v=(-?\d{1,3}),(-?\d{1,3})`)

	matches := re.FindAllStringSubmatch(input, -1)

	finalPositions := [][]int{}

	for _, match := range matches {
		vals := parseValues(match)

		finalX, finalY := runMovements(vals, 100, xSize, ySize)

		finalPositions = append(finalPositions, []int{finalX, finalY})
	}

	// spew.Dump(finalPositions)

	quadrants := getQuadrants(finalPositions, xSize, ySize)

	// spew.Dump(quadrants)

	fmt.Println(quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3])

	return nil
}

func fourteenTwo(input string, xSize, ySize int) error {
	fmt.Println("hello")
	defer wrapper.ProfileFunction("fourteenOne")()

	re := regexp.MustCompile(`p=(\d{1,3}),(\d{1,3}) v=(-?\d{1,3}),(-?\d{1,3})`)

	matches := re.FindAllStringSubmatch(input, -1)

	for simTime := range 10000 {
		// 3397 too low
		if simTime < 3000 {
			continue
		}
		finalPositions := [][]int{}
		for _, match := range matches {
			vals := parseValues(match)

			finalX, finalY := runMovements(vals, simTime, xSize, ySize)

			finalPositions = append(finalPositions, []int{finalX, finalY})
		}

		grid := expandPositions(finalPositions, xSize, ySize)

		potentialTree := checkForPotentialTrees(grid)

		if potentialTree {

			fmt.Println("time " + strconv.Itoa(simTime))
			printPositions(finalPositions, xSize, ySize)
		}

	}

	return nil
}

func checkForPotentialTrees(grid [][]bool) bool {
	for r, row := range grid {
		for c, spot := range row {
			if spot {
				// check for top of an arrow
				if r+3 >= len(grid) || c+3 >= len(row) || c-3 < 0 {
					continue
				}
				if grid[r+1][c+1] &&
					grid[r+1][c-1] &&
					grid[r+2][c+2] &&
					grid[r+2][c-2] &&
					grid[r+3][c-3] &&
					grid[r+3][c+3] {
					return true
				}
			}
		}
	}

	return false
}

// take the xy list of robots and their positions and expand it to a full 2d
// grid with true if there is a robot in a spot
func expandPositions(poses [][]int, xSize, ySize int) [][]bool {
	expanded := make([][]bool, ySize+2)

	for i := range expanded {
		expanded[i] = make([]bool, xSize+2)
	}

	for _, pos := range poses {
		expanded[pos[0]][pos[1]] = true
	}

	return expanded
}

func printPositions(poses [][]int, xSize, ySize int) {
	expanded := make([][]bool, ySize+2)

	for i := range expanded {
		expanded[i] = make([]bool, xSize+2)
	}

	for _, pos := range poses {
		expanded[pos[0]][pos[1]] = true
	}

	for r := 0; r < len(expanded); r++ {
		row := expanded[r]
		for c, spot := range row {
			spotUpper := spot
			spotLower := false
			if r+1 < len(expanded) {
				spotLower = expanded[r+1][c]
			}
			if spotUpper && !spotLower {
				fmt.Print("▀")
			} else if spotUpper && spotLower {
				fmt.Print("█")
			} else if !spotUpper && spotLower {
				fmt.Print("▄")
			} else if !spotUpper && !spotLower {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func getQuadrants(positions [][]int, xSize, ySize int) []int {
	quadrants := make([]int, 4)

	// quadrant 0,1,2,3
	// 0  1
	// 2  3

	for _, pos := range positions {
		// in middle
		if pos[0] == (xSize/2) || pos[1] == (ySize/2) {
			continue
		}

		quadrant := 0

		// x
		if pos[0] < (xSize / 2) {
			// 0 or 2
			// y
			if pos[1] < (ySize / 2) {
				quadrant = 0
			} else {
				quadrant = 2
			}
		} else {
			// 1 or 3
			// y
			if pos[1] < (ySize / 2) {
				quadrant = 1
			} else {
				quadrant = 3
			}
		}

		quadrants[quadrant]++
	}

	return quadrants
}

func parseValues(match []string) posVelValues {
	xPos, err := strconv.Atoi(match[1])
	if err != nil {
		panic(err)
	}

	yPos, err := strconv.Atoi(match[2])
	if err != nil {
		panic(err)
	}

	xVel, err := strconv.Atoi(match[3])
	if err != nil {
		panic(err)
	}

	yVel, err := strconv.Atoi(match[4])
	if err != nil {
		panic(err)
	}

	return posVelValues{
		xPos: xPos,
		yPos: yPos,
		xVel: xVel,
		yVel: yVel,
	}
}

type posVelValues struct {
	xPos int
	yPos int
	xVel int
	yVel int
}

func runMovements(vals posVelValues, numOfMoves, xSize, ySize int) (int, int) {
	newX := (vals.xPos + (vals.xVel * numOfMoves)) % xSize
	newY := (vals.yPos + (vals.yVel * numOfMoves)) % ySize

	if newX < 0 {
		newX += xSize
	}

	if newY < 0 {
		newY += ySize
	}

	return newX, newY
}
