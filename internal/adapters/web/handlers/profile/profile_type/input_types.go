package profiletype

import domain "github.com/tapiaw38/cardon-tour-be/internal/domain/profile"

type inputProfileType struct {
	Name string `json:"name"`
}

func toProfileTypeInput(profileTypeInput inputProfileType) domain.ProfileType {
	return domain.ProfileType{
		Name: profileTypeInput.Name,
	}
}
