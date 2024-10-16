/*
An "even ended number" is a number whose first and last digit are the same.

You mission, should you choose to accept it, is to count how many "even ended numbers" are
there that are a multiplication of two 4 digit numbers.
*/

package main

import (
	"fmt"
)

func main() {
	count := 0

	for i := 1000; i < 9999; i++ {
		for j := i; j < 9999; j++ {
			n := i * j
			s := fmt.Sprintf("%d", n)

			if s[0] == s[len(s) - 1] {
				count++
			}
		}
	}

	fmt.Printf("There are %d even-ended numbers between 1000 and 9999.\n", count)
}
