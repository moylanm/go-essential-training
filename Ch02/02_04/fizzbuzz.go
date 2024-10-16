/*
Print the numbers between 1 to 20, however
- If the number is divisible by 3, print "fizz" instead
- If the number is divisible by 5, print "buzz" instead
- If the number is divisible by 3 and 5, print "fizz buzz" instead
- Otherwise print the number
*/

package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 20; i++ {
		switch {
		case i % 3 == 0 && i % 5 == 0:
			fmt.Println("fizz buzz")
		case i % 3 == 0:
			fmt.Println("fizz")
		case i % 5 == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(i)
		}
	}
}
