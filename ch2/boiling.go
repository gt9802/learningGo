package main

import "fmt"

// Visible to all files that belong to main package
const boilingF = 212.0

func main() {
	//visible to this  local scope
	var f = 212.0
	var c = (f - 32) * 5 / 9
	fmt.Printf("Boiling point  = %g°F or %g°C\n", f, c) // %g - format verb used to represent floating point numbers.
}
