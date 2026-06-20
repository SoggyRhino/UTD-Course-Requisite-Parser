package conditions

import (
	"encoding/json"
	"fmt"
	"parser/objects/constants"
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

func (g *GPACondition) Fulfils(userInfo constants.UserInfo, _ bool) *constants.Evaluation {
	if userInfo.GPA >= float64(g.GPA) {
		return &constants.Evaluation{
			Name:    "GPA",
			Status:  constants.StatusPass,
			Summary: fmt.Sprintf("GPA is %.2f (requires %.2f)", userInfo.GPA, g.GPA),
		}
	}
	return &constants.Evaluation{
		Name:    "GPA",
		Status:  constants.StatusDefiniteFail,
		Summary: fmt.Sprintf("GPA is %.2f but requires %.2f", userInfo.GPA, g.GPA),
	}
}
