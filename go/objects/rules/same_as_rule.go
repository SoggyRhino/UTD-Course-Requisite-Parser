package rules

import (
	"encoding/json"
	"parser/objects/constants"
)

type SameAsRule struct {
	Courses []constants.Course `json:"courses,omitempty"`
}

func (r *SameAsRule) Fulfils(_ constants.UserInfo) *constants.Evaluation {
	return nil
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
