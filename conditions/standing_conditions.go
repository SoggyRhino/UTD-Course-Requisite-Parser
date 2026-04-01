package conditions

import "parser/utils"

type GradeLevelCondition struct {
	GradeLevel utils.GradeLevel
	Degree     string
	School     string
}

func NewGradeLevelCondition(level utils.GradeLevel) *GradeLevelCondition {
	return &GradeLevelCondition{
		GradeLevel: level,
	}
}

func NewGradeLevelConditionWithDegree(level utils.GradeLevel, degree string) *GradeLevelCondition {
	return &GradeLevelCondition{
		GradeLevel: level,
		Degree:     degree,
	}
}

func (g *GradeLevelCondition) Fulfils(userInfo utils.UserInfo) (bool, error) {
	return false, nil
}

type GraduateStandingInCondition struct {
	Degree string
}

func NewGraduateStandingInCondition() *GraduateStandingInCondition {
	return &GraduateStandingInCondition{}
}

func NewGraduateStandingInConditionWithDegree(degree string) *GraduateStandingInCondition {
	return &GraduateStandingInCondition{
		Degree: degree}
}

func (g *GraduateStandingInCondition) Fulfils(userInfo utils.UserInfo) (bool, error) {
	return false, nil
}

type GenericStandingCondition struct {
	Standing utils.Standing
}

func NewGenericStandingCondition(standing utils.Standing) *GenericStandingCondition {
	return &GenericStandingCondition{
		Standing: standing,
	}
}

func (g *GenericStandingCondition) Fulfils(userInfo utils.UserInfo) (bool, error) {
	return false, nil
}
