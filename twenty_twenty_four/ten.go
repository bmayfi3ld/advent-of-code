package twentytwentyfour

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bmayfi3ld/advent-of-code/pkg/cmd"
	wrapper "github.com/bmayfi3ld/advent-of-code/pkg/timer"
	"github.com/pkg/errors"
)

func init() {
	cmd.RegisterCommand("2024-10-1-test", func() error { return tenOne(tenTestInput) })
	cmd.RegisterCommand("2024-10-1-real", func() error { return tenOne(tenRealInput) })
	cmd.RegisterCommand("2024-10-2-test", func() error { return tenTwo(tenTestInput) })
	cmd.RegisterCommand("2024-10-2-real", func() error { return tenTwo(tenRealInput) })
}

func tenTwo(input string) error {
	fmt.Println("hello")
	defer wrapper.ProfileFunction("tenOne")()

	trees := parseTree(input)
	total9s := 0

	for _, tree := range trees {
		// spew.Dump(tree)
		count := count9sInTree(tree)

		// fmt.Printf("tree: %s has %d 9s\n", tree.originalGridString, count)

		total9s += count

		// return nil
	}

	fmt.Println(total9s)

	return nil
}

func tenOne(input string) error {
	fmt.Println("hello")
	defer wrapper.ProfileFunction("tenOne")()

	trees := parseTree(input)
	total9s := 0

	for _, tree := range trees {
		// spew.Dump(tree)
		count, _ := countUnique9sInTree(tree, nil)

		// fmt.Printf("tree: %s has %d 9s\n", tree.originalGridString, count)

		total9s += count

		// return nil
	}

	fmt.Println(total9s)

	return nil
}

func countUnique9sInTree(rootNode *node, alreadyFound9s []string) (int, []string) {
	count := 0

	if alreadyFound9s == nil {
		alreadyFound9s = []string{}
	}

	// check if we've already counted this tree
	nineAlreadyAdded := false

	for _, alreadyAdded := range alreadyFound9s {
		if alreadyAdded == rootNode.originalGridString {
			nineAlreadyAdded = true

			break
		}
	}

	if rootNode.value == 9 && !nineAlreadyAdded {
		count++

		alreadyFound9s = append(alreadyFound9s, rootNode.originalGridString)
	}

	for _, child := range rootNode.children {
		moreCount := 0
		moreCount, alreadyFound9s = countUnique9sInTree(child, alreadyFound9s)
		count += moreCount
	}

	return count, alreadyFound9s
}

func count9sInTree(rootNode *node) int {
	count := 0

	if rootNode.value == 9 {
		count++
	}

	for _, child := range rootNode.children {
		count += count9sInTree(child)
	}

	return count
}

// parse something like this
// 89010123
// 78121874
// 87430965
// 96549874
// 45678903
// 32019012
// 01329801
// 10456732
// into a slice of trees
// with 0 always being the root of the tree
// and 9 always being the end of the tree
// each number can only connect to the numbers above,below, and to the left and
// right of it
func parseTree(input string) []*node {
	// first make a 2d slice of ints
	grid := [][]int{}
	outNodes := []*node{}

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		row := []int{}

		for _, char := range line {
			charAsInt, err := strconv.Atoi(string(char))
			if err != nil {
				panic(errors.WithStack(err))
			}

			row = append(row, charAsInt)
		}

		grid = append(grid, row)
	}

	for r, row := range grid {
		for c, num := range row {
			if num == 0 {
				// start the tree
				outNodes = append(outNodes, &node{value: num, originalGridString: fmt.Sprintf("%d%d", r, c)})
				findRestOfTree(grid, r, c, outNodes[len(outNodes)-1])
			}
		}
	}

	return outNodes
}

type node struct {
	value              int // should be 0-9
	originalGridString string
	children           []*node
}

// given the grid, the row and col of the root node, and the root node
// find the rest of the tree
func findRestOfTree(grid [][]int, startRow, startCol int, rootNode *node) {
	nodeToSearchFor := rootNode.value + 1
	if nodeToSearchFor == 10 {
		return
	}

	// check the row above
	if startRow > 0 {
		if grid[startRow-1][startCol] == nodeToSearchFor {
			rootNode.children = append(rootNode.children, &node{value: nodeToSearchFor, originalGridString: fmt.Sprintf("%d%d", startRow-1, startCol)})
			findRestOfTree(grid, startRow-1, startCol, rootNode.children[len(rootNode.children)-1])
		}
	}

	// below
	if startRow < len(grid)-1 {
		if grid[startRow+1][startCol] == nodeToSearchFor {
			rootNode.children = append(rootNode.children, &node{value: nodeToSearchFor, originalGridString: fmt.Sprintf("%d%d", startRow+1, startCol)})
			findRestOfTree(grid, startRow+1, startCol, rootNode.children[len(rootNode.children)-1])
		}
	}

	// left
	if startCol > 0 {
		if grid[startRow][startCol-1] == nodeToSearchFor {
			rootNode.children = append(rootNode.children, &node{value: nodeToSearchFor, originalGridString: fmt.Sprintf("%d%d", startRow, startCol-1)})
			findRestOfTree(grid, startRow, startCol-1, rootNode.children[len(rootNode.children)-1])
		}
	}

	// right
	if startCol < len(grid[0])-1 {
		if grid[startRow][startCol+1] == nodeToSearchFor {
			rootNode.children = append(rootNode.children, &node{value: nodeToSearchFor, originalGridString: fmt.Sprintf("%d%d", startRow, startCol+1)})
			findRestOfTree(grid, startRow, startCol+1, rootNode.children[len(rootNode.children)-1])
		}
	}
}
