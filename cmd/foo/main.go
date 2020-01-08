/*
	foo was a Foo-er for Foo-ing things.
	It now Baz-es things, but for a while it Bar-ed them
*/

package main

import (
	"fmt"

	"github.com/Intermernet/deprecated"
)

// Foo is the old hotness
func Foo() {
	fmt.Println("This code causes the program to terminate!")
	panic("Aaaaargh!")
}

// Bar was the new hotness, but could be better
func Bar() {
	fmt.Println("This code is generally better!")
}

// Baz is the NEW hotness!
func Baz() {
	fmt.Println("This is obviously the way to go!")
}

func main() {
	fmt.Printf("Error: %v\n", deprecated.ThisIs())
	fmt.Printf("Error: %v\n", deprecated.ThisIs(Foo, Bar))
	Baz()
}
