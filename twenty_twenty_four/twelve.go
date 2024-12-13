package twentytwentyfour

import (
	"fmt"
	"strings"

	"github.com/bmayfi3ld/advent-of-code/pkg/cmd"
	wrapper "github.com/bmayfi3ld/advent-of-code/pkg/timer"
)

func init() {
	cmd.RegisterCommand("2024-12-1-test", func() error { return twelveOne(twelveTestInput) })
	cmd.RegisterCommand("2024-12-1-simple", func() error { return twelveOne(twelveSimpleInput) })
	cmd.RegisterCommand("2024-12-1-real", func() error { return twelveOne(twelveRealInput) })
	cmd.RegisterCommand("2024-12-2-test", func() error { return twelveTwo(twelveTestInput) })
	cmd.RegisterCommand("2024-12-2-simple", func() error { return twelveTwo(twelveSimpleInput) })
	cmd.RegisterCommand("2024-12-2-mob", func() error { return twelveTwo(twelveMobiusInput) })
	cmd.RegisterCommand("2024-12-2-e", func() error { return twelveTwo(twelveEInput) })
	cmd.RegisterCommand("2024-12-2-real", func() error { return twelveTwo(twelveRealInput) })
}

func twelveOne(input string) error {
	fmt.Println("hello")
	defer wrapper.ProfileFunction("twelveOne")()

	trees := parseRegionTree(input)

	total := 0
	for _, tree := range trees {
		// spew.Dump(tree)

		total += calcRegionCostPerimeter(tree)
	}

	fmt.Println(total)

	return nil
}

func twelveTwo(input string) error {
	fmt.Println("hello")
	defer wrapper.ProfileFunction("twelveOne")()

	trees := parseRegionTree(input)

	total := 0
	for _, tree := range trees {
		// spew.Dump(tree)

		total += calcRegionCostSides(tree)
	}

	// 895486 is too low

	fmt.Println(total)

	return nil
}

func calcRegionCostPerimeter(reg *region) int {
	// get size
	area := len(reg.spots)

	// get perimeter
	perimeter := getSpotsPerimeter(reg.spots)

	return area * perimeter
}

func calcRegionCostSides(reg *region) int {
	// get size
	area := len(reg.spots)

	// get perimeter
	perimeter := getSpotsSides(reg.spots)

	return area * perimeter
}

func getSpotsPerimeter(spots map[gridLocation]bool) int {
	type faceRef struct {
		row float64
		col float64
	}

	faces := map[faceRef]int{}

	for spot := range spots {
		// each spot has 4
		faces[faceRef{row: float64(spot.row), col: float64(spot.col) + .5}] += 1
		faces[faceRef{row: float64(spot.row), col: float64(spot.col) - .5}] += 1
		faces[faceRef{row: float64(spot.row) + .5, col: float64(spot.col)}] += 1
		faces[faceRef{row: float64(spot.row) - .5, col: float64(spot.col)}] += 1
	}

	perimeter := 0
	for _, occurences := range faces {
		if occurences == 1 {
			perimeter++
		}
	}

	return perimeter
}

// should be same as perimeter, but odd corners
func getSpotsSides(spots map[gridLocation]bool) int {
	type cornerRef struct {
		row float64
		col float64
	}

	faces := map[cornerRef]int{}

	for spot := range spots {
		// each spot has 4
		faces[cornerRef{row: float64(spot.row) - .5, col: float64(spot.col) - .5}] += 1
		faces[cornerRef{row: float64(spot.row) - .5, col: float64(spot.col) + .5}] += 1
		faces[cornerRef{row: float64(spot.row) + .5, col: float64(spot.col) + .5}] += 1
		faces[cornerRef{row: float64(spot.row) + .5, col: float64(spot.col) - .5}] += 1
	}

	sides := 0
	for _, occurences := range faces {
		if occurences == 1 || occurences == 3 {
			sides++
		}
	}

	return sides
}

// parse something like this
// RRRRIICCFF
// RRRRIICCCF
// VVRRRCCFFF
// VVRCCCJFFF
// VVVVCJJCFE
// VVIVCCJJEE
// VVIIICJJEE
// MIIIIIJJEE
// MIIISIJEEE
// MMMISSJEEE
// into a slice of trees
func parseRegionTree(input string) []*region {
	// first make a 2d slice of ints
	grid := [][]rune{}
	outNodes := []*region{}

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		row := []rune{}

		for _, char := range line {
			row = append(row, char)
		}

		grid = append(grid, row)
	}

	for r, row := range grid {
		for c, num := range row {
			// check all the regions so far
			foundClaimedSpot := false
			for _, region := range outNodes {
				for spot := range region.spots {
					if spot.col == c && spot.row == r {
						foundClaimedSpot = true
						break
					}
				}
				if foundClaimedSpot {
					break
				}
			}

			if foundClaimedSpot {
				continue
			}

			// if not already in a region
			outNodes = append(outNodes,
				&region{
					value: num,
					spots: map[gridLocation]bool{},
				},
			)
			findRestOfRegion(grid, r, c, outNodes[len(outNodes)-1])

		}
	}

	return outNodes
}

type region struct {
	value rune
	spots map[gridLocation]bool
}

// given the grid, the row and col of the root node, and the root node
// find the rest of the tree
func findRestOfRegion(grid [][]rune, startRow, startCol int, reg *region) {
	_, found := reg.spots[gridLocation{row: startRow, col: startCol}]
	if found {
		// was already checked
		return
	}

	reg.spots[gridLocation{row: startRow, col: startCol}] = true

	nodeToSearchFor := reg.value

	// check the row above
	if startRow > 0 {
		if grid[startRow-1][startCol] == nodeToSearchFor {
			findRestOfRegion(grid, startRow-1, startCol, reg)
		}
	}

	// below
	if startRow < len(grid)-1 {
		if grid[startRow+1][startCol] == nodeToSearchFor {
			findRestOfRegion(grid, startRow+1, startCol, reg)
		}
	}

	// left
	if startCol > 0 {
		if grid[startRow][startCol-1] == nodeToSearchFor {
			findRestOfRegion(grid, startRow, startCol-1, reg)
		}
	}

	// right
	if startCol < len(grid[0])-1 {
		if grid[startRow][startCol+1] == nodeToSearchFor {
			findRestOfRegion(grid, startRow, startCol+1, reg)
		}
	}
}
