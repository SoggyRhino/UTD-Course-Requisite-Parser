package rules

import (
	"encoding/json"
	"fmt"
	"parser/objects/constants"
)

type RepeatRule struct {
	Count           int                `json:"count,omitempty"`
	Hours           int                `json:"hours,omitempty"`
	Courses         []constants.Course `json:"courses,omitempty"`
	MajorInternship string             `json:"major_internship,omitempty"`
}

func (r *RepeatRule) Fulfils(userInfo constants.UserInfo) constants.Evaluation {
	if r.MajorInternship != "" {
		return constants.Evaluation{
			Name:    "Repeat Rule (Major Internship)",
			Status:  constants.StatusPossibleFail,
			Summary: "Major internship repeat rule not implemented",
		}
	}

	takenCount := 0
	takenHours := 0

	for _, c := range r.Courses {
		if _, taken := userInfo.Taken[c]; taken {
			takenCount++
			takenHours += c.Hours()
		}

		for _, enrolled := range userInfo.CurrentEnrollment {
			if enrolled == c {
				takenCount++
				takenHours += c.Hours()
			}
		}
	}

	if r.Count > 0 && takenCount >= r.Count {
		return constants.Evaluation{
			Name:    "Repeat Rule",
			Status:  constants.StatusDefiniteFail,
			Summary: fmt.Sprintf("Repeated courses %d times (limit %d)", takenCount, r.Count),
		}
	}

	if r.Hours > 0 && takenHours >= r.Hours {
		return constants.Evaluation{
			Name:    "Repeat Rule",
			Status:  constants.StatusDefiniteFail,
			Summary: fmt.Sprintf("Repeated courses for %d hours (limit %d)", takenHours, r.Hours),
		}
	}

	return constants.Evaluation{
		Name:    "Repeat Rule",
		Status:  constants.StatusPass,
		Summary: "Repeat limits not exceeded",
	}
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

func (r *GpaRepeatRule) Fulfils(userInfo constants.UserInfo) constants.Evaluation {
	return constants.Evaluation{
		Name:    "GPA Repeat Rule",
		Status:  constants.StatusPossibleFail,
		Summary: "GPA repeat rule not implemented",
	}
}
