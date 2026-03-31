package rules

import (
	"parser/utils"
)

type RepeatRule struct {
	Count           int
	Hours           int
	Course          utils.Course
	MajorInternship string
}

func NewRepeatRule(count int, hours int, course utils.Course, internship string) *RepeatRule {
	return &RepeatRule{
		Count:           count,
		Hours:           hours,
		Course:          course,
		MajorInternship: internship,
	}
}

func NewCourseRepeatRule(course utils.Course) *RepeatRule {
	return &RepeatRule{
		Count:  1,
		Course: course,
	}
}

func NewInternshipRepeatRule(internship string) *RepeatRule {
	return &RepeatRule{
		Count:           1,
		MajorInternship: internship,
	}
}
