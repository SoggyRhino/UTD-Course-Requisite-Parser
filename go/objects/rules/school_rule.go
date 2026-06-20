package rules

import (
	"encoding/json"
	"fmt"
	"parser/objects/constants"
	"strings"
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

func (r *SchoolRule) Fulfils(userInfo constants.UserInfo) *constants.Evaluation {
	for _, school := range r.Schools {
		if userInfo.School == school {
			return &constants.Evaluation{
				Name:    "School Rule",
				Status:  constants.StatusPass,
				Summary: fmt.Sprintf("Student is in required school: %s", school),
			}
		}
	}

	return &constants.Evaluation{
		Name:    "School Rule",
		Status:  constants.StatusDefiniteFail,
		Summary: fmt.Sprintf("Student is not in any of the required schools: %s", strings.Join(r.Schools, ", ")),
	}
}

func NewSchoolRule(schools []string) *SchoolRule {
	return &SchoolRule{
		Schools: schools,
	}
}
