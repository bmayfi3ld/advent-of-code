//go:build mage

package twentytwentyfour

import (
	"fmt"

	"github.com/bmayfi3ld/advent-of-code/pkg/cmd"
	wrapper "github.com/bmayfi3ld/advent-of-code/pkg/timer"
)

func init() {
	cmd.RegisterCommand("2024-6-2-test", func() error { return sixB(sixTestInput, 6) })
	cmd.RegisterCommand("2024-6-2-real", func() error { return sixB(sixInput, 1831) })
}

func TempA() error {
	fmt.Println("hello")
	defer wrapper.ProfileFunction("TempA")()

	return nil
}
