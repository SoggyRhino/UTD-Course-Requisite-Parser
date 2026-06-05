package conditions

import (
	"encoding/json"
	"fmt"
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

func (g *GradeLevelCondition) Fulfils(userInfo constants.UserInfo) *constants.Evaluation {
	if g.GradeLevel != "" && g.GradeLevel != constants.AnyGrade && userInfo.GradeLevel != g.GradeLevel {
		return &constants.Evaluation{
			Status:  constants.StatusDefiniteFail,
			Summary: fmt.Sprintf("Grade level is %s; requires %s", userInfo.GradeLevel, g.GradeLevel),
		}
	}
	// Note: School/Degree fields in GradeLevelCondition are currently unused in verification
	return &constants.Evaluation{
		Status:  constants.StatusPass,
		Summary: fmt.Sprintf("Grade level requirement %s satisfied", g.GradeLevel),
	}
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

func (g *GraduateStandingInCondition) Fulfils(userInfo constants.UserInfo) *constants.Evaluation {
	if userInfo.DegreeLevel != constants.Graduate && userInfo.DegreeLevel != constants.PhD {
		return &constants.Evaluation{
			Status:  constants.StatusDefiniteFail,
			Summary: fmt.Sprintf("Degree level is %s; requires Graduate standing", userInfo.DegreeLevel),
		}
	}
	return &constants.Evaluation{
		Status:  constants.StatusPass,
		Summary: "Graduate standing requirement satisfied",
	}
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

func (g *GenericStandingCondition) Fulfils(userInfo constants.UserInfo) *constants.Evaluation {
	if g.Standing != "" && userInfo.Standing != g.Standing {
		return &constants.Evaluation{
			Status:  constants.StatusDefiniteFail,
			Summary: fmt.Sprintf("Academic standing is %q; requires %q", userInfo.Standing, g.Standing),
		}
	}
	return &constants.Evaluation{
		Status:  constants.StatusPass,
		Summary: fmt.Sprintf("Academic standing %q satisfied", g.Standing),
	}
}
