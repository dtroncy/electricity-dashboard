package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	tempo "github.com/dtroncy/tempo-calendar"
)

type TempoStruct struct {
	Data TempoDataStruct `json:"tempo_like_calendars"`
}

type TempoDataStruct struct {
	StartDate string             `json:"start_date"`
	EndDate   string             `json:"end_date"`
	Values    []TempoValueStruct `json:"values"`
}

type TempoValueStruct struct {
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Value       string `json:"value"`
	UpdatedDate string `json:"updated_date"`
}

type TempoTodayAndTomorrow struct {
	Yesterday string `json:"yesterday"`
	Today     string `json:"today"`
	Tomorrow  string `json:"tomorrow"`
}

func main() {

	http.HandleFunc("/lastdaydata", getLastDayData)

	fmt.Println("Démarrage du serveur sur le port 8080 ...")
	fmt.Println("Aller sur http://localhost:8080/lastdaydata")
	http.ListenAndServe(":8080", nil)

}

func getLastDayData(w http.ResponseWriter, r *http.Request) {

	today := time.Now()
	today = time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
	afterTomorrow := today.AddDate(0, 0, 2)
	yesterday := today.AddDate(0, 0, -1)
	formattedYesterdayDateTempo := yesterday.Format("2006-01-02T15:04:05-07:00")
	formattedAfterTomorrowDateTempo := afterTomorrow.Format("2006-01-02T15:04:05-07:00")

	tempoRequestResponse, err := tempo.GetTempoCalendar(formattedYesterdayDateTempo, formattedAfterTomorrowDateTempo)

	if err != nil {
		fmt.Println("Erreur lors de la récupération des informations Tempo : ", err)
		return
	}

	var tempoResponse TempoStruct

	err = json.Unmarshal(tempoRequestResponse, &tempoResponse)

	if err != nil {
		fmt.Println("Erreur lors du décodage JSON:", err)
		return
	}

	fmt.Println("Demain : " + tempoResponse.Data.Values[0].Value)
	fmt.Println("Aujourd'hui : " + tempoResponse.Data.Values[1].Value)
	fmt.Println("Demain : " + tempoResponse.Data.Values[2].Value)

	tempoTodayAndTomorrow := TempoTodayAndTomorrow{
		Yesterday: tempoResponse.Data.Values[2].Value,
		Today:     tempoResponse.Data.Values[1].Value,
		Tomorrow:  tempoResponse.Data.Values[0].Value,
	}

	// Encode the data as JSON
	jsonData, err := json.Marshal(tempoTodayAndTomorrow)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON data to the response writer
	w.Write(jsonData)

}
