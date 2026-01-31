package Handlers

import (
	"encoding/json"
	"fmt"
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

var location Locations
var dates Dates
var relations Relations

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

	// Locations
	locResp, err := http.Get(data.Artist.Locations)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "Impossible de récupérer les lieux")
		return
	}
	defer locResp.Body.Close()
	if err := json.NewDecoder(locResp.Body).Decode(&location); err != nil {
		HandleError(w, http.StatusInternalServerError, "Erreur décodage des lieux")
		return
	}
	data.Locations = location.Locations

	// ConcertDates
	dateResp, err := http.Get(data.Artist.ConcertDates)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "Impossible de récupérer les dates de concert")
		return
	}
	defer dateResp.Body.Close()
	if err := json.NewDecoder(dateResp.Body).Decode(&dates); err != nil {
		HandleError(w, http.StatusInternalServerError, "Erreur décodage des dates de concert")
		return
	}
	data.ConcertDates = dates.Dates

	// Relations
	relResp, err := http.Get(data.Artist.Relations)
	if err == nil {
		defer relResp.Body.Close()
		var relData Relations
		if err := json.NewDecoder(relResp.Body).Decode(&relData); err != nil {
			HandleError(w, http.StatusInternalServerError, "Erreur décodage des relations")
			return
		}
		for loc, dates := range relData.DatesLocations {
			for _, d := range dates {
				data.Relations = append(data.Relations, fmt.Sprintf("%s : %s", loc, d))
			}
		}
	}

	tmpl.Execute(w, data)
}
