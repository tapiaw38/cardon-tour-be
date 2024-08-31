package domain

import domain "github.com/tapiaw38/cardon-tour-be/internal/domain/location"

type (
	Site struct {
		ID                string
		Slug              string
		Name              string
		Description       string
		ImageURL          string
		CityID            string
		City              *domain.City
		BusinessTypeSlugs []string
	}
)
