package twentytwentyfour

import (
	"fmt"

	"github.com/bmayfi3ld/advent-of-code/pkg/cmd"
	wrapper "github.com/bmayfi3ld/advent-of-code/pkg/timer"
)

func init() {
	cmd.RegisterCommand("2024-t-1-template", func() error { return tempOne(sixTestInput) })
	cmd.RegisterCommand("2024-t-1-template", func() error { return tempOne(sixTestInput) })
	cmd.RegisterCommand("2024-t-1-template", func() error { return tempOne(sixInput) })
	cmd.RegisterCommand("2024-t-1-template", func() error { return tempOne(sixInput) })
}

func tempOne(input string) error {
	fmt.Println("hello")
	defer wrapper.ProfileFunction("tempA")()

	return nil
}
