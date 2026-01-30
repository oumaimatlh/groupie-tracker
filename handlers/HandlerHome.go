package Handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	//"io"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func Home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "ERROR", 404)
	}

	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	var artistes []Artist

	err = json.NewDecoder(res.Body).Decode(&artistes)
	if err != nil {
		log.Fatal(err)
	}

	template, _ := template.ParseFiles("./web/templates/home.html")
	template.Execute(w, artistes)
}


