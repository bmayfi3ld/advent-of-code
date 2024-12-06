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

func SixB() error {
	fmt.Println("hello")
	defer wrapper.ProfileFunction("SixB")()

	// test is 6
	// grid, startingPos := parseSixInput(sixTestInput)
	grid, startingPos := parseSixInput(sixInput)

	stuckRoutes := 0

	fmt.Println("initial route")

	gotOut, initialRoute := findRoute(startingPos, grid)
	if !gotOut {
		panic("route is broken")
	}

	fmt.Printf("checking %d blockers\n", len(initialRoute))
	bar := progressbar.Default(int64(len(initialRoute)))
	for _, potentialBlockSpot := range initialRoute {
		// fmt.Printf("checking if block r %d c %d\n", potentialBlockSpot.rowPos, potentialBlockSpot.colPos)
		bar.Add(1)

		// update grid
		grid[potentialBlockSpot.rowPos][potentialBlockSpot.colPos] = true

		gotOut, _ := findRoute(startingPos, grid)
		if !gotOut {
			stuckRoutes++
		}

		// revert grid
		grid[potentialBlockSpot.rowPos][potentialBlockSpot.colPos] = false
	}

	fmt.Println(stuckRoutes)

	return nil
}

// runs to see if the guard can get out
// returns true if he gets out
// returns the set of locations he visited
// if ha has visited each existing spot 2 times assume he is stuck
func findRoute(startingPostion guardPos, grid [][]bool) (bool, map[string]guardPos) {
	visitedSpots := map[string]guardPos{}
	currentPosition := startingPostion

	for {
		// record or update location
		//// by updating times visited
		pos, exists := visitedSpots[currentPosition.positionKey()]
		if exists {
			pos.timesVisited++
			visitedSpots[currentPosition.positionKey()] = pos
		} else {
			visitedSpots[currentPosition.positionKey()] = currentPosition
		}

		//// but also check if stuck
		atLeastOneStuckSpot := false
		for _, spots := range visitedSpots {
			// times stuck being 3, might accidentally hit it 2 times,
			// three should be stuck
			if spots.timesVisited > 3 {
				atLeastOneStuckSpot = true
				break
			}
		}

		if atLeastOneStuckSpot {
			return false, visitedSpots
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
	rowPos       int
	colPos       int
	dir          gridDirection
	timesVisited int // can be how many times the guard visited this spot
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
