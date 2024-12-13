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
	cmd.RegisterCommand("2024-11-1-test", func() error { return eleven(elevenTestInput, 25, 55312) })
	cmd.RegisterCommand("2024-11-1-real", func() error { return eleven(elevenRealInput, 25, 193607) })
	cmd.RegisterCommand("2024-11-2-real", func() error { return eleven(elevenRealInput, 75, 229557103025807) })
}

func eleven(input string, loopCount, answer int) error {
	fmt.Println("hello")

	defer wrapper.ProfileFunction("elevenOne")()

	currentSliceStr := strings.Split(strings.TrimSpace(input), " ")
	numberOfNumberOccurances := map[int]int{}

	for _, s := range currentSliceStr {
		num, err := strconv.Atoi(s)
		if err != nil {
			return errors.WithStack(err)
		}

		numberOfNumberOccurances[num]++
	}

	for range loopCount {
		numberOfNumberOccurances = applyRulesToSlice(numberOfNumberOccurances)
	}

	totalItems := 0
	for _, n := range numberOfNumberOccurances {
		totalItems += n
	}

	fmt.Printf("\nfound len %d, wanted len %d\n", totalItems, answer)

	return nil
}

// apply rules from day 11
func applyRulesToSlice(currentMap map[int]int) map[int]int {
	rule3Multiplier := 2024

	newMap := map[int]int{}

	// apply rules
	for num, existingOccuranceCount := range currentMap {
		if num == 0 {
			newMap[1] += existingOccuranceCount

			continue
		}

		if digits := getDigits(num); digits%2 == 0 {
			first, second := splitInt(num, digits)
			newMap[first] += existingOccuranceCount
			newMap[second] += existingOccuranceCount

			continue
		}

		newMap[num*rule3Multiplier] += existingOccuranceCount
	}

	return newMap
}

// number of base10 digits in the int
func getDigits(n int) int {
	if n == 0 {
		return 1
	}

	digits := 0
	temp := n

	for temp != 0 {
		temp /= 10
		digits++
	}

	return digits
}

// eg: 1234 -> 12, 34
func splitInt(intToSplit, digits int) (int, int) {
	// Count the number of digits
	if intToSplit == 0 {
		panic("Cannot split a single-digit number")
	}

	// Check if the number of digits is even
	if digits%2 != 0 {
		panic("Number of digits must be even")
	}

	// Calculate the divisor to split the number in half
	half := digits / 2
	divisor := 1

	for range half {
		divisor *= 10
	}

	// Split the number into two parts
	first := intToSplit / divisor
	second := intToSplit % divisor

	return first, second
}
