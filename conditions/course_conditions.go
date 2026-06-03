package conditions

import (
	"encoding/json"
	"parser/constants"
)

type CourseCondition struct {
	Course   constants.Course `json:"course,omitempty"`
	MinGrade constants.Grade  `json:"min_grade,omitempty"`
}

func (c *CourseCondition) MarshalJSON() ([]byte, error) {
	type Alias CourseCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "course",
		Alias: (*Alias)(c),
	})
}

func NewCourseCondition(prefix, number, grade string) *CourseCondition {
	return &CourseCondition{
		Course: constants.Course{
			Prefix: prefix,
			Number: number,
		},
		MinGrade: constants.Grade(grade),
	}
}

func (c *CourseCondition) Fulfils(userInfo constants.UserInfo) (bool, error) {
	return false, nil
}

func (c *CourseCondition) AppendGrade(grade constants.Grade) {
	c.MinGrade = grade
}

type CoreCondition struct {
	CoreNumber    string `json:"core_number,omitempty"`
	CoreTitle     string `json:"core_title,omitempty"`
	SemesterHours int    `json:"semester_hours,omitempty"`
}

func (c *CoreCondition) MarshalJSON() ([]byte, error) {
	type Alias CoreCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "core",
		Alias: (*Alias)(c),
	})
}

func (c *CoreCondition) Fulfils(userInfo constants.UserInfo) (bool, error) {
	return false, nil
}

func NewCoreCondition(courseNumber, coreTitle string) *CoreCondition {
	return &CoreCondition{
		CoreNumber: courseNumber,
		CoreTitle:  coreTitle,
	}
}

func NewCoreConditionWithSemesterHours(courseNumber, coreTitle string, semesterHours int) *CoreCondition {
	return &CoreCondition{
		CoreNumber:    courseNumber,
		SemesterHours: semesterHours,
	}
}

type CreditHoursCondition struct {
	Hours int `json:"hours"`
}

func (c *CreditHoursCondition) MarshalJSON() ([]byte, error) {
	type Alias CreditHoursCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "credit_hours",
		Alias: (*Alias)(c),
	})
}

func NewCreditHoursCondition(hours int) *CreditHoursCondition {
	return &CreditHoursCondition{
		Hours: hours,
	}
}

func (c *CreditHoursCondition) Fulfils(userInfo constants.UserInfo) (bool, error) {
	return false, nil
}

type CreditHoursFromCondition struct {
	Hours   int                `json:"hours,omitempty"`
	Courses []constants.Course `json:"courses,omitempty"`
}

func (c *CreditHoursFromCondition) MarshalJSON() ([]byte, error) {
	type Alias CreditHoursFromCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "credit_hours_from",
		Alias: (*Alias)(c),
	})
}

func NewCreditHoursFromCondition(hours int, courses []constants.Course) *CreditHoursFromCondition {
	return &CreditHoursFromCondition{
		Hours:   hours,
		Courses: courses,
	}
}

func (c *CreditHoursFromCondition) Fulfils(userInfo constants.UserInfo) (bool, error) {
	return false, nil
}

type UpperDivisionCoursesCondition struct {
	Hours  int    `json:"hours,omitempty"`
	Count  int    `json:"count,omitempty"`
	Prefix string `json:"prefix,omitempty"`
}

func (c *UpperDivisionCoursesCondition) MarshalJSON() ([]byte, error) {
	type Alias UpperDivisionCoursesCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "upper_division_courses",
		Alias: (*Alias)(c),
	})
}

func NewUpperDivisionCreditHoursCondition(hours int, prefix string) *UpperDivisionCoursesCondition {
	return &UpperDivisionCoursesCondition{
		Hours:  hours,
		Prefix: prefix,
	}
}

func NewUpperDivisionCountCondition(count int, prefix string) *UpperDivisionCoursesCondition {
	return &UpperDivisionCoursesCondition{
		Count:  count,
		Prefix: prefix,
	}
}

func (c *UpperDivisionCoursesCondition) Fulfils(userInfo constants.UserInfo) (bool, error) {
	return false, nil
}

type ResearchCondition struct {
	Hours       int                   `json:"hours,omitempty"`
	DegreeLevel constants.DegreeLevel `json:"degree_level,omitempty"`
}

func (c *ResearchCondition) MarshalJSON() ([]byte, error) {
	type Alias ResearchCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "research",
		Alias: (*Alias)(c),
	})
}

func NewResearchCondition(hours int, degreeLevel constants.DegreeLevel) *ResearchCondition {
	return &ResearchCondition{
		Hours:       hours,
		DegreeLevel: degreeLevel,
	}
}

func (c *ResearchCondition) Fulfils(userInfo constants.UserInfo) (bool, error) {
	return false, nil
}

type NCoursesCondition struct {
	N       int                `json:"n,omitempty"`
	Courses []constants.Course `json:"courses,omitempty"`
}

func (c *NCoursesCondition) MarshalJSON() ([]byte, error) {
	type Alias NCoursesCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "n_courses",
		Alias: (*Alias)(c),
	})
}

func NewNCoursesCondition(n int, courses []constants.Course) *NCoursesCondition {
	return &NCoursesCondition{
		N:       n,
		Courses: courses,
	}
}

func (c *NCoursesCondition) Fulfils(userInfo constants.UserInfo) (bool, error) {
	return false, nil
}
