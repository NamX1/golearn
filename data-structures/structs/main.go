package main

import "fmt"

// section prints a heading so consecutive demos stay readable instead of
// running together into one wall of lines.
func section(title string) {
	fmt.Println()
	fmt.Println("== " + title + " ==")
}

// Book is one item on a shelf. A capitalised field is exported and visible
// outside the package; a lowercase one is not.
type Book struct {
	Title  string
	Author string
	Copies int
	OnLoan bool
	tags   []string
}

// Member is a library member. Structs can hold other structs, which is how you
// model a thing made of parts rather than one flat pile of fields.
type Member struct {
	Name  string
	Email string
	Books []Book
}

// Shelf is a numbered shelf. Every field is comparable, so two Shelf values can
// be compared with ==.
type Shelf struct {
	Number int
	Label  string
}

func main() {
	section("Named vs positional literals")

	// A named literal says which field is which, so it keeps working when a
	// field is added or reordered.
	named := Book{
		Title:  "Ghost in the Wires",
		Author: "Kevin Mitnick",
		Copies: 2,
	}
	fmt.Printf("%+v\n", named)

	// A positional literal gives values in order only, so adding or reordering
	// a field silently changes what lands where.
	positional := Book{"White Nights", "Dostoevsky", 1, false, nil}
	fmt.Printf("%+v\n", positional)

	section("Zero value")

	// Declaring a struct without assigning gives every field its own zero
	// value: "" for string, 0 for int, false for bool, nil for a slice.
	var blank Book
	fmt.Printf("%+v\n", blank)

	// Fields are reached with a dot, and assigning to one changes the struct.
	blank.Title = "Linux Basics"
	blank.Copies = 3
	fmt.Println(blank.Title, blank.Copies)

	section("Nested structs")

	// A struct field can itself be a struct or a slice of them.
	member := Member{
		Name:  "NamX1",
		Email: "namx1@example.com",
		Books: []Book{
			{Title: "Atomic Habits", Author: "James Clear", Copies: 1},
			{Title: "White Nights", Author: "Dostoevsky", Copies: 2},
		},
	}
	fmt.Println(member.Name, member.Email)
	fmt.Println("books held:", len(member.Books))
	fmt.Println("first book:", member.Books[0].Title)

	section("Copy semantics")

	// Assignment copies the whole struct. The two start equal, but changing
	// one afterwards leaves the other alone.
	original := Book{Title: "Dune", Copies: 1}
	duplicate := original
	duplicate.Copies = 99
	fmt.Println("original:", original.Copies)
	fmt.Println("duplicate:", duplicate.Copies)

	section("Comparability")

	// == compares every field, but only if every field is itself comparable.
	// Book has a []string field, so comparing two Books will not compile.
	//
	//	twin := Book{Title: "Dune", Copies: 1}
	//	fmt.Println(original == twin) // compile error: Book contains []string
	shelf := Shelf{Number: 1, Label: "Fiction"}
	other := Shelf{Number: 1, Label: "Fiction"}
	fmt.Println("same shelf:", shelf == other)
	fmt.Println("other shelf:", shelf == Shelf{Number: 2, Label: "Fiction"})
}
