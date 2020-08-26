package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Station struct {
	Name string `json:"name"`
	Crs  string `json:"crs"`
}

var station Station
var stations []Station

func readCsv() [][]string {

	resp, err := os.Open("./stationcrs.csv")

	if err != nil {
		panic(err)
	}

	defer resp.Close()

	if err != nil {
		panic(err)
	}

	r := csv.NewReader(resp)
	r.FieldsPerRecord = -1

	rows, err := r.ReadAll()
	if err != nil {
		log.Fatalln("Cannot read csv data:", err.Error())
		os.Exit(1)
	}
	return rows
}

func getCrs(rows [][]string) []Station {

	for _, each := range rows {
		station.Name = each[0]
		station.Crs = each[1]
		stations = append(stations, station)
	}

	return stations

}

func readCrs(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(getCrs(readCsv()))

}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/crs", readCrs)
	http.Handle("/", r)
	log.Println(http.ListenAndServe(":"+os.Getenv("PORT"), nil))

}
