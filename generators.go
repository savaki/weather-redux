package main

import (
	"math/rand"
)

func fakeWeather(city City) Weather {
	_city := city
	temperature := rand.Intn(50)
	high := temperature + rand.Intn(5)
	low := temperature - rand.Intn(5)
	return Weather{
		City:        _city,
		Temperature: temperature,
		Highs:       high,
		Lows:        low,
	}
}
