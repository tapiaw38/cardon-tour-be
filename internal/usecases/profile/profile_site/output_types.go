package profilesite

import domain "github.com/tapiaw38/cardon-tour-be/internal/domain/profile"

type (
	ProfileSiteOutputData struct {
		ProfileID string `json:"profile_id"`
		SiteID    string `json:"site_id"`
	}
)

func toProfileSiteOutputData(profileSite domain.ProfileSite) ProfileSiteOutputData {
	return ProfileSiteOutputData{
		ProfileID: profileSite.ProfileID,
		SiteID:    profileSite.SiteID,
	}
}
