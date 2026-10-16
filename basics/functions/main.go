package main

import "fmt"

// section prints a heading so consecutive demos stay readable instead of
// running together into one wall of lines.
func section(title string) {
	fmt.Println()
	fmt.Println("== " + title + " ==")
}

// greetMember prints a welcome message. In `name string`, the name is the
// parameter and string is its type.
func greetMember(name string) {
	fmt.Printf("Welcome %s, into our library!\n", name)
}

// canBorrow reports whether a copy is available. Go can return several values
// at once; here a bool and a short reason travel together.
func canBorrow(copiesLeft int) (bool, string) {
	if copiesLeft <= 0 {
		return false, "no copies available"
	}
	return true, "copy available"
}

func main() {
	section("greetMember")
	greetMember("NamX1")

	section("canBorrow")
	// Two return values are assigned with the usual = or :=.
	allowed, reason := canBorrow(0)
	fmt.Println(allowed, reason)
}
