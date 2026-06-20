package rules

import (
	"encoding/json"
	"parser/objects/constants"
)

type SameAsRule struct {
	Courses []constants.Course `json:"courses,omitempty"`
}

func (r *SameAsRule) Fulfils(userInfo constants.UserInfo) constants.Evaluation {
	return constants.Evaluation{
		Name:    "Same As Rule",
		Status:  constants.StatusDefiniteFail,
		Summary: "Not implemented",
	}
}

func (r *SameAsRule) MarshalJSON() ([]byte, error) {
	type Alias SameAsRule
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "same_as",
		Alias: (*Alias)(r),
	})
}

func NewSameAsRule(courses []constants.Course) *SameAsRule {
	return &SameAsRule{Courses: courses}
}
