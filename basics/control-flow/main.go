package main

import (
	"errors"
	"fmt"
)

// printCardType shows which membership card a member qualifies for.
func printCardType(age int) {
	if age >= 18 {
		fmt.Println("Adult Card")
	} else if age >= 13 {
		fmt.Println("Teen Card")
	} else {
		fmt.Println("Junior Card")
	}
}

// reserveCopy holds a copy for a member, refusing when none are left.
func reserveCopy(copiesLeft int) error {
	if copiesLeft <= 0 {
		return errors.New("no copies available")
	}
	return nil
}

// printMembershipStatus explains a membership status using a tagged switch,
// which compares status against each case value.
func printMembershipStatus(status string) {
	switch status {
	case "Active":
		fmt.Println("Welcome back!")
	case "Suspended":
		fmt.Println("Membership suspended!")
	case "Expired":
		fmt.Println("Membership expired!")
	default:
		fmt.Println("Unknown membership status!")
	}
}

// printOverdueFine reports how large the fine is, using a tagless switch where
// each case states its own boolean condition. Cases are tried top to bottom and
// stop at the first match, so they must cover every input without overlapping.
func printOverdueFine(days int) {
	switch {
	case days < 0:
		fmt.Println("Invalid date")
	case days == 0:
		fmt.Println("No fine")
	case days <= 7:
		fmt.Println("Small fine")
	case days <= 30:
		fmt.Println("Medium fine")
	default:
		fmt.Println("Large fine")
	}
}

// printShelfNumbers lists shelf numbers with the three-clause for loop:
// init, condition, post. Shelf numbers are 1-based, so the counter starts at 1.
func printShelfNumbers(shelves int) {
	for shelf := 1; shelf <= shelves; shelf++ {
		fmt.Printf("Shelf number %d\n", shelf)
	}
}

// printBookTitles uses the range form. With two variables the first is the
// index and the second is a copy of the element at that index.
func printBookTitles() {
	titles := []string{
		"Ghost in the wires",
		"Psychology of money",
		"White nights",
	}

	for i, title := range titles {
		fmt.Println(i, title)
	}
}

// printBooksByShelf walks a slice of slices and prints one line per book as
// "shelf title". The shelf number is the index of the OUTER range.
func printBooksByShelf() {
	bookLists := [][]string{
		{"White Night"},
		{"Ghost In The Wire"},
		{"Atomic Habits"},
	}

	for shelf, titles := range bookLists {
		for _, title := range titles {
			fmt.Printf("%d %s\n", shelf+1, title)
		}
	}
}

func main() {
	printCardType(15)
	printMembershipStatus("Active")
	printOverdueFine(2)
	printShelfNumbers(5)
	printBookTitles()
	printBooksByShelf()

	// reserveCopy returns an error nothing inspects yet. See interfaces-and-errors.
	// reserveCopy(10)
}
