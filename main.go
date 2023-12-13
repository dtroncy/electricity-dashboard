package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	tempo "github.com/dtroncy/tempo-calendar"
)

type TempoStruct struct {
	StartDate string              `json:"start_date"`
	EndDate   string              `json:"end_date"`
	Values    []TempoValuesStruct `json:"values"`
}

type TempoValuesStruct struct {
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Value       string `json:"value"`
	UpdatedDate string `json:"updated_date"`
}

func main() {

	http.HandleFunc("/lastdaydata", getLastDayData)

	fmt.Println("Démarrage du serveur sur le port 8080 ...")
	http.ListenAndServe(":8080", nil)

}

func getLastDayData(w http.ResponseWriter, r *http.Request) {

	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)
	formattedTodayDateTempo := today.Format("2006-01-02T15:04:05-07:00")
	formattedYesterdayDateTempo := yesterday.Format("2006-01-02T15:04:05-07:00")
	//formattedTodayDateConso := today.Format("2006-01-02")
	//formattedYesterdayDateConso := yesterday.Format("2006-01-02")

	tempoRequestResponse, err := tempo.GetTempoCalendar(formattedYesterdayDateTempo, formattedTodayDateTempo)

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

	fmt.Println(tempoResponse.Values[0].Value)

	//conso.GetDailyConsumption(formattedYesterdayDateConso, formattedTodayDateConso)
	//conso.GetConsumptionLoadCurve(formattedYesterdayDateConso, formattedTodayDateConso)
	//conso.GetConsumptionMaxPower(formattedYesterdayDateConso, formattedTodayDateConso)

	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(data)
}
