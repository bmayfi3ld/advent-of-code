package twentytwentyfour

import (
	"fmt"
	"strings"

	"github.com/bmayfi3ld/advent-of-code/pkg/cmd"
	wrapper "github.com/bmayfi3ld/advent-of-code/pkg/timer"
	"github.com/pkg/errors"
	"github.com/schollz/progressbar/v3"
)

func init() {
	cmd.RegisterCommand("2024-6-2-test", func() error { return sixB(sixTestInput, 6) })
	cmd.RegisterCommand("2024-6-2-real", func() error { return sixB(sixInput, 1831) })
}

func SixA() error {
	fmt.Println("hello")
	defer wrapper.ProfileFunction("SixA")()

	// test is 41
	// grid, startingPos := parseSixInput(sixTestInput)
	grid, startingPos := parseSixInput(sixInput)

	visitedSet := map[string]bool{}
	currentPosition := startingPos

	for {
		// record location
		visitedSet[currentPosition.positionKey()] = true

		// move
		newRow := currentPosition.getNewRowPos()
		newCol := currentPosition.getNewColPos()

		//// if leaving map exit
		if newRow >= len(grid) ||
			newCol >= len(grid[newRow]) ||
			newRow < 0 ||
			newCol < 0 {
			break
		}

		if grid[newRow][newCol] {
			//// if hitting a box turn
			turnedDir := currentPosition.getTurnedDirection()
			currentPosition.dir = turnedDir

		} else {
			//// just move forward
			currentPosition.rowPos = newRow
			currentPosition.colPos = newCol
		}

		// fmt.Printf("moved to %v\n", currentPosition)
		// time.Sleep(time.Second)
	}

	fmt.Println(len(visitedSet))

	return nil
}

func sixB(input string, expected int) error {
	fmt.Println("hello")
	defer wrapper.ProfileFunction("SixB")()

	grid, startingPos := parseSixInput(input)
	// preallocate the visit grid, grid is square
	visitedSpots := make([][]int, len(grid))
	for i := range visitedSpots {
		visitedSpots[i] = make([]int, len(grid[i]))
	}

	stuckRoutes := 0

	gotOut, visitedSpots := findRoute(startingPos, grid, visitedSpots)
	if !gotOut {
		visitGrid := strings.Builder{}
		for _, row := range visitedSpots {
			visitGrid.WriteString(fmt.Sprintf("%v\n", row))
		}

		return errors.Errorf("err could not find initial route, visited grid \n%s", visitGrid.String())
	}

	initialRoute := getMapFromVisitedGrid(visitedSpots)

	// cut it by 3/4 for profiling
	// keepCount := len(initialRoute) / 4
	// i := 0
	// for k := range initialRoute {
	// 	i++
	// 	if i > keepCount {
	// 		delete(initialRoute, k)
	// 	}
	// }

	fmt.Printf("checking %d blockers\n", len(initialRoute))
	bar := progressbar.Default(int64(len(initialRoute)))
	for _, potentialBlockSpot := range initialRoute {
		// fmt.Printf("checking if block r %d c %d\n", potentialBlockSpot.rowPos, potentialBlockSpot.colPos)
		err := bar.Add(1)
		if err != nil {
			return errors.WithStack(err)
		}

		// update grid
		grid[potentialBlockSpot.rowPos][potentialBlockSpot.colPos] = true

		gotOut, _ := findRoute(startingPos, grid, visitedSpots)
		if !gotOut {
			stuckRoutes++
		}

		clearVisitedGrid(visitedSpots)

		// revert grid
		grid[potentialBlockSpot.rowPos][potentialBlockSpot.colPos] = false
	}

	if stuckRoutes == expected {
		fmt.Printf("correct! with %d routes\n", expected)
	} else {
		fmt.Printf("wrong! with %d routes, expected %d\n", stuckRoutes, expected)
	}
	// fmt.Println(stuckRoutes)

	return nil
}

func clearVisitedGrid(visitedSpots [][]int) {
	for row := range visitedSpots {
		for col := range visitedSpots[row] {
			visitedSpots[row][col] = 0
		}
	}
}

func getMapFromVisitedGrid(visitedSpots [][]int) map[string]guardPos {
	result := make(map[string]guardPos)

	for row := range visitedSpots {
		for col, times := range visitedSpots[row] {
			if times > 0 {
				key := fmt.Sprintf("%d-%d", row, col)
				result[key] = guardPos{
					rowPos: row,
					colPos: col,
				}
			}
		}
	}

	return result
}

// runs to see if the guard can get out
// returns true if he gets out
// returns the set of locations he visited
// if ha has visited each existing spot 2 times assume he is stuck
func findRoute(startingPostion guardPos, grid [][]bool, visitedSpots [][]int) (bool, [][]int) {
	currentPosition := startingPostion

	// don't check if stuck until 4 turns have been made
	turnsSinceStuckCheck := 0
	turnsToCheck := 100 // check after turning this much

	// there is a strange bug here, where if this number is set too low, the program
	// will *sometimes* give the wrong answer (off by one)
	howManyVisitsConsideredStuck := 10

	for {
		// record or update location
		//// by updating times visited
		visitedSpots[currentPosition.rowPos][currentPosition.colPos]++

		//// but also check if stuck if we've made a square
		if turnsSinceStuckCheck == turnsToCheck {
			turnsSinceStuckCheck = 0
			atLeastOneStuckSpot := false
			for _, row := range visitedSpots {
				for _, spot := range row {
					if spot > howManyVisitsConsideredStuck {
						atLeastOneStuckSpot = true
						break
					}
				}
				if atLeastOneStuckSpot {
					break
				}
			}

			if atLeastOneStuckSpot {
				return false, nil
			}
		}

		// move
		newRow := currentPosition.getNewRowPos()
		newCol := currentPosition.getNewColPos()

		//// if leaving map exit
		if newRow < 0 ||
			newCol < 0 ||
			newRow >= len(grid) ||
			newCol >= len(grid[newRow]) {
			break
		}

		if grid[newRow][newCol] {
			//// if hitting a box turn
			turnedDir := currentPosition.getTurnedDirection()
			turnsSinceStuckCheck++
			currentPosition.dir = turnedDir

		} else {
			//// just move forward
			currentPosition.rowPos = newRow
			currentPosition.colPos = newCol
		}

		// fmt.Printf("moved to %v\n", currentPosition)
	}

	return true, visitedSpots
}

type gridDirection int

const (
	unknownGridDirection gridDirection = iota
	upGridDirection
	rightGridDirection
	downGridDirection
	leftGridDirection
)

type guardPos struct {
	rowPos int
	colPos int
	dir    gridDirection
}

// just smash x and y together
func (gp guardPos) positionKey() string {
	return fmt.Sprintf("%d-%d", gp.rowPos, gp.colPos)
}

func (gp guardPos) getNewRowPos() int {
	// if up x+1 if down x-1

	if gp.dir == upGridDirection {
		return gp.rowPos - 1
	}

	if gp.dir == downGridDirection {
		return gp.rowPos + 1
	}

	return gp.rowPos
}

func (gp guardPos) getNewColPos() int {
	if gp.dir == leftGridDirection {
		return gp.colPos - 1
	}

	if gp.dir == rightGridDirection {
		return gp.colPos + 1
	}

	return gp.colPos
}

func (gp guardPos) getTurnedDirection() gridDirection {
	newDir := gp.dir + 1

	if newDir == 5 { // out of bounds
		newDir = upGridDirection
	}

	return newDir
}

// point grid,  true if blocked
// x,y of the start
func parseSixInput(in string) ([][]bool, guardPos) {
	out := [][]bool{}
	startingPos := guardPos{}

	rows := strings.Split(strings.TrimSpace(in), "\n")

	for r, row := range rows {
		out = append(out, []bool{})
		for c, point := range row {
			sPoint := string(point)

			out[r] = append(out[r], sPoint == "#")

			if sPoint == "^" {
				startingPos = guardPos{
					rowPos: r,
					colPos: c,
					dir:    upGridDirection,
				}
			}
		}
	}

	return out, startingPos
}
