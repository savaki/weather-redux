package openweathermap

type Coord struct {
	Lon float32
	Lat float32
}

type Sys struct {
	Country string
	Sunrise int64
	Sunset  int64
}

type Weather struct {
	Id          int
	Main        string
	Description string
	Icon        string
}

type Temperature struct {
	Temp     float32
	Humidity int
	Pressure int
	TempMin  float32 `json:"temp_min"`
	TempMax  float32 `json:"temp_max"`
}

type Wind struct {
	Speed float32
	Deg   float32
}

type Rain struct {
	ThreeHour int
}

type Clouds struct {
	All int
}

type Forecast struct {
	Coord       Coord
	Sys         Sys
	Weather     []Weather
	Temperature Temperature `json:"main"`
	Wind        Wind
	Rain        Rain
	Clouds      Clouds
	Dt          int
	Id          int
	Name        string
	Cod         int
}
