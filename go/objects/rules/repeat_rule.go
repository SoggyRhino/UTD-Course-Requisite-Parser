package rules

import (
	"encoding/json"
	"parser/objects/constants"
)

type RepeatRule struct {
	Count           int                `json:"count,omitempty"`
	Hours           int                `json:"hours,omitempty"`
	Courses         []constants.Course `json:"courses,omitempty"`
	MajorInternship string             `json:"major_internship,omitempty"`
}

func (r *RepeatRule) isRule() bool {
	return true
}

func (r *RepeatRule) MarshalJSON() ([]byte, error) {
	type Alias RepeatRule
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "repeat",
		Alias: (*Alias)(r),
	})
}

func NewRepeatRule(count int, hours int, courses []constants.Course, internship string) *RepeatRule {
	return &RepeatRule{
		Count:           count,
		Hours:           hours,
		Courses:         courses,
		MajorInternship: internship,
	}
}

func NewCourseRepeatRule(course []constants.Course) *RepeatRule {
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

type GpaRepeatRule struct {
	Course       constants.Course `json:"course,omitempty"`
	AcademicPlan string           `json:"academic_plan,omitempty"`
}

func (r *GpaRepeatRule) MarshalJSON() ([]byte, error) {
	type Alias GpaRepeatRule
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "gpa_repeat",
		Alias: (*Alias)(r),
	})
}

func NewGpaRepeatRule(course constants.Course) *GpaRepeatRule {
	return &GpaRepeatRule{
		Course: course,
	}
}

func (r *GpaRepeatRule) isRule() bool {
	return true
}
