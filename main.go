package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	add(3, 4)
	sub(5, 2)
}

func add(a, b int) {
	fmt.Println(a + b)
}

func sub(a, b int) {
	fmt.Println(a - b)
}
