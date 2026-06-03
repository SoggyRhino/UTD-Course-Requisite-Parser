package conditions

import (
	"encoding/json"
	"parser/constants"
)

type GradeLevelCondition struct {
	GradeLevel constants.GradeLevel `json:"grade_level,omitempty"`
	Degree     string               `json:"degree,omitempty"`
	School     string               `json:"school,omitempty"`
}

func (g *GradeLevelCondition) MarshalJSON() ([]byte, error) {
	type Alias GradeLevelCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "grade_level",
		Alias: (*Alias)(g),
	})
}

func NewGradeLevelCondition(level constants.GradeLevel) *GradeLevelCondition {
	return &GradeLevelCondition{
		GradeLevel: level,
	}
}

func NewGradeLevelConditionWithDegree(level constants.GradeLevel, degree string) *GradeLevelCondition {
	return &GradeLevelCondition{
		GradeLevel: level,
		Degree:     degree,
	}
}

func (g *GradeLevelCondition) Fulfils(userInfo constants.UserInfo) (bool, error) {
	return false, nil
}

type GraduateStandingInCondition struct {
	Degree string `json:"degree,omitempty"`
}

func (g *GraduateStandingInCondition) MarshalJSON() ([]byte, error) {
	type Alias GraduateStandingInCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "graduate_standing_in",
		Alias: (*Alias)(g),
	})
}

func NewGraduateStandingInCondition() *GraduateStandingInCondition {
	return &GraduateStandingInCondition{}
}

func NewGraduateStandingInConditionWithDegree(degree string) *GraduateStandingInCondition {
	return &GraduateStandingInCondition{
		Degree: degree}
}

func (g *GraduateStandingInCondition) Fulfils(userInfo constants.UserInfo) (bool, error) {
	return false, nil
}

type GenericStandingCondition struct {
	Standing constants.Standing `json:"standing,omitempty"`
}

func (g *GenericStandingCondition) MarshalJSON() ([]byte, error) {
	type Alias GenericStandingCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "generic_standing",
		Alias: (*Alias)(g),
	})
}

func NewGenericStandingCondition(standing constants.Standing) *GenericStandingCondition {
	return &GenericStandingCondition{
		Standing: standing,
	}
}

func (g *GenericStandingCondition) Fulfils(userInfo constants.UserInfo) (bool, error) {
	return false, nil
}
