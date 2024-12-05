//go:build mage

package main

import (
	"fmt"

	wrapper "github.com/bmayfi3ld/advent-of-code/pkg/timer"
)

func TempA() error {
	fmt.Println("hello")
	defer wrapper.ProfileFunction("TempA")()

	return nil
}
