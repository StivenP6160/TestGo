package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	if t := time.Now(); t.Hour() < 12 {
		fmt.Println("Mañana!")
	} else if t.Hour() < 17 {
		fmt.Println("Tarde!")
	} else {
		fmt.Println("Noche!")
	}

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Go run -> Windows")
	case "linux":
		fmt.Println("Go run -> linux")
	case "darwin":
		fmt.Println("Go run -> macOS")
	default:
		fmt.Println("Go run -> Otro OS")
	}

	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println(t.Hour(), "Mañana!")
	case t.Hour() < 17:
		fmt.Println("Tarde!")
	default:
		fmt.Println("Noche!")
	}
}
