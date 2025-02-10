package main

import (
	"errors"
	"fmt"
)

func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No se puede devidir por cero")
	}
	return dividendo / divisor, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	println("Resultado:", result)
}

/*
func main() {
	str := "123f"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Número:", num)
}
*/
