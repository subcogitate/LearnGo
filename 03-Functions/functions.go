/**
 * Multiple return values
 * Use it like any other type: Similar to functions in js
 * Function literals: When declaring a function inside another function, it remembers context.
 * Variadic functions: multiple variable parameters in func
 * http://golang.org/doc/codewalk/functions/
 * http://golang.org/ref/spec#Function_Types
 */
package main

import (
	"fmt"
)

type Salutation struct {
	name     string
	greeting string
}

type Printer func(string) () // Declaring a type Printer that is a function that takes a string and returns nothing hence the ()

func Greet(salutation Salutation, do Printer) { //Greet takes in a salutation and a Printer
	last, alternate := CreateMessage(salutation.name, salutation.greeting, "yo!", "sup!") //first takes first return from CreateMessage,

	do(last) //calls the passed in function with last greeting
	do(alternate)
}

func CreateMessage(name string, greeting ...string) (message string, alternate string) { //CreateMessage takes name and greeting as strings and returns two named strings

	message = greeting[len(greeting)-1] + " " + name //using last greeting passed "sup!"
	alternate = "Hey! " + name
	return //returns both message and alternate
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func CreatePrintFunction(custom string) Printer {
	return func(s string) { //this s will be passed in somewhere else.
		fmt.Println(s + custom) //uses context from when it was created and imprinting it onto the new function. In this case "!!closure!"
	}
}
func main() {
	var s = Salutation{"Bob", "Hello"}
	Greet(s, Print)                             //calls Greet with s, and Print func
	Greet(s, PrintLine)                         //calls Greet with s, and PrintLine func
	Greet(s, CreatePrintFunction("!!closure!")) //calls Greet with s and a function that is of type Printer that was created by CreatePrintFunction
}
