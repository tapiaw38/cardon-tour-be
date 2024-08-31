package domain

import (
	"time"

	profile "github.com/tapiaw38/cardon-tour-be/internal/domain/profile"
	site "github.com/tapiaw38/cardon-tour-be/internal/domain/site"
)

type (
	BusinessType struct {
		ID          string
		Slug        string
		Name        string
		Color       string
		Description string
		ImageURL    string
	}

	Business struct {
		ID             string
		BusinessTypeID string
		BusinessType   *BusinessType
		SiteID         string
		Site           *site.Site
		ProfileID      string
		Profile        *profile.Profile
		Name           string
		PhoneNumber    string
		Email          string
		Address        string
		Description    string
		Latitude       float64
		Longitude      float64
		Active         bool
		Images         []BusinessImage
		CreatedAt      time.Time
	}

	BusinessImage struct {
		ID  string
		URL string
	}
)
