package main

import "fmt"

// section prints a heading so consecutive demos stay readable instead of
// running together into one wall of lines.
func section(title string) {
	fmt.Println()
	fmt.Println("== " + title + " ==")
}

// Every value in Go has a type the compiler decides at compile time.
// Here is the same library member described with each basic type.
func main() {
	section("Basic types")

	// string holds text, like a member's name.
	var name string = "NamX1"
	fmt.Println(name)

	// int holds whole numbers; two of one type can share a declaration.
	var copiesBorrowed, copiesLimit int = 2, 5
	fmt.Println(copiesBorrowed, copiesLimit)

	// bool records a yes/no fact, like an active membership.
	var membershipActive bool = true
	fmt.Println(membershipActive)

	// float64 holds numbers with a fractional part, like a late fee.
	var lateFee float64 = 1.5
	fmt.Println(lateFee)

	// rune is a single Unicode code point; byte is an alias for uint8 (0-255).
	// A rune is an integer underneath, so Println would show 66. Use %c for B.
	var gradeLevel rune = 'B'
	var copiesInHand byte = 3
	fmt.Printf("grade %c, copies in hand %d\n", gradeLevel, copiesInHand)

	section("Declaration forms")

	// const fixes a value at compile time and cannot be reassigned.
	const libraryName = "City Library"
	fmt.Println(libraryName)

	// := declares and infers the type from the value on the right.
	dueDate := "2026-10-30"
	fmt.Println(dueDate)

	// An unassigned variable holds the zero value for its type: "" here.
	var shelfLabel string
	fmt.Printf("shelfLabel is %q\n", shelfLabel)

	section("All together")
	fmt.Printf("%s: %s, grade %c, borrowed %d of %d, late fee %.2f, shelf %q, due %s\n",
		libraryName, name, gradeLevel, copiesBorrowed, copiesLimit, lateFee, shelfLabel, dueDate)
}
