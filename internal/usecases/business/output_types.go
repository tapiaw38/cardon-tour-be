package business

import (
	"time"

	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/business"
)

type BusinessOutputData struct {
	ID             string    `json:"id"`
	ProfileID      string    `json:"profile_id"`
	BusinessTypeID string    `json:"business_type_id"`
	SiteID         string    `json:"site_id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	PhoneNumber    string    `json:"phone_number"`
	Email          string    `json:"email"`
	Address        string    `json:"address"`
	Active         bool      `json:"active"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	CreatedAt      time.Time `json:"created_at"`
}

func toBusinessOutputData(business *domain.Business) BusinessOutputData {
	return BusinessOutputData{
		ID:             business.ID,
		ProfileID:      business.ProfileID,
		BusinessTypeID: business.BusinessTypeID,
		SiteID:         business.SiteID,
		Name:           business.Name,
		Description:    business.Description,
		PhoneNumber:    business.PhoneNumber,
		Email:          business.Email,
		Address:        business.Address,
		Active:         business.Active,
		Latitude:       business.Latitude,
		Longitude:      business.Longitude,
		CreatedAt:      business.CreatedAt,
	}
}
