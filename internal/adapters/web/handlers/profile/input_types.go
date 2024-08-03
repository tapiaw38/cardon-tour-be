package profile

import domain "github.com/tapiaw38/cardon-tour-be/internal/domain/profile"

type inputProfile struct {
	UserID        string `json:"user_id"`
	ProfileTypeID string `json:"profile_type_id"`
}

func toProfileInput(profileInput inputProfile) domain.Profile {
	return domain.Profile{
		UserID:        profileInput.UserID,
		ProfileTypeID: profileInput.ProfileTypeID,
	}
}
