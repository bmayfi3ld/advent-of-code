package wrapper

import (
	"fmt"
	"time"
)

// usage: 	defer ProfileFunction("exampleFunction")()
func ProfileFunction(name string) func() {
	start := time.Now()
	return func() {
		elapsed := time.Since(start)
		fmt.Printf("%s execution time: %s\n", name, elapsed)
	}
}
