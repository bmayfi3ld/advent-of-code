package main

import (
	"fmt"

	"github.com/bmayfi3ld/advent-of-code/pkg/cmd"
	"github.com/pkg/errors"

	_ "github.com/bmayfi3ld/advent-of-code/twenty_twenty_four"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("Error: %+v", err)
	}
}

func run() error {
	// cpu profiling
	// cpuFile, err := os.Create("cpu.prof")
	// if err != nil {
	// 	return errors.WithStack(err)
	// }
	// defer cpuFile.Close()

	// err = pprof.StartCPUProfile(cpuFile)
	// if err != nil {
	// 	return errors.WithStack(err)
	// }
	// defer pprof.StopCPUProfile()

	err := cmd.GetRootCmd().Execute()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
