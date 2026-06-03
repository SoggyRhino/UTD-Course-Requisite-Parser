package conditions

import (
	"encoding/json"
	"parser/constants"
)

type ConcurrentEnrollmentCondition struct {
	Course constants.Course `json:"course,omitempty"`
}

func (c *ConcurrentEnrollmentCondition) MarshalJSON() ([]byte, error) {
	type Alias ConcurrentEnrollmentCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "concurrent_enrollment",
		Alias: (*Alias)(c),
	})
}

func NewConcurrentEnrollmentCondition(course constants.Course) *ConcurrentEnrollmentCondition {
	return &ConcurrentEnrollmentCondition{
		Course: course,
	}
}

func (c *ConcurrentEnrollmentCondition) Fulfils(userInfo constants.UserInfo) (bool, error) {
	return false, nil
}

type ExactSectionCondition struct {
	Course constants.Course `json:"course,omitempty"`
}

func (c *ExactSectionCondition) MarshalJSON() ([]byte, error) {
	type Alias ExactSectionCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "exact_section",
		Alias: (*Alias)(c),
	})
}

func NewExactSectionCondition(course constants.Course) *ExactSectionCondition {
	return &ExactSectionCondition{
		Course: course,
	}
}

func (c *ExactSectionCondition) Fulfils(userInfo constants.UserInfo) (bool, error) {
	return false, nil
}

type AnyPreviousMajorCourseCondition struct {
	Prefix string `json:"prefix,omitempty"`
}

func (c *AnyPreviousMajorCourseCondition) MarshalJSON() ([]byte, error) {
	type Alias AnyPreviousMajorCourseCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "any_previous_major_course",
		Alias: (*Alias)(c),
	})
}

func NewAnyPreviousMajorCourseCondition(prefix string) *AnyPreviousMajorCourseCondition {
	return &AnyPreviousMajorCourseCondition{
		Prefix: prefix,
	}
}

func (c *AnyPreviousMajorCourseCondition) Fulfils(userInfo constants.UserInfo) (bool, error) {
	return false, nil
}

type AcademicYearCondition struct {
	Plan  string `json:"plan,omitempty"`
	Equal bool   `json:"equal,omitempty"`
}

func (c *AcademicYearCondition) MarshalJSON() ([]byte, error) {
	type Alias AcademicYearCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "academic_year",
		Alias: (*Alias)(c),
	})
}

func NewAcademicYearCondition(plan string, equal bool) *AcademicYearCondition {
	return &AcademicYearCondition{
		Plan:  plan,
		Equal: equal,
	}
}

func (c *AcademicYearCondition) Fulfils(userInfo constants.UserInfo) (bool, error) {
	return false, nil
}
