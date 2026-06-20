package rules

import (
	"encoding/json"
	"parser/objects/constants"
)

type SchoolRule struct {
	Schools []string `json:"schools,omitempty"`
}

func (r *SchoolRule) MarshalJSON() ([]byte, error) {
	type Alias SchoolRule
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "school",
		Alias: (*Alias)(r),
	})
}

func (r *SchoolRule) Fulfils(userInfo constants.UserInfo) constants.Evaluation {
	return constants.Evaluation{
		Name:    "School Rule",
		Status:  constants.StatusDefiniteFail,
		Summary: "Not implemented",
	}
}

func NewSchoolRule(schools []string) *SchoolRule {
	return &SchoolRule{
		Schools: schools,
	}
}
