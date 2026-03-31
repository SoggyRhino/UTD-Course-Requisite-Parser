package conditions

import "parser/utils"

type GPAType float64
type GPACondition struct {
	GPA    GPAType
	Degree string
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

func (g *GPACondition) Fulfils(userInfo utils.UserInfo) (bool, error) {
	return false, nil
}
