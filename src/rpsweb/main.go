package main

import (
	"net/http"
	"rpsweb/handlers"
	"log"
)

func main() {
	// Crear enrutador
	router := http.NewServeMux()

	// Manejador para servir los archivos estáticos
	fs := http.FileServer(http.Dir("static"))
	// Ruta para acceder alos archivos estáticos
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	// Configurar rutas
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	port := ":8080"
	log.Printf("Servidor escuchando en http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))	
}