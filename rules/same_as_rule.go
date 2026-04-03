package rules

import "parser/utils"

type SameAsRule struct {
	Courses []utils.Course
}

func (r *SameAsRule) isRule() bool {
	panic("implement me")
}

func NewSameAsRule(courses []utils.Course) *SameAsRule {
	return &SameAsRule{Courses: courses}
}

func (r *SameAsRule) IsRule() bool {
	return true
}
