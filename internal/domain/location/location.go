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
		CountryID string
		Country   *Country
	}

	City struct {
		ID         string
		Name       string
		Code       string
		ProvinceID string
		Province   *Province
		CountryID  string
		Country    *Country
	}
)
