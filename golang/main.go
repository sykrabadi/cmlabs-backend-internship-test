package main

import "fmt"

func FizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%5 == 0 && i&3 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
func main() {
	FizzBuzz()
}
