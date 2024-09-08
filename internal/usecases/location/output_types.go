package location

import domain "github.com/tapiaw38/cardon-tour-be/internal/domain/location"

type (
	ProvinceOutputData struct {
		ID          string  `json:"id"`
		Name        string  `json:"name"`
		Code        string  `json:"code"`
		Description string  `json:"description"`
		ImageURL    string  `json:"image_url"`
		Latitude    float64 `json:"latitude"`
		Longitude   float64 `json:"longitude"`
		CountryID   string  `json:"country_id"`
	}
)

func toProvinceOutputData(province *domain.Province) ProvinceOutputData {
	return ProvinceOutputData{
		ID:          province.ID,
		Name:        province.Name,
		Code:        province.Code,
		Description: province.Description,
		ImageURL:    province.ImageURL,
		Latitude:    province.Latitude,
		Longitude:   province.Longitude,
		CountryID:   province.CountryID,
	}
}
