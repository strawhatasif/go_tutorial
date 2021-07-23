package main

import "fmt"

/*
A couple oddities compared to languages like Python or Java:
1. argument types come after the variable name
2. the return type is declared before the opening curly brace
*/
func prepare_greeting(name string, color string) string {

	//similar to Java 10+ with local variable type inference
	var middle_text = "! We are thrilled your favorite color is "

	return "Welcome, " + name + middle_text + color + "!"
}

// This is a simple example demonstrating that functions in Go can return more than result!
func switch_names(last_name, first_name string) (string, string) {
	var corrected_first_name = last_name
	var corrected_last_name = first_name

	return corrected_first_name, corrected_last_name
}

func main() {
	//simple example (no user input, yet)
	//hard coded portions of a name
	var last_name = "Rachel "
	var first_name = "Pepperstaffere "
	var middle_name = "Belle "

	//whoops, we messed up in our hard coding of names! let's switch them
	corrected_first_name, corrected_last_name := switch_names(last_name, first_name)

	var full_name = corrected_first_name + middle_name + corrected_last_name

	fmt.Println(prepare_greeting(full_name, "black"))
}
