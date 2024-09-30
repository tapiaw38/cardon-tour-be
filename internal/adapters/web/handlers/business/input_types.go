package business

import (
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/business"
	"time"
)

type BusinessInput struct {
	ProfileID      string    `json:"profile_id"`
	BusinessTypeID string    `json:"business_type_id"`
	SiteID         string    `json:"site_id"`
	Name           string    `json:"name"`
	PhoneNumber    string    `json:"phone_number"`
	Email          string    `json:"email"`
	Description    string    `json:"description"`
	Address        string    `json:"address"`
	Active         bool      `json:"active"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	CreatedAt      time.Time `json:"created_at"`
}

func toBusinessInput(input BusinessInput) domain.Business {
	return domain.Business{
		ProfileID:      input.ProfileID,
		BusinessTypeID: input.BusinessTypeID,
		SiteID:         input.SiteID,
		Name:           input.Name,
		PhoneNumber:    input.PhoneNumber,
		Email:          input.Email,
		Description:    input.Description,
		Address:        input.Address,
		Active:         input.Active,
		Latitude:       input.Latitude,
		Longitude:      input.Longitude,
		CreatedAt:      input.CreatedAt,
	}
}
