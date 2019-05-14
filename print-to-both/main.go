package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "Stdout")
	fmt.Fprintln(os.Stderr, "Stderr")

}
