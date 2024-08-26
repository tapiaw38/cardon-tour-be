package site

import domain "github.com/tapiaw38/cardon-tour-be/internal/domain/site"

type SiteOutputData struct {
	ID                string   `json:"id"`
	Slug              string   `json:"slug"`
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	ImageURL          string   `json:"image_url"`
	BusinessTypeSlugs []string `json:"business_type_slugs,omitempty"`
}

func toSiteOutputData(site *domain.Site) SiteOutputData {
	return SiteOutputData{
		ID:                site.ID,
		Slug:              site.Slug,
		Name:              site.Name,
		Description:       site.Description,
		ImageURL:          site.ImageURL,
		BusinessTypeSlugs: site.BusinessTypeSlugs,
	}
}
