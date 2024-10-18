package site

import domain_site "github.com/tapiaw38/cardon-tour-be/internal/domain/site"

type SiteInput struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	CityID      string `json:"city_id"`
	IsPromoted  bool   `json:"is_promoted"`
}

func toSiteInput(input SiteInput) domain_site.Site {
	return domain_site.Site{
		Name:        input.Name,
		Slug:        input.Slug,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		CityID:      input.CityID,
		IsPromoted:  input.IsPromoted,
	}
}
