package main 

import (
	"time"
	"fmt"
	"net/http"
)

func main() {
	start := time.Now()

	apis := []string {
		"http://management.azure.com",
		"http://dev.azure.com",
		"http://api.github.com",
		"http://outlook.office.com/",
		"http://api.somewhereintheinternet.com/",
		"http://graph.microsoft.com",
	}

	// Sin concurrencia
	/*
	for _, api := range apis {
		checkAPI(api)
	}
	*/

	ch := make(chan string)

	// Con concurrencia
	for _, api := range apis {
		go checkAPI(api, ch)
	}
	
	for i := 0;i < len(apis); i++ {
		fmt.Print(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("¡Listo! ¡Tomó %v segundos!\n", elapsed.Seconds())
}

func checkAPI(api string, ch chan string) {
	if _, err := http.Get(api); err != nil {
		ch <- fmt.Sprintf("ERROR: ¡%s está caído!\n", api)
		return
	}

	ch <- fmt.Sprintf("SUCCESS: ¡%s está en funcionamiento!\n", api)
}