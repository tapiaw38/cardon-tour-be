package domain

type (
	ProfileType struct {
		ID   string
		Name string
	}

	Profile struct {
		ID            string
		UserID        string
		ProfileTypeID string
		ProfileType   *ProfileType
	}
)
