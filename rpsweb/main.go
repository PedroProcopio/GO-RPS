package main

import (
	"log"
	"net/http"
	"rpsweb/handlers"
)

func main() {

	// Enrutador
	router := http.NewServeMux()

	//Manejador de Archivos estaticos
	fs := http.FileServer(http.Dir("static"))

	// Rutas de archivos estaticos
	router.Handle("/static/", http.StripPrefix("/static/", fs))
	// Rutas
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)
	router.HandleFunc("/newgame", handlers.NewGame)

	port := ":8080"
	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(port, router))

}
