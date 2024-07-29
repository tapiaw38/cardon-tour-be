package profiletype

import (
	domain "github.com/tapiaw38/cardon-tour-be/internal/domain/profile"
)

type ProfileTypeOutputData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func toProfileTypeOutputData(profileType *domain.ProfileType) ProfileTypeOutputData {
	return ProfileTypeOutputData{
		ID:   profileType.ID,
		Name: profileType.Name,
	}
}
