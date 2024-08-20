package businesstype

import domain "github.com/tapiaw38/cardon-tour-be/internal/domain/business"

type BusinessTypeOutputData struct {
	ID          string `json:"id"`
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

func toBusinessTypeOutputData(businessType *domain.BusinessType) BusinessTypeOutputData {
	return BusinessTypeOutputData{
		ID:          businessType.ID,
		Slug:        businessType.Slug,
		Name:        businessType.Name,
		Color:       businessType.Color,
		Description: businessType.Description,
		ImageURL:    businessType.ImageURL,
	}
}
