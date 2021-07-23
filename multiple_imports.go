package main

/* Multiple imports
Syntactically, the below would be equivalent to:
import "fmt"
import "math"
*/
import (
	"fmt"
	"math"
)

func main() {
	// When referring to a function, you can only refer to its exported name (first letter capitalized)
	// printf is not valid, Printf is. the same applies to Sqrt
	fmt.Printf("Look ma! The square root of 9 is: %g", math.Sqrt(9))
}
