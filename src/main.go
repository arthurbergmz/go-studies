package main

import (
	. "./tests"
	"./tests/closure"
	"fmt"
)

func ensureTrue (result bool) {
	if result {
		fmt.Println("Test passed successfully.")
	} else {
		fmt.Println("Test didn't pass: unexpected result.")
	}
}

func main () {
	Test(ensureTrue, closure.TwoIntegersTest([]int{ 3, 7 }))
}