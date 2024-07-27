package domain

import (
	"time"

	location "github.com/tapiaw38/cardon-tour-be/internal/domain/location"
	profile "github.com/tapiaw38/cardon-tour-be/internal/domain/profile"
)

type (
	BusinessType struct {
		ID   string
		Name string
	}

	Business struct {
		ID             string
		BusinessTypeID string
		BusinessType   *BusinessType
		ProfileID      string
		Profile        *profile.Profile
		Name           string
		Phone          string
		Email          string
		Description    string
		Latitude       float64
		Longitude      float64
		CountryID      string
		Country        *location.Country
		ProvinceID     string
		Province       *location.Province
		CityID         string
		City           *location.City
		Images         []BusinessImage
		CreatedAt      time.Time
	}

	BusinessImage struct {
		ID  string
		URL string
	}
)
