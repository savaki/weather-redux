package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func WeatherHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	state := req.FormValue("state")
	weathers, err := FindWeather(State(state))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	json.NewEncoder(w).Encode(weathers)
}

func main() {
	http.HandleFunc("/weather", WeatherHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
