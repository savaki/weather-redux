package citiesbystate

import (
	"code.google.com/p/go.net/context"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
)

const (
	baseUrl = "http://api.sba.gov/geodata/city_links_for_state_of/"
)

var states = map[string]string{
	"ALABAMA": "AL", "AL": "AL",
	"ALASKA": "AK", "AK": "AK",
	"ARIZONA": "AZ", "AZ": "AZ",
	"ARKANSAS": "AR", "AR": "AR",
	"CALIFORNIA": "CA", "CA": "CA",
	"COLORADO": "CO", "CO": "CO",
	"CONNECTICUT": "CT", "CT": "CT",
	"DELAWARE": "DE", "DE": "DE",
	"FLORIDA": "FL", "FL": "FL",
	"GEORGIA": "GA", "GA": "GA",
	"HAWAII": "HI", "HI": "HI",
	"IDAHO": "ID", "ID": "ID",
	"ILLINOIS": "IL", "IL": "IL",
	"INDIANA": "IN", "IN": "IN",
	"IOWA": "IA", "IA": "IA",
	"KANSAS": "KS", "KS": "KS",
	"KENTUCKY": "KY", "KY": "KY",
	"LOUISIANA": "LA", "LA": "LA",
	"MAINE": "ME", "ME": "ME",
	"MARYLAND": "MD", "MD": "MD",
	"MASSACHUSETTS": "MA", "MA": "MA",
	"MICHIGAN": "MI", "MI": "MI",
	"MINNESOTA": "MN", "MN": "MN",
	"MISSISSIPPI": "MS", "MS": "MS",
	"MISSOURI": "MO", "MO": "MO",
	"MONTANA": "MT", "MT": "MT",
	"NEBRASKA": "NE", "NE": "NE",
	"NEVADA": "NV", "NV": "NV",
	"NEW HAMPSHIRE": "NH", "NH": "NH",
	"NEW JERSEY": "NJ", "NJ": "NJ",
	"NEW MEXICO": "NM", "NM": "NM",
	"NEW YORK": "NY", "NY": "NY",
	"NORTH CAROLINA": "NC", "NC": "NC",
	"NORTH DAKOTA": "ND", "ND": "ND",
	"OHIO": "OH", "OH": "OH",
	"OKLAHOMA": "OK", "OK": "OK",
	"OREGON": "OR", "OR": "OR",
	"PENNSYLVANIA": "PA", "PA": "PA",
	"RHODE ISLAND": "RI", "RI": "RI",
	"SOUTH CAROLINA": "SC", "SC": "SC",
	"SOUTH DAKOTA": "SD", "SD": "SD",
	"TENNESSEE": "TN", "TN": "TN",
	"TEXAS": "TX", "TX": "TX",
	"UTAH": "UT", "UT": "UT",
	"VERMONT": "VT", "VT": "VT",
	"VIRGINIA": "VA", "VA": "VA",
	"WASHINGTON": "WA", "WA": "WA",
	"WEST VIRGINIA": "WV", "WV": "WV",
	"WISCONSIN": "WI", "WI": "WI",
	"WYOMING": "WY", "WY": "WY",
}

type CityService interface {
	ByState(name string) ([]City, error)
}

func New() CityService {
	return &cityService{
		ctx: context.Background(),
	}
}

func WithContext(ctx context.Context) CityService {
	return &cityService{
		ctx: ctx,
	}
}

type cityService struct {
	ctx context.Context
}

func (c *cityService) ByState(name string) ([]City, error) {
	upper := strings.ToUpper(name)
	abbrev, ok := states[upper]
	if !ok {
		return nil, fmt.Errorf("no state with name or abbreviation, %s", name)
	}

	req, _ := http.NewRequest("GET", baseUrl+abbrev, nil)
	var sites Sites
	err := httpDo(c.ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		return xml.NewDecoder(resp.Body).Decode(&sites)
	})

	return sites.Cities, err
}
