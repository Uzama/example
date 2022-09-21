package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	add(3, 4)
}

func add(a, b int) {
	fmt.Println(a + b)
}
