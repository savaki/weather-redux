package main

import (
	"code.google.com/p/go.net/context"
	"encoding/json"
	"net/http"
	"os"
	"time"
)

func Fail(w http.ResponseWriter, err error) {
	w.WriteHeader(500)
}

func WeatherHandler(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// read the user's input
	req.ParseForm()
	state := req.FormValue("state")

	// find the weather for all the cities in this state
	weathers, _ := findWeatherForState(ctx, state)

	// return as a json encoded block
	json.NewEncoder(w).Encode(weathers)
}

func main() {
	http.HandleFunc("/weather", WeatherHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
