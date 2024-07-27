package profile

import domain "github.com/tapiaw38/cardon-tour-be/internal/domain/profile"

type ProfileTypeOutputData struct {
	Name string `json:"name"`
}

type ProfileOutputData struct {
	ID          string                `json:"id"`
	UserID      string                `json:"user_id"`
	ProfileType ProfileTypeOutputData `json:"profile_type"`
}

func toProfileOutputData(profile *domain.Profile) ProfileOutputData {
	return ProfileOutputData{
		ID:     profile.ID,
		UserID: profile.UserID,
		ProfileType: ProfileTypeOutputData{
			Name: profile.ProfileType.Name,
		},
	}
}
