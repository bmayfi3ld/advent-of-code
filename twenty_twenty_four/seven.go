package twentytwentyfour

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/bmayfi3ld/advent-of-code/pkg/cmd"
	wrapper "github.com/bmayfi3ld/advent-of-code/pkg/timer"
	"github.com/pkg/errors"
)

func init() {
	cmd.RegisterCommand("2024-7-1-test", func() error { return sevenA(sevenTestInput) })
	cmd.RegisterCommand("2024-7-1-real", func() error { return sevenA(sevenRealInput) })
	cmd.RegisterCommand("2024-7-2-test", func() error { return sevenB(sevenTestInput) })
	cmd.RegisterCommand("2024-7-2-real", func() error { return sevenB(sevenRealInput) })
}

func sevenA(input string) error {
	fmt.Println("hello")
	defer wrapper.ProfileFunction("SevenA")()

	validSums := 0

	rows := strings.Split(strings.TrimSpace(input), "\n")

	for _, row := range rows {
		answer, nums, err := parseSevenRow(row)
		if err != nil {
			return errors.WithStack(err)
		}

		valid := validRowSeven(answer, nums, []operator{operatorAdd, operatorMul})
		if valid {
			validSums += answer
		}
	}

	// test answer is 3749
	fmt.Println(validSums)

	return nil
}

func parseSevenRow(row string) (int, []int, error) {
	split := strings.Split(row, " ")
	answerS := strings.TrimSuffix(split[0], ":")

	answer, err := strconv.Atoi(answerS)
	if err != nil {
		return 0, nil, errors.WithStack(err)
	}

	nums := []int{}
	for _, num := range split[1:] {
		op, err := strconv.Atoi(num)
		if err != nil {
			return 0, nil, errors.WithStack(err)
		}

		nums = append(nums, op)
	}

	return answer, nums, nil
}

// checks if operators + or * could be added
func validRowSeven(answer int, nums []int, ops []operator) bool {
	combinations := getAllOperators(len(nums)-1, ops)

	for _, combination := range combinations {
		found := checkCombinationAnswer(nums, combination, answer)
		if found {
			return true
		}
	}

	return false
}

func checkCombinationAnswer(nums []int, ops []operator, answer int) bool {
	type numOrOperator struct {
		isNum bool
		num   int
		op    operator
	}

	stack := []numOrOperator{}

	numsIndex := 0
	operatorsIndex := 0

	// build stack
	for numsIndex < len(nums) {
		// Push the current number onto the stack
		stack = append(stack, numOrOperator{
			isNum: true,
			num:   nums[numsIndex],
		})
		numsIndex++

		// If there's an operator at this point, push that as well
		if operatorsIndex < len(ops) {
			stack = append(stack, numOrOperator{
				isNum: false,
				op:    ops[operatorsIndex],
			})
			operatorsIndex++
		}
	}

	result := stack[0].num

	for i := 1; i < len(stack); i += 2 {
		currentOp := stack[i].op
		nextNum := stack[i+1].num
		switch currentOp {
		case operatorAdd:
			result = result + nextNum
		case operatorMul:
			result = result * nextNum
		case operatorConc:
			result = concat(result, nextNum)
		}

		if result > answer {
			return false
		}
	}

	return result == answer
}

func concat(a, b int) int {
	digits := 0
	temp := b
	for temp > 0 {
		digits++
		temp /= 10
	}

	// shift 'a' by 10^(digits) and add 'b'
	return a*int(math.Pow10(digits)) + b
}

type operator int

const (
	operatorUnknown operator = iota
	operatorAdd
	operatorMul
	operatorConc
)

func getAllOperators(numPositions int, ops []operator) [][]operator {
	if numPositions <= 0 {
		return [][]operator{}
	}

	// Result holds all combinations
	var result [][]operator

	// We'll use a recursive helper function
	var helper func(combination []operator)
	helper = func(combination []operator) {
		if len(combination) == numPositions {
			// Make a copy of combination to avoid referencing the same slice
			temp := make([]operator, numPositions)
			copy(temp, combination)
			result = append(result, temp)
			return
		}

		for _, v := range ops {
			// Append next value and recurse
			combination = append(combination, v)
			helper(combination)
			// Backtrack
			combination = combination[:len(combination)-1]
		}
	}

	helper([]operator{})
	return result
}

func sevenB(input string) error {
	fmt.Println("hello")
	defer wrapper.ProfileFunction("sevenB")()

	validSums := 0

	rows := strings.Split(strings.TrimSpace(input), "\n")

	for _, row := range rows {
		answer, nums, err := parseSevenRow(row)
		if err != nil {
			return errors.WithStack(err)
		}

		valid := validRowSeven(answer, nums, []operator{operatorAdd, operatorMul, operatorConc})
		if valid {
			validSums += answer
		}
	}

	// test answer is 11387
	fmt.Println(validSums)

	return nil
}
