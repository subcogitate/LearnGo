package greeting

import (
	"fmt"
)

type Salutation struct {
	Name     string
	Greeting string
}

type RenameType struct {
	Name string
}

type Renamable interface {
	// define our interface
	Rename(newName string) //if the type implements a method of Rename, it can be passed as a Renamable
}

func (salutation *Salutation) Write(p []byte) (n int, err error) { // *Salutation is now a Writer!
	s := string(p)
	salutation.Rename(s)
	n = len(s)
	err = nil
	return
}

func (salutation *Salutation) Rename(newName string) { //must work on a pointer of Salutation or it will be working on a copy
	salutation.Name = newName
}

func (renametype *RenameType) Rename(newName string) { // This should also be implementing Renamable's methods
	renametype.Name = newName
}

type Printer func(string) ()

type Salutations []Salutation //setting up our named type

func (salutations Salutations) Greet(do Printer, isFormal bool) { //changing our function to be a method that acts on Salutations type

	for i, s := range salutations {
		message, alternate := CreateMessage(s.Name, s.Greeting)

		if prefix := GetPrefix(s.Name); isFormal {
			fmt.Print(i)
			do(prefix + message)
		} else {
			do(alternate)
		}
	}
}

func GetPrefix(name string) (prefix string) {

	prefixMap := map[string]string{
		"Bob":  "Mr ",
		"Joe":  "Dr ",
		"Amy":  "Dr ",
		"Mary": "Mrs ",
	}

	prefixMap["Joe"] = "Jr "

	delete(prefixMap, "Mary")

	if value, exists := prefixMap[name]; exists {
		return value
	}

	return "Dude"
}

func CreateMessage(name string, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "Hey! " + name
	return
}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}
