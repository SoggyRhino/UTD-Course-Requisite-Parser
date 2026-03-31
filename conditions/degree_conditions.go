package conditions

import "parser/utils"

type DegreeLevel string

const (
	Undergraduate DegreeLevel = "Undergraduate"
	Graduate      DegreeLevel = "Graduate"
	PhD           DegreeLevel = "PhD"
)

type MajorCondition struct {
	Degree      string
	DegreeLevel DegreeLevel
	GradeLevel  GradeLevel
}

func NewMajorCondition(degree string) *MajorCondition {
	return &MajorCondition{
		Degree: degree,
	}
}

func NewMajorConditionWithGradeLevel(degree string, level GradeLevel) *MajorCondition {
	return &MajorCondition{
		Degree:     degree,
		GradeLevel: level,
	}
}

func NewMajorConditionWithDegreeLevel(degree string, level DegreeLevel) *MajorCondition {
	return &MajorCondition{
		Degree:      degree,
		DegreeLevel: level,
	}
}

func NewMajorConditionWithDegreeAndGradeLevel(degree string, level DegreeLevel, gradeLevel GradeLevel) *MajorCondition {
	return &MajorCondition{
		Degree:      degree,
		DegreeLevel: level,
		GradeLevel:  gradeLevel,
	}
}

func (m *MajorCondition) Fulfils(userInfo utils.UserInfo) (bool, error) {
	return false, nil
}

type DegreeCondition struct {
	Degree string
}

func NewDegreeCondition(degree string) *DegreeCondition {
	return &DegreeCondition{
		Degree: degree,
	}
}

func (d *DegreeCondition) Fulfils(userInfo utils.UserInfo) (bool, error) {
	return false, nil
}
