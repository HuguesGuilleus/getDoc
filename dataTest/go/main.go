package main

import "fmt"

// The recipient of a message
const WHO = "world"

// The message for save the world
const (
	MESSAGE = "hello"
)

// Information of the calcul
var todo bool = true

// The sum of the calcul
var (
	sum = 0
)

// A example of structure
type fondation struct {
	a, b int
}

// The great main function
func main() {
	f := &fondation{
		a: 5,
		b: 14,
	}
	f.talk()
	if todo {
		sum = f.a + f.b
	}
	fmt.Println("Sum:", sum)
}

// A function to talk information
func (f *fondation) talk() {
	fmt.Println(MESSAGE, WHO)
	fmt.Println("Number:", f.a, f.b)
}
