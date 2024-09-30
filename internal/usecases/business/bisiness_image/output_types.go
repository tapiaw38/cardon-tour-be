package bisiness_image

import domain "github.com/tapiaw38/cardon-tour-be/internal/domain/business"

type BusinessImageOutputData struct {
	ID         string `json:"id"`
	URL        string `json:"url"`
	BusinessID string `json:"business_id"`
}

func toBusinessImageOutputData(businessImages domain.BusinessImage) BusinessImageOutputData {
	return BusinessImageOutputData{
		ID:         businessImages.ID,
		URL:        businessImages.URL,
		BusinessID: businessImages.BusinessID,
	}
}
