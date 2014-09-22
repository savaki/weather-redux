package openweathermap

import (
	"code.google.com/p/go.net/context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

const (
	baseUrl = "http://api.openweathermap.org/data/2.5/weather"
)

type WeatherService interface {
	ByCityName(name string) (*Forecast, error)
	ByCityId(id int) (*Forecast, error)
}

func New() WeatherService {
	return &weatherService{
		apiKey: "",
		ctx:    context.Background(),
	}
}

func WithApiKey(apiKey string) WeatherService {
	return &weatherService{
		apiKey: apiKey,
		ctx:    context.Background(),
	}
}

func WithContext(ctx context.Context, apiKey string) WeatherService {
	return &weatherService{
		apiKey: apiKey,
		ctx:    ctx,
	}
}

type weatherService struct {
	apiKey string
	ctx    context.Context
}

func (w *weatherService) ByCityName(city string) (*Forecast, error) {
	values := url.Values{}
	values.Add("q", city)

	return w.query(values)
}

func (w *weatherService) ByCityId(id int) (*Forecast, error) {
	values := url.Values{}
	values.Add("id", strconv.Itoa(id))

	return w.query(values)
}

func (w *weatherService) query(values url.Values) (*Forecast, error) {
	if w.apiKey != "" {
		values.Add("APPID", w.apiKey)
	}

	u, _ := url.Parse(baseUrl)
	u.RawQuery = values.Encode()
	req, _ := http.NewRequest("GET", u.String(), nil)

	var forecast Forecast
	err := httpDo(w.ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		return json.NewDecoder(resp.Body).Decode(&forecast)
	})

	return &forecast, err
}
