package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "123f"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Número:", num)
}
