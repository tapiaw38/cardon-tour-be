package profile

import (
	profile_domain "github.com/tapiaw38/cardon-tour-be/internal/domain/profile"
	site_domain "github.com/tapiaw38/cardon-tour-be/internal/domain/site"
)

type ProfileTypeOutputData struct {
	Name string `json:"name"`
}

type ProfileSitesOutputData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	CityID      string `json:"city_id"`
}

type ProfileOutputData struct {
	ID           string                   `json:"id"`
	UserID       string                   `json:"user_id"`
	ProfileType  ProfileTypeOutputData    `json:"profile_type"`
	ProfileSites []ProfileSitesOutputData `json:"profile_sites"`
}

func toProfileOutputData(profile *profile_domain.Profile) ProfileOutputData {
	profileSites := make([]ProfileSitesOutputData, len(profile.ProfileSites))
	for i, site := range profile.ProfileSites {
		profileSites[i] = toProfileSitesOutputData(&site)
	}

	return ProfileOutputData{
		ID:           profile.ID,
		UserID:       profile.UserID,
		ProfileType:  toProfileTypeOutputData(profile.ProfileType),
		ProfileSites: profileSites,
	}
}

func toProfileTypeOutputData(profileType *profile_domain.ProfileType) ProfileTypeOutputData {
	return ProfileTypeOutputData{
		Name: profileType.Name,
	}
}

func toProfileSitesOutputData(profileSite *site_domain.Site) ProfileSitesOutputData {
	return ProfileSitesOutputData{
		ID:          profileSite.ID,
		Name:        profileSite.Name,
		Description: profileSite.Description,
		ImageURL:    profileSite.ImageURL,
		CityID:      profileSite.CityID,
	}
}
