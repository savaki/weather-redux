package main

import (
	"code.google.com/p/go.net/context"
	"github.com/savaki/openweathermap"
	"github.com/savaki/par"
)

type City string

type Weather struct {
	City        City
	Temperature int
	Highs       int
	Lows        int
}

func findWeatherForCities(ctx context.Context, cities []City) ([]Weather, error) {
	responses, err := func() (chan Weather, error) {
		redundancy := 2

		// create a channel to hold our responses
		responses := make(chan Weather, len(cities)*redundancy)
		defer close(responses)

		// convert our list of cities into functions that can be executed redundantly
		// and in parallel
		requests := make(chan par.RequestFunc, len(cities))
		for _, city := range cities {
			requests <- findWeatherForCityFunc(city, responses)
		}
		close(requests)

		// Here's the meat of the issue.  Rather than using generics, we're using a
		// closure to wrap the details of our weather implementation.
		//
		// resolver.Merge() will:
		// 1. execute <parallelism> number calls for each city (to reduce latency)
		// 2. always return within 5 seconds
		// 3. return as soon as there is at least one response from each city
		// 4. cancel any requests in progress upon returning
		//
		// All without any knowledge of our weather implementation.  Not too bad for
		// a language without generics!
		//
		resolver := par.Requests(requests).WithRedundancy(redundancy).WithConcurrency(10)
		return responses, resolver.Do()
	}()

	// extract and convert results
	var weathers []Weather
	var dedup map[City]City
	for weather := range responses {
		if _, ok := dedup[weather.City]; !ok {
			dedup[weather.City] = weather.City
			weathers = append(weathers, weather)
		}
	}

	return weathers, err
}

func findWeatherForCityFunc(city City, results chan Weather) func(context.Context) error {
	return func(ctx context.Context) error {
		forecast, err := openweathermap.WithContext(ctx, "").ByCityName(string(city))
		if err != nil {
			return err
		}

		results <- Weather{
			City:        City(forecast.Name),
			Temperature: int(forecast.Temperature.Temp),
			Highs:       int(forecast.Temperature.TempMax),
			Lows:        int(forecast.Temperature.TempMin),
		}

		return nil
	}
}
