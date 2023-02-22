package server

import (
	"encoding/json"
	"net/http"
	"strings"
	//"fmt"
)

type relation struct {
	Index []struct {
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

type AllData struct {
	Id             int                 `json:"id"`
	Image          string              `json:"image"`
	Name           string              `json:"name"`
	Members        []string            `json:"members"`
	CreationDate   int                 `json:"creationDate"`
	FirstAlbum     string              `json:"firstAlbum"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Date struct {
	Index []struct {
		Id    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}
type Location struct {
	Index []struct {
		Id       int      `json:"id"`
		Location []string `json:"locations"`
	} `json:"index"`
}

func PageData(w http.ResponseWriter) []AllData {

	popup := []AllData{}
	Testo := relation{}
	data, err1 := Harvest("https://groupietrackers.herokuapp.com/api/artists")
	data2, err2 := Harvest("https://groupietrackers.herokuapp.com/api/relation")
	data3, _ := Harvest("https://groupietrackers.herokuapp.com/api/dates")
	data4, _ := Harvest("https://groupietrackers.herokuapp.com/api/locations")
	_ = data3
	_ = data4
	if err1 != nil || err2 != nil {
		error500(w)
	}

	json.Unmarshal(data, &popup)
	json.Unmarshal(data2, &Testo)
	test, _ := json.Marshal(Testo.Index)
	test = []byte(strings.Title(strings.ReplaceAll(string(test), "_", " ")))
	json.Unmarshal(test, &popup)
	return popup
}
