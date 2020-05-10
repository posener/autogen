// Package main is a simple example that shows how to use autogen.
package main

import (
	"log"

	"github.com/posener/autogen"
)

//go:generate go run .

// animal is an example struct that is used to create the template.
type animal struct {
	Name  string
	Sound string
}

// NameLen is an example that shows that the given variable can have methods that than can be used
// in the template (Look for NameLen in animals.go.gotmpl.)
func (a animal) NameLen() int {
	return len(a.Name)
}

// animals is the variables that is used to generate the template.
var animals = []animal{{"cat", "meow"}, {"dog", "bark"}}

// main invokes the template generation.
func main() {
	// Automatically loads all *.gotmpl files in this directory and execute them to files in the
	// parent directory. This location can be configured with the `autogen.Location` function.
	err := autogen.Execute(animals)
	if err != nil {
		log.Fatal(err)
	}
}
