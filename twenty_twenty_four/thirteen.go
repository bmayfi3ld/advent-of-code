package twentytwentyfour

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/bmayfi3ld/advent-of-code/pkg/cmd"
	wrapper "github.com/bmayfi3ld/advent-of-code/pkg/timer"
	"github.com/schollz/progressbar/v3"
)

func init() {
	cmd.RegisterCommand("2024-13-1-test", func() error { return thirteen(thirteenTestInput, false) })
	cmd.RegisterCommand("2024-13-1-real", func() error { return thirteen(thirteenRealInput, false) })
	cmd.RegisterCommand("2024-13-2-real", func() error { return thirteen(thirteenRealInput, true) })
}

func thirteen(input string, partTwo bool) error {
	fmt.Println("hello")
	defer wrapper.ProfileFunction("thirteenOne")()

	re := regexp.MustCompile(`Button A: X\+(\d{1,3}), Y\+(\d{1,3})\nButton B: X\+(\d{1,3}), Y\+(\d{1,3})\nPrize: X=(\d{1,6}), Y=(\d{1,6})`)

	matches := re.FindAllStringSubmatch(input, -1)

	totalTokens := 0

	bar := progressbar.Default(int64(len(matches)))

	for _, match := range matches {
		_ = bar.Add(1)

		tokensUsed := processOneGame(match, partTwo)

		if tokensUsed == 0 {
			continue
		}

		totalTokens += tokensUsed
	}

	fmt.Println(totalTokens)

	return nil
}

func processOneGame(inputs []string, increase bool) int {
	values := getAllNumbers(inputs)

	if increase {
		values.pX += 10000000000000
		values.pY += 10000000000000
	}

	// linear algebra magic, (solved by chatgpt, not me)
	aInt := int(float64(values.bY*values.pX-values.pY*values.bX) / float64(values.aX*values.bY-values.aY*values.bX))
	bInt := int(float64(values.aX*values.pY-values.pX*values.aY) / float64(values.aX*values.bY-values.aY*values.bX))

	// check if it solved it (might not of been 2 ints)
	if aInt*values.aX+bInt*values.bX != values.pX ||
		aInt*values.aY+bInt*values.bY != values.pY {
		return 0
	}

	if !increase && (aInt > 100 || bInt > 100) {
		return 0
	}

	return aInt*3 + bInt
}

type gameRound struct {
	aX int
	aY int
	bX int
	bY int
	pX int
	pY int
}

func getAllNumbers(inputs []string) gameRound {
	aX, err := strconv.Atoi(inputs[1])
	if err != nil {
		panic(err)
	}

	aY, err := strconv.Atoi(inputs[2])
	if err != nil {
		panic(err)
	}

	bX, err := strconv.Atoi(inputs[3])
	if err != nil {
		panic(err)
	}

	bY, err := strconv.Atoi(inputs[4])
	if err != nil {
		panic(err)
	}

	pX, err := strconv.Atoi(inputs[5])
	if err != nil {
		panic(err)
	}

	pY, err := strconv.Atoi(inputs[6])
	if err != nil {
		panic(err)
	}

	return gameRound{
		aX: aX,
		aY: aY,
		bX: bX,
		bY: bY,
		pX: pX,
		pY: pY,
	}
}
