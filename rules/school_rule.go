package rules

import (
	"encoding/json"
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

func (r *SchoolRule) isRule() bool {
	return true
}

func NewSchoolRule(schools []string) *SchoolRule {
	return &SchoolRule{
		Schools: schools,
	}
}

func (r *SchoolRule) IsRule() bool {
	return true
}
