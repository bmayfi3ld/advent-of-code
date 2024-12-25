package twentytwentyfour

import (
	"fmt"
	"regexp"

	"github.com/bmayfi3ld/advent-of-code/pkg/cmd"
	wrapper "github.com/bmayfi3ld/advent-of-code/pkg/timer"
)

func init() {
	cmd.RegisterCommand("2024-25-1-test", func() error { return twentyfiveOne(twentyFiveTestInput) })
	cmd.RegisterCommand("2024-25-1-real", func() error { return twentyfiveOne(twentyFiveRealInput) })
	cmd.RegisterCommand("2024-t-1-template", func() error { return tempOne(sixInput) })
	cmd.RegisterCommand("2024-t-1-template", func() error { return tempOne(sixInput) })
}

func twentyfiveOne(input string) error {
	fmt.Println("hello")
	defer wrapper.ProfileFunction("twentyfiveOne")()

	locks, keys := getLockAndKeys(input)

	matches := getLockAndKeyMatches(locks, keys)

	fmt.Println(matches)

	return nil
}

func getLockAndKeyMatches(locks [][]int, keys [][]int) int {
	matches := 0
	for _, lock := range locks {
		matches += getAllMatches(lock, keys)
	}

	return matches
}

func getAllMatches(lock []int, keys [][]int) int {
	matches := 0
	for _, key := range keys {
		if isNoOverlap(lock, key) {
			matches++
		}
	}

	return matches
}

// don't think it matters which order
func isNoOverlap(lock, key []int) bool {
	for i := range lock {
		if (lock[i] + key[i]) > 5 {
			return false
		}
	}

	return true
}

// returns locks and keys
func getLockAndKeys(input string) ([][]int, [][]int) {
	keyMatcher := regexp.MustCompile(`#####\n([\.#]{5})\n([\.#]{5})\n([\.#]{5})\n([\.#]{5})\n([\.#]{5})\n\.\.\.\.\.`)

	keySchematics := keyMatcher.FindAllStringSubmatch(input, -1)

	keys := [][]int{}

	for _, key := range keySchematics {
		keyMap := getHeightMap(key)

		keys = append(keys, keyMap)
	}

	lockMatcher := regexp.MustCompile(`\.\.\.\.\.\n([\.#]{5})\n([\.#]{5})\n([\.#]{5})\n([\.#]{5})\n([\.#]{5})\n#####`)

	lockSchematics := lockMatcher.FindAllStringSubmatch(input, -1)

	locks := [][]int{}

	for _, lock := range lockSchematics {
		lockMap := getHeightMap(lock)

		locks = append(locks, lockMap)
	}

	return locks, keys
}

func getHeightMap(heightSchematic []string) []int {
	heights := make([]int, 5)

	for i, heightLayer := range heightSchematic {
		if i == 0 {
			// regex 0 is the whole match
			continue
		}

		for ii, heightSlotForLayer := range heightLayer {
			if heightSlotForLayer == '#' {
				heights[ii]++
			}
		}
	}

	return heights
}
