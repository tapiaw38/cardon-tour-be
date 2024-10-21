package site_business_type

import domain "github.com/tapiaw38/cardon-tour-be/internal/domain/site"

type SiteBusinessTypeOutputData struct {
	SiteID         string `json:"site_id"`
	BusinessTypeID string `json:"business_type_id"`
}

func toSiteBusinessTypeOutputData(siteBusinessType *domain.SiteBusinessType) SiteBusinessTypeOutputData {
	return SiteBusinessTypeOutputData{
		SiteID:         siteBusinessType.SiteID,
		BusinessTypeID: siteBusinessType.BusinessTypeID,
	}
}
