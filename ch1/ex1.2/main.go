// Ex1.1 modified the echo program to print the first argument.
package main

import (
	"fmt"
	"os"
)

func main() {
	for ind, val := range os.Args[:] {
		fmt.Println(ind, val)
	}
}
