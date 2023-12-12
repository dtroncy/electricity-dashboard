package main

import (
	"fmt"
	"net/http"
	"time"

	conso "github.com/dtroncy/enedis-consumption"
	tempo "github.com/dtroncy/tempo-calendar"
)

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
	formattedTodayDateConso := today.Format("2006-01-02")
	formattedYesterdayDateConso := yesterday.Format("2006-01-02")

	tempo.GetTempoCalendar(formattedYesterdayDateTempo, formattedTodayDateTempo)
	conso.GetDailyConsumption(formattedYesterdayDateConso, formattedTodayDateConso)
	conso.GetConsumptionLoadCurve(formattedYesterdayDateConso, formattedTodayDateConso)
	conso.GetConsumptionMaxPower(formattedYesterdayDateConso, formattedTodayDateConso)

	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(data)
}
