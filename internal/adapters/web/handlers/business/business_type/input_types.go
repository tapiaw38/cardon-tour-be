package businesstype

import domain "github.com/tapiaw38/cardon-tour-be/internal/domain/business"

type BusinessTypeInput struct {
	ID          string `json:"id"`
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

func toBusinessTypeInput(businessTypeInput BusinessTypeInput) domain.BusinessType {
	return domain.BusinessType{
		ID:          businessTypeInput.ID,
		Slug:        businessTypeInput.Slug,
		Name:        businessTypeInput.Name,
		Color:       businessTypeInput.Color,
		Description: businessTypeInput.Description,
		ImageURL:    businessTypeInput.ImageURL,
	}
}
