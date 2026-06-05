package conditions

import (
	"encoding/json"
	"fmt"
	"strings"

	"parser/constants"
)

type MajorCondition struct {
	Degree      string                `json:"degree,omitempty"`
	DegreeLevel constants.DegreeLevel `json:"degree_level,omitempty"`
	GradeLevel  constants.GradeLevel  `json:"grade_level,omitempty"`
}

func (m *MajorCondition) MarshalJSON() ([]byte, error) {
	type Alias MajorCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "major",
		Alias: (*Alias)(m),
	})
}

func NewMajorCondition(degree string) *MajorCondition {
	return &MajorCondition{
		Degree: degree,
	}
}

func NewMajorConditionWithGradeLevel(degree string, level constants.GradeLevel) *MajorCondition {
	return &MajorCondition{
		Degree:     degree,
		GradeLevel: level,
	}
}

func NewMajorConditionWithDegreeLevel(degree string, level constants.DegreeLevel) *MajorCondition {
	return &MajorCondition{
		Degree:      degree,
		DegreeLevel: level,
	}
}

func NewMajorConditionWithDegreeAndGradeLevel(degree string, level constants.DegreeLevel, gradeLevel constants.GradeLevel) *MajorCondition {
	return &MajorCondition{
		Degree:      degree,
		DegreeLevel: level,
		GradeLevel:  gradeLevel,
	}
}

func (m *MajorCondition) Fulfils(userInfo constants.UserInfo) *constants.Evaluation {
	if m.Degree != "" && !strings.Contains(strings.ToLower(userInfo.Major), strings.ToLower(m.Degree)) {
		return &constants.Evaluation{
			Status:  constants.StatusDefiniteFail,
			Summary: fmt.Sprintf("Major is %q; requires %q", userInfo.Major, m.Degree),
		}
	}

	if m.DegreeLevel != "" && m.DegreeLevel != constants.AnyDegree && userInfo.DegreeLevel != m.DegreeLevel {
		return &constants.Evaluation{
			Status:  constants.StatusDefiniteFail,
			Summary: fmt.Sprintf("Degree level is %q; requires %q", userInfo.DegreeLevel, m.DegreeLevel),
		}
	}

	if m.GradeLevel != "" && m.GradeLevel != constants.AnyGrade && userInfo.GradeLevel != m.GradeLevel {
		return &constants.Evaluation{
			Status:  constants.StatusDefiniteFail,
			Summary: fmt.Sprintf("Grade level is %q; requires %q", userInfo.GradeLevel, m.GradeLevel),
		}
	}

	return &constants.Evaluation{
		Status:  constants.StatusPass,
		Summary: "Major and level requirements satisfied",
	}
}

type DegreeCondition struct {
	Degree string `json:"degree,omitempty"`
}

func (d *DegreeCondition) MarshalJSON() ([]byte, error) {
	type Alias DegreeCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "degree",
		Alias: (*Alias)(d),
	})
}

func NewDegreeCondition(degree string) *DegreeCondition {
	return &DegreeCondition{
		Degree: degree,
	}
}

func (d *DegreeCondition) Fulfils(userInfo constants.UserInfo) *constants.Evaluation {
	if d.Degree != "" && !strings.Contains(strings.ToLower(userInfo.Major), strings.ToLower(d.Degree)) {
		return &constants.Evaluation{
			Status:  constants.StatusDefiniteFail,
			Summary: fmt.Sprintf("Major is %q; requires degree in %q", userInfo.Major, d.Degree),
		}
	}

	return &constants.Evaluation{
		Status:  constants.StatusPass,
		Summary: fmt.Sprintf("Degree requirement %q satisfied", d.Degree),
	}
}
