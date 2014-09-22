package openweathermap

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestJsonUnmarshall(t *testing.T) {
	var content string
	var forecast Forecast
	var err error

	Convey("Given a json weather forecase", t, func() {
		content = `{
    "coord": {
        "lon": 139,
        "lat": 35
    },
    "sys": {
        "country": "JP",
        "sunrise": 1369769524,
        "sunset": 1369821049
    },
    "weather": [
        {
            "id": 804,
            "main": "clouds",
            "description": "overcast clouds",
            "icon": "04n"
        }
    ],
    "main": {
        "temp": 289.5,
        "humidity": 89,
        "pressure": 1013,
        "temp_min": 287.04,
        "temp_max": 292.04
    },
    "wind": {
        "speed": 7.31,
        "deg": 187.002
    },
    "rain": {
        "3h": 0
    },
    "clouds": {
        "all": 92
    },
    "dt": 1369824698,
    "id": 1851632,
    "name": "Shuzenji",
    "cod": 200
}`

		Convey("When I decode the json", func() {
			err = json.NewDecoder(strings.NewReader(content)).Decode(&forecast)

			Convey("Then I expect no errors", func() {
				So(err, ShouldBeNil)
			})

			Convey("And I expect Temperature to be extracted", func() {
				temperature := Temperature{
					Temp:     289.5,
					Humidity: 89,
					Pressure: 1013,
					TempMin:  287.04,
					TempMax:  292.04,
				}

				So(forecast.Temperature, ShouldResemble, temperature)
			})

			Convey("And I expect weather to be extracted", func() {
				weather := Weather{
					Id:          804,
					Main:        "clouds",
					Description: "overcast clouds",
					Icon:        "04n",
				}

				So(forecast.Weather, ShouldResemble, []Weather{weather})
			})

			Convey("And I expect coord to be extracted", func() {
				coord := Coord{
					Lon: 139,
					Lat: 35,
				}

				So(forecast.Coord, ShouldResemble, coord)
			})

			Convey("And I expect sys to be extracted", func() {
				sys := Sys{
					Country: "JP",
					Sunrise: 1369769524,
					Sunset:  1369821049,
				}

				So(forecast.Sys, ShouldResemble, sys)
			})

			Convey("And I expect wind to be extracted", func() {
				wind := Wind{
					Speed: 7.31,
					Deg:   187.002,
				}

				So(forecast.Wind, ShouldResemble, wind)
			})

			Convey("And I expect id to be set", func() {
				So(forecast.Id, ShouldEqual, 1851632)
			})

			Convey("And I expect name to be set", func() {
				So(forecast.Name, ShouldEqual, "Shuzenji")
			})

			Convey("And I expect cod to be set", func() {
				So(forecast.Cod, ShouldEqual, 200)
			})

			Convey("And I expect dt to be set", func() {
				So(forecast.Dt, ShouldEqual, 1369824698)
			})
		})
	})
}
