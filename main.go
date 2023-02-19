package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		fizzbuzz(i)
	}
}

func fizzbuzz(i int) {
	if i%3 == 0 && i%5 == 0 {
		fmt.Println(i, "FizzBuzz")
	} else if i%5 == 0 {
		fmt.Println(i, "Buzz")
	} else if i%3 == 0 {
		fmt.Println(i, "fizz")
	} else {
		fmt.Println(i)
	}
}
