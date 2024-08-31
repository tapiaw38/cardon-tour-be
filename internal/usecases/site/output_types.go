package site

import domain "github.com/tapiaw38/cardon-tour-be/internal/domain/site"

type (
	CityOutputData struct {
		Name      string  `json:"name"`
		Code      string  `json:"code"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	SiteOutputData struct {
		ID                string         `json:"id"`
		Slug              string         `json:"slug"`
		Name              string         `json:"name"`
		Description       string         `json:"description"`
		ImageURL          string         `json:"image_url"`
		City              CityOutputData `json:"city"`
		BusinessTypeSlugs []string       `json:"business_type_slugs,omitempty"`
	}
)

func toSiteOutputData(site *domain.Site) SiteOutputData {
	return SiteOutputData{
		ID:          site.ID,
		Slug:        site.Slug,
		Name:        site.Name,
		Description: site.Description,
		ImageURL:    site.ImageURL,
		City: CityOutputData{
			Name:      site.City.Name,
			Code:      site.City.Code,
			Latitude:  site.City.Latitude,
			Longitude: site.City.Longitude,
		},
		BusinessTypeSlugs: site.BusinessTypeSlugs,
	}
}
