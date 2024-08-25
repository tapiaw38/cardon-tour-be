package domain

import site "github.com/tapiaw38/cardon-tour-be/internal/domain/site"

type (
	ProfileType struct {
		ID   string
		Name string
	}

	Profile struct {
		ID             string
		UserID         string
		ProfileTypeID  string
		ProfileType    *ProfileType
		ProfileSitesID []string
		ProfileSites   []site.Site
	}
)
