package profile

import domain "github.com/tapiaw38/cardon-tour-be/internal/domain/profile"

func unmarshalProfile(
	id string,
	userID string,
	profileName string,
	profileSitesID []string,
) *domain.Profile {
	return &domain.Profile{
		ID:     id,
		UserID: userID,
		ProfileType: &domain.ProfileType{
			Name: profileName,
		},
		ProfileSitesID: profileSitesID,
	}
}
