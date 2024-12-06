package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/bmayfi3ld/advent-of-code/pkg/cmd"

	_ "github.com/bmayfi3ld/advent-of-code/twenty_twenty_four"
)

func main() {
	cpuFile, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer cpuFile.Close()

	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	if err := cmd.GetRootCmd().Execute(); err != nil {
		fmt.Printf("Error: %+v", err)
	}
}
