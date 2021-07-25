// This demonstrates a named or "naked" return

package main

import (
	"fmt"
	"strings"
)

/*
First, we use the Split function to split on the space delimiter
Then, we assign the first token to the last name variable, and second token to the first name variable
*/
func divideFirstAndLastName(fullName string) (first_name, last_name string) {
	name := strings.Split(fullName, " ")

	last_name = name[0]
	first_name = name[1]

	//notice, we are not specifiying what arguments to return
	return
}

func main() {
	//whoops! we are passing the last name and first name in the wrong order!
	fmt.Println(divideFirstAndLastName("Pepperstaffere Rachel"))
}
