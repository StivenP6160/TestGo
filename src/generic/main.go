package main

import "fmt"

func PrintList(list ...any) {
	for _, value := range list {
		fmt.Println(value)
	}
}

func main() {
	PrintList("Alex", "Roel", "Juan", "Pedro")
	PrintList(100, 456, 789, 456, 626)
	PrintList("Hola", 452, 6.87, true)
}