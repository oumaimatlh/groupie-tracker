package Handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
)

type Locations struct {
	Idlocation int      `json:"id"`
	Locations  []string `json:"locations"`
}

type Dates struct {
	IdDates int      `json:"id"`
	Dates   []string `json:"dates"`
}

type Relations struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

var (
	location  Locations
	dates     Dates
	relations Relations
)

func HandlerGroupe(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		HandleError(w, http.StatusBadRequest, "ID du groupe manquant")
		return
	}

	Id, err := strconv.Atoi(idStr)
	if err != nil {
		HandleError(w, http.StatusBadRequest, "ID invalide")
		return
	}

	var data struct {
		Artist
		Locations    []string
		ConcertDates []string
		Relations    []string
	}

	found := false
	for _, a := range artistes {
		if a.Id == Id {
			data.Artist = a
			found = true
			break
		}
	}
	if !found {
		HandleError(w, http.StatusNotFound, "Artiste non trouvé")
		return
	}

	tmpl, err := template.ParseFiles("./web/templates/groupArtist.html")
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "Erreur lors du chargement de la page")
		return
	}

	// ---------------------------------//
	loc, err := http.Get(data.Artist.Locations)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "Impossible de récupérer les lieux")
		return
	}
	defer loc.Body.Close()
	if err := json.NewDecoder(loc.Body).Decode(&location); err != nil {
		HandleError(w, http.StatusInternalServerError, "Erreur décodage des lieux")
		return
	}
	data.Locations = location.Locations

	// ---------------------------------//
	date, err := http.Get(data.Artist.ConcertDates)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "Impossible de récupérer les dates de concert")
		return
	}
	defer date.Body.Close()
	if err := json.NewDecoder(date.Body).Decode(&dates); err != nil {
		HandleError(w, http.StatusInternalServerError, "Erreur décodage des dates de concert")
		return
	}
	data.ConcertDates = dates.Dates

	// ---------------------------------//
	relation, err := http.Get(data.Artist.Relations)
	if err == nil {
		defer relation.Body.Close()
		var relData Relations
		if err := json.NewDecoder(relation.Body).Decode(&relData); err != nil {
			HandleError(w, http.StatusInternalServerError, "Erreur décodage des relations")
			return
		}
		for loc, dates := range relData.DatesLocations {
			for _, d := range dates {
				data.Relations = append(data.Relations, loc+" : "+d)
			}
		}
	}

	tmpl.Execute(w, data)
}
