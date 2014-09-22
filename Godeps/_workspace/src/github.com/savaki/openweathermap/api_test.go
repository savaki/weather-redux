package openweathermap

import (
	"code.google.com/p/go.net/context"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLiveCall(t *testing.T) {
	var weatherService WeatherService
	var forecast *Forecast
	var err error

	Convey("Given a WeatherService with a custom context", t, func() {
		weatherService = WithContext(context.Background(), "")

		Convey("When I find the weather for San Francisco", func() {
			forecast, err = weatherService.ByCityName("San Francisco")

			Convey("Then I expect no errors", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then I expect the weather to have been returned", func() {
				So(forecast.Name, ShouldEqual, "San Francisco")
			})
		})
	})
}
