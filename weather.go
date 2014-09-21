package main

import (
	"code.google.com/p/go.net/context"
	"errors"
	"github.com/savaki/merge"
	"math/rand"
	"time"
)

type State string

const (
	CA State = "CA"
)

type City string

type Weather struct {
	City        City
	Temperature int
	Highs       int
	Lows        int
}

func FindCities(state State) []City {
	return []City{
		"San Francisco",
		"Oakland",
		"San Jose",
		"Berkeley",
	}
}

func FindWeatherAtFunc(city City, results chan Weather) func(context.Context) error {
	return func(ctx context.Context) error {
		// simulate a random delay in connection
		delay := time.Duration(rand.Int63n(1000)) * time.Millisecond
		timer := time.NewTimer(delay)
		defer timer.Stop()

		// sends the results after delay ms or returns if this request has been canceled
		select {
		case <-ctx.Done():
			return errors.New("timeout!")
		case <-timer.C:
			results <- fakeWeather(city)
		}

		return nil
	}
}

func findWeather(cities []City, parallelism int, timeout time.Duration) (<-chan Weather, error) {
	responses := make(chan Weather, len(cities)*parallelism)
	defer close(responses)

	requests := make(chan merge.RequestFunc, len(cities))
	for _, city := range cities {
		requests <- FindWeatherAtFunc(city, responses)
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
	resolver := merge.Requests(requests, timeout).WithParallelism(parallelism)
	err := resolver.Merge()

	return (<-chan Weather)(responses), err
}

func FindWeather(state State) (map[City]Weather, error) {
	// Unlike the original poster who described a system that made one query per
	// city, we're going to make two queries per city.  While this increases our
	// api consumption rate, we can always mitigate this through the use of a proxy
	// server like varnish which is anyways better suited to caching than we are.
	parallelism := 2
	timeout := 5 * time.Second

	// Delegate to a helper the creation of a channel for our requests as well as
	// a channel for the responses.
	responses, err := findWeather(FindCities(state), parallelism, timeout)
	if err != nil {
		return nil, err
	}

	// weathers now contains all the responses we're looking for
	// (and perhaps a few extra ;)
	weathers := map[City]Weather{}
	for result := range responses {
		weathers[result.City] = result
	}
	return weathers, nil
}
