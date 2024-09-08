package profilesite

type inputProfileSite struct {
	ProfileID string `json:"profile_id" binding:"required"`
	SiteID    string `json:"site_id" binding:"required"`
}
