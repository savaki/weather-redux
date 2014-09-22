package citiesbystate

type City struct {
	CountyName        string  `xml:"county_name"`
	Description       string  `xml:"description"`
	FeatClass         string  `xml:"feat_class"`
	FeatureId         int     `xml:"feature_id"`
	FipsClass         string  `xml:"fips_class"`
	FipsCountyCd      int     `xml:"fips_county_cd"`
	FullCountyName    string  `xml:"full_county_name"`
	Url               string  `xml:"url"`
	Name              string  `xml:"name"`
	PrimaryLatitude   float32 `xml:"primary_latitude"`
	PrimaryLongitude  float32 `xml:"primary_longitude"`
	StateAbbreviation string  `xml:"state_abbreviation"`
	StateName         string  `xml:"state_name"`
}

type Sites struct {
	Count  int    `xml:"count,attr"`
	Cities []City `xml:"site"`
}
