package domain

type (
	Country struct {
		ID   string
		Name string
		Code string
	}

	Province struct {
		ID        string
		Name      string
		Code      string
		Latitude  float64
		Longitude float64
		CountryID string
		Country   *Country
	}

	City struct {
		ID         string
		Name       string
		Code       string
		Latitude   float64
		Longitude  float64
		ProvinceID string
		Province   *Province
	}
)
