package rules

import (
	"parser/utils"
)

type RepeatRule struct {
	Count           int
	Hours           int
	Courses         []utils.Course
	MajorInternship string
}

func (r *RepeatRule) isRule() bool {
	return true
}

func NewRepeatRule(count int, hours int, courses []utils.Course, internship string) *RepeatRule {
	return &RepeatRule{
		Count:           count,
		Hours:           hours,
		Courses:         courses,
		MajorInternship: internship,
	}
}

func NewCourseRepeatRule(course []utils.Course) *RepeatRule {
	return &RepeatRule{
		Count:   1,
		Courses: course,
	}
}

func NewInternshipRepeatRule(internship string) *RepeatRule {
	return &RepeatRule{
		Count:           1,
		MajorInternship: internship,
	}
}

func (r *RepeatRule) IsRule() bool {
	return true
}

type GpaRepeatRule struct {
	Course       utils.Course
	AcademicPlan string
}

func NewGpaRepeatRule(course utils.Course) *GpaRepeatRule {
	return &GpaRepeatRule{
		Course: course,
	}
}

func (r *GpaRepeatRule) isRule() bool {
	return true
}
