package Handlers

import (
	"encoding/json"
	"net/http"
	"html/template"
)

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

var artistes []Artist

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		HandleError(w, http.StatusNotFound, "Page non trouvée")
		return
	}

	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "Impossible de récupérer les artistes")
		return
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&artistes)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "Erreur lors du décodage des artistes")
		return
	}

	tmpl, err := template.ParseFiles("./web/templates/home.html")
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "Erreur lors du chargement de la page")
		return
	}

	tmpl.Execute(w, artistes)
}
