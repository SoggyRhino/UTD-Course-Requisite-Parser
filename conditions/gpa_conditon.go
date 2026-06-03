package conditions

import (
	"encoding/json"
	"parser/constants"
)

type GPAType float64
type GPACondition struct {
	GPA    GPAType `json:"gpa,omitempty"`
	Degree string  `json:"degree,omitempty"`
}

func (g *GPACondition) MarshalJSON() ([]byte, error) {
	type Alias GPACondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "gpa",
		Alias: (*Alias)(g),
	})
}

func NewGpaCondition(gpa float64) *GPACondition {
	return &GPACondition{
		GPA: GPAType(gpa),
	}
}

func NewGpaConditionWithDegree(gpa float64, degree string) *GPACondition {
	return &GPACondition{
		GPA:    GPAType(gpa),
		Degree: degree,
	}
}

func (g *GPACondition) Fulfils(userInfo constants.UserInfo) (bool, error) {
	return false, nil
}
