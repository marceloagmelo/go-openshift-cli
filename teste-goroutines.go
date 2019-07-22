package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direto")
	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("indo")

	var input string
	fmt.Scanln(&input)
	fmt.Println("pronto")
}
