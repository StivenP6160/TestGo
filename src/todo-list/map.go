package main

import "fmt"

func main() {
	colors := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}

	fmt.Println(colors["rojo"])
	colors["Negro"] = "#000000"

	fmt.Println(colors)

	if Valor, ok := colors["verder"]; ok {
		fmt.Println(Valor)
	} else {
		fmt.Println("No existe esta clave")
	}

	delete(colors, "azul")
	fmt.Println(colors)

}
