// This demonstrates a named or "naked" return

package main

import (
	"fmt"
	"strings"
)

/*
First, we use the Split function to split on the space delimiter
Then, we assign the first token to the first name variable, and second token to the last name variable
*/
func divideFirstAndLastName(fullName string) (first_name, last_name string) {
	name := strings.Split(fullName, " ")

	first_name = name[0]
	last_name = name[1]

	//notice, we are not specifiying what arguments to return
	return
}

func main() {
	fmt.Println(divideFirstAndLastName("Rachel Pepperstaffere"))
}
