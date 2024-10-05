package site

import domain "github.com/tapiaw38/cardon-tour-be/internal/domain/site"

type (
	CityOutputData struct {
		Name       string  `json:"name"`
		Slug       string  `json:"slug"`
		Latitude   float64 `json:"latitude"`
		Longitude  float64 `json:"longitude"`
		ProvinceID string  `json:"province_id"`
	}

	SiteOutputData struct {
		ID             string         `json:"id"`
		Slug           string         `json:"slug"`
		Name           string         `json:"name"`
		Description    string         `json:"description"`
		ImageURL       string         `json:"image_url"`
		IsPromoted     bool           `json:"is_promoted"`
		City           CityOutputData `json:"city"`
		BusinessTypeID []string       `json:"business_type_ids"`
	}
)

func toSiteOutputData(site *domain.Site) SiteOutputData {
	return SiteOutputData{
		ID:          site.ID,
		Slug:        site.Slug,
		Name:        site.Name,
		Description: site.Description,
		ImageURL:    site.ImageURL,
		IsPromoted:  site.IsPromoted,
		City: CityOutputData{
			Name:       site.City.Name,
			Slug:       site.City.Slug,
			ProvinceID: site.City.ProvinceID,
			Latitude:   site.City.Latitude,
			Longitude:  site.City.Longitude,
		},
		BusinessTypeID: site.BusinessTypeID,
	}
}
