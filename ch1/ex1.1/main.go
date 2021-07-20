// Ex1.1 modified the echo program to print the first argument.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep = "", ""
	for _, val := range os.Args[:] {
		s += sep + val
		sep = " "
	}
	fmt.Println(s)
}
