package main

import "fmt"

func divi(dividendo, divisor int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	validateZero(divisor)
	fmt.Println(dividendo / divisor)
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("No puedes dividir por cero")
	}
}

func main() {
	divi(100, 10)
	divi(200, 25)
	divi(34, 0)
	divi(100, 4)
}
