package twentytwentyfour

import (
	"fmt"
	"strings"

	"github.com/bmayfi3ld/advent-of-code/pkg/cmd"
	wrapper "github.com/bmayfi3ld/advent-of-code/pkg/timer"
)

func init() {
	cmd.RegisterCommand("2024-8-1-test", func() error { return eightA(eightTestInput) })
	cmd.RegisterCommand("2024-8-1-real", func() error { return eightA(eightRealInput) })
	cmd.RegisterCommand("2024-8-2-test", func() error { return eightB(eightTestInput) })
	cmd.RegisterCommand("2024-8-2-test2", func() error { return eightB(eightOtherTestInput) })
	cmd.RegisterCommand("2024-8-2-real", func() error { return eightB(eightRealInput) })
}

func eightB(input string) error {
	fmt.Println("hello")
	defer wrapper.ProfileFunction("eightA")()

	grid := parseEightInput(input)

	matchedFreqs := map[gridLocation]bool{}

	for r, row := range grid {
		for c, spot := range row {
			if spot == '.' {
				continue
			}

			for rr, rrow := range grid {
				for cc, sspot := range rrow {
					if sspot == spot && rr != r && cc != c {
						diffR := (rr - r)
						diffC := (cc - c)
						curAntiR := rr
						curAntiC := cc
						for {
							if curAntiR < 0 || curAntiR >= len(grid) || curAntiC < 0 || curAntiC >= len(rrow) {
								break
							}

							matchedFreqs[gridLocation{row: curAntiR, col: curAntiC}] = true

							curAntiR += diffR
							curAntiC += diffC
						}
					}
				}
			}
		}
	}

	fmt.Println(len(matchedFreqs))

	return nil
}

func eightA(input string) error {
	fmt.Println("hello")
	defer wrapper.ProfileFunction("eightA")()

	grid := parseEightInput(input)

	matchedFreqs := map[gridLocation]bool{}

	for r, row := range grid {
		for c, spot := range row {
			if spot == '.' {
				continue
			}

			for rr, rrow := range grid {
				for cc, sspot := range rrow {
					if sspot == spot && rr != r && cc != c {
						// check if antinode would fit
						antiR := (rr - r) + rr
						antiC := (cc - c) + cc

						if antiR < 0 || antiR >= len(grid) || antiC < 0 || antiC >= len(rrow) {
							continue
						}

						matchedFreqs[gridLocation{row: antiR, col: antiC}] = true
					}
				}
			}

			// fancy wasn't working,  well... actually it was
			// it was just off by 1,
			// // get all possible matching freq locations
			// potLocs := getAllMatchingLocations(r, c, len(grid), len(row))

			// // count how many have the same freq
			// for _, potLoc := range potLocs {
			// 	if grid[potLoc.row][potLoc.col] == spot {
			// 		matchedFreqs[gridLocation{
			// 			row: potLoc.row - (r - potLoc.row),
			// 			col: potLoc.col - (c - potLoc.col),
			// 		}] = true
			// 	}
			// }
		}
	}

	fmt.Println(len(matchedFreqs))

	return nil
}

type gridLocation struct {
	row int
	col int
}

func getAllMatchingLocations(r, c, rLen, cLen int) []gridLocation {
	out := []gridLocation{}

	minR := calcGridMin(r)
	maxR := calcGridMax(r, rLen)

	minC := calcGridMin(c)
	maxC := calcGridMax(c, cLen)

	for rIter := minR; rIter <= maxR; rIter++ {
		for cIter := minC; cIter <= maxC; cIter++ {
			if rIter == r && cIter == c {
				continue
			}

			out = append(out, gridLocation{row: rIter, col: cIter})
		}
	}

	return out
}

func calcGridMin(cur int) int {
	// to close to edge to get anywhere
	if cur == 0 {
		return 0
	}

	if cur == 1 {
		return 1
	}

	return (cur / 2) + (cur % 2)
}

func calcGridMax(cur, len int) int {
	maxIndex := len - 1

	// to close to edge to get anywhere
	if cur == maxIndex {
		return maxIndex
	}

	if cur == maxIndex-1 {
		return maxIndex
	}

	return ((maxIndex - cur) / 2) + cur
}

// point grid,  stores the letter of the freq
// x,y of the start
func parseEightInput(in string) [][]rune {
	out := [][]rune{}

	rows := strings.Split(strings.TrimSpace(in), "\n")

	for r, row := range rows {
		out = append(out, []rune{})
		for _, point := range row {
			out[r] = append(out[r], point)
		}
	}

	return out
}
