package main

import (
	"fmt"
	"strconv"

	"github.com/golang/example/stringutil"
)

func main() {
	fmt.Println("Largest is", largestpalindrome(999))
}
func largestpalindrome(digits int) int {
	largestp := 0
	for b := digits; b != 0; b-- {
		for a := digits; a != 0; a-- {
			c := a * b
			if stringutil.Reverse(strconv.Itoa(c)) == strconv.Itoa(c) {
				if c > largestp {
					largestp = c
					fmt.Println(largestp)
				}

			}
		}
	}
	return largestp
}
