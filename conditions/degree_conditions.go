package conditions

import (
	"encoding/json"
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

func (m *MajorCondition) Fulfils(userInfo constants.UserInfo) (bool, error) {
	return false, nil
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

func (d *DegreeCondition) Fulfils(userInfo constants.UserInfo) (bool, error) {
	return false, nil
}
