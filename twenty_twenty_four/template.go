//go:build mage

package main

import (
	"fmt"

	"github.com/bmayfi3ld/advent-of-code/utils/wrapper"
)

func TempA() error {
	fmt.Println("hello")
	defer wrapper.ProfileFunction("TempA")()

	return nil
}
