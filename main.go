package main

import (
	Handlers "GROUPIE-TRACKER/handlers"
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Serveur lancé sur : http://localhost:8080")

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("web/static")),
		),
	)


	http.HandleFunc("/", Handlers.Home)
	http.HandleFunc("/artist/", Handlers.HandlerGroupe)

	http.ListenAndServe(":8080", nil)
}
