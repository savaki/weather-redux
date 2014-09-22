package main

import (
	"code.google.com/p/go.net/context"
	"github.com/savaki/cities-by-state"
)

func findCities(ctx context.Context, state string) ([]City, error) {
	// find our list of cities
	sites, err := citiesbystate.WithContext(ctx).ByState(state)
	if err != nil {
		return nil, err
	}

	// convert sites into a city array
	var cities []City
	for _, site := range sites {
		cities = append(cities, City(site.Name))
	}

	return cities, nil
}

func findWeatherForState(ctx context.Context, state string) ([]Weather, error) {
	cities, err := findCities(ctx, state)
	if err != nil {
		return nil, err
	}

	return findWeatherForCities(ctx, cities)
}
