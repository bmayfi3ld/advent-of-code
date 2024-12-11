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
	cmd.RegisterCommand("2024-9-1-test", func() error { return nineOne(nineTestInput) })
	cmd.RegisterCommand("2024-9-1-real", func() error { return nineOne(nineRealInput) })
	cmd.RegisterCommand("2024-9-2-test", func() error { return nineTwo(nineTestInput) })
	cmd.RegisterCommand("2024-9-2-real", func() error { return nineTwo(nineRealInput) })
}

func nineOne(input string) error {
	fmt.Println("hello")

	defer wrapper.ProfileFunction("nineOne")()

	parsed := expandDiskFormat(input)

	for i, spot := range parsed {
		if spot == -1 {
			for ii := len(parsed) - 1; ii > 0; ii-- {
				if parsed[ii] != -1 {
					parsed[i] = parsed[ii]
					parsed[ii] = -1

					// fmt.Println(parsed)

					break
				}
			}
		}

		// dumb way to check if the rest of the spots are free space
		noMoreFiles := true

		for ii := i + 1; ii < len(parsed); ii++ {
			if parsed[ii] != -1 {
				noMoreFiles = false

				break
			}
		}

		if noMoreFiles {
			break
		}
	}

	// fmt.Println(parsed)

	fmt.Println(calcSum(parsed))

	return nil
}

func nineTwo(input string) error {
	fmt.Println("hello")

	defer wrapper.ProfileFunction("nineOne")()

	parsed := expandDiskFormat(input)

	for i := len(parsed) - 1; i > 0; i-- {
		if parsed[i] == -1 {
			continue
		}

		// check how big the file is
		sizeOfFile := 0
		fileNum := parsed[i]

		for ii := i; ii > 0; ii-- {
			if parsed[ii] != fileNum {
				break
			}

			sizeOfFile++
		}

		// from the front, find the next spot that will fit
		for ii := 0; ii < len(parsed); ii++ { //nolint: varnamelen
			if parsed[ii] != -1 {
				continue
			}

			// don't move the file to the right
			// not past where it already is
			if ii > i-sizeOfFile {
				break
			}

			// see how big it is
			sizeOfEmptySpot := 0

			for iii := ii; iii < len(parsed); iii++ {
				if parsed[iii] != -1 {
					break
				}

				sizeOfEmptySpot++
			}

			// if it doesn't fit, skip the empty spot
			if sizeOfFile > sizeOfEmptySpot {
				ii += sizeOfEmptySpot

				continue
			}

			// if fits swap it
			for f := range sizeOfFile {
				parsed[i-f] = -1
				parsed[ii+f] = fileNum
			}

			// fmt.Println(parsed)

			break
		}

		i -= (sizeOfFile - 1)

	}

	// fmt.Println(parsed)

	fmt.Println(calcSum(parsed))

	return nil
}

func calcSum(formatted []int) int {
	sum := 0

	for i, item := range formatted {
		if item == -1 {
			continue
		}

		sum += i * item
	}

	return sum
}

// eg from
// 2333133121414131402 to 00...111...2...333.44.5555.6666.777.888899
func expandDiskFormat(in string) []int {
	out := []int{}
	file := true // true if adding file rune, false if adding empty space runes
	fileIDIndex := 0

	for _, spot := range strings.TrimSpace(in) {
		spotD, err := strconv.Atoi(string(spot))
		if err != nil {
			panic(errors.WithStack(err))
		}

		if file {
			for range spotD {
				out = append(out, fileIDIndex)
			}

			file = false
			fileIDIndex++

			continue
		}

		for range spotD {
			out = append(out, -1)
		}

		file = true
	}

	return out
}
