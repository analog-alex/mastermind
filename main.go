package main

import (
	"fmt"
	"mastermind/pkg/listener"
)

func main() {
	fmt.Println("Booting up TCP listener...")

	// run listener on main goroutine
	listener.TPCListener(":3400")
}
