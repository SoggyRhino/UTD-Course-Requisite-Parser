package rules

import "parser/utils"

type SameAsRule struct {
	Courses []utils.Course
}

func NewSameAsRule(courses []utils.Course) *SameAsRule {
	return &SameAsRule{Courses: courses}
}
