package twentytwentyfour

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/bmayfi3ld/advent-of-code/pkg/cmd"
	wrapper "github.com/bmayfi3ld/advent-of-code/pkg/timer"
	"github.com/davecgh/go-spew/spew"
)

func init() {
	cmd.RegisterCommand("2024-24-1-test", func() error { return twentyfourOne(twentyFourTestInput) })
	cmd.RegisterCommand("2024-24-1-real", func() error { return twentyfourOne(sixTestInput) })
	cmd.RegisterCommand("2024-t-1-template", func() error { return tempOne(sixInput) })
	cmd.RegisterCommand("2024-t-1-template", func() error { return tempOne(sixInput) })
}

func twentyfourOne(input string) error {
	fmt.Println("hello")
	defer wrapper.ProfileFunction("twentyfourOne")()

	inputMatcher := regexp.MustCompile(`([xy]\d\d): (\d)`)

	inputMatches := inputMatcher.FindAllStringSubmatch(input, -1)

	knownWires := map[string]bool{}

	for _, inputMatch := range inputMatches {
		value := false
		if inputMatch[2] == "1" {
			value = true
		}

		knownWires[inputMatch[1]] = value
	}

	operationMatcher := regexp.MustCompile(`(.{3}) (XOR|OR|AND) (.{3}) -> (.{3})`)

	operationMatches := operationMatcher.FindAllStringSubmatch(input, -1)

	wires := []wireGate{}

	for _, operationMatch := range operationMatches {
		op := andOperation

		if operationMatch[2] == "XOR" {
			op = xorOperation
		}
		if operationMatch[2] == "OR" {
			op = orOperation
		}

		wires = append(wires, wireGate{
			firstWire:  operationMatch[1],
			secondWire: operationMatch[3],
			op:         op,
			resultWire: operationMatch[4],
		})
	}

	solveWires(knownWires, wires)

	// spew.Dump(knownWires)

	binaryThing := []bool{}

	for knownWire, value := range knownWires {
		if !strings.HasPrefix(knownWire, "z") {
			continue
		}


	}

	return nil
}

func solveWires(knownWires map[string]bool, wires []wireGate) {
	for {
		solvedEverything := true
		for _, thingToSolve := range wires {
			first, haveFirst := knownWires[thingToSolve.firstWire]
			second, haveSecond := knownWires[thingToSolve.firstWire]

			if !haveFirst || !haveSecond {
				solvedEverything = false

				continue
			}

			switch thingToSolve.op {
			case andOperation:
				knownWires[thingToSolve.resultWire] = first && second
			case orOperation:
				knownWires[thingToSolve.resultWire] = first || second
			case xorOperation:
				knownWires[thingToSolve.resultWire] = first != second
			}
		}

		if solvedEverything {
			break
		}

	}
}

type operation int

const (
	andOperation operation = iota
	orOperation
	xorOperation
)

type wireGate struct {
	firstWire  string
	op         operation
	secondWire string
	resultWire string
}
