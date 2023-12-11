package main

import (
	conso "github.com/dtroncy/enedis-consumption"
	tempo "github.com/dtroncy/tempo-calendar"
)

func main() {

	tempo.GetTempoCalendar("2023-12-08T00:00:00+01:00", "2023-12-11T00:00:00+01:00")
	conso.GetDailyConsumption("2023-12-08", "2023-12-11")
	conso.GetConsumptionLoadCurve("2023-12-08", "2023-12-11")
	conso.GetConsumptionMaxPower("2023-12-08", "2023-12-11")

}
