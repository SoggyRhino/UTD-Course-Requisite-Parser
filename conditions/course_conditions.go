package conditions

import "parser/utils"

type CourseCondition struct {
	Course   utils.Course
	MinGrade utils.Grade
}

func NewCourseCondition(prefix, number, grade string) *CourseCondition {
	return &CourseCondition{
		Course: utils.Course{
			Prefix: prefix,
			Number: number,
		},
		MinGrade: utils.Grade(grade),
	}
}

func (c *CourseCondition) Fulfils(userInfo utils.UserInfo) (bool, error) {
	return false, nil
}

func (c *CourseCondition) AppendGrade(grade utils.Grade) {
	c.MinGrade = grade
}

type CoreCondition struct {
	CoreNumber    string
	CoreTitle     string
	SemesterHours int
}

func (c *CoreCondition) Fulfils(userInfo utils.UserInfo) (bool, error) {
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
	Hours int
}

func NewCreditHoursCondition(hours int) *CreditHoursCondition {
	return &CreditHoursCondition{
		Hours: hours,
	}
}

func (c *CreditHoursCondition) Fulfils(userInfo utils.UserInfo) (bool, error) {
	return false, nil
}

type CreditHoursFromCondition struct {
	Hours   int
	Courses []utils.Course
}

func NewCreditHoursFromCondition(hours int, courses []utils.Course) *CreditHoursFromCondition {
	return &CreditHoursFromCondition{
		Hours:   hours,
		Courses: courses,
	}
}

func (c *CreditHoursFromCondition) Fulfils(userInfo utils.UserInfo) (bool, error) {
	return false, nil
}

type UpperDivisionCoursesCondition struct {
	Hours  int
	Count  int
	Prefix string
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

func (c *UpperDivisionCoursesCondition) Fulfils(userInfo utils.UserInfo) (bool, error) {
	return false, nil
}

type ResearchCondition struct {
	Hours       int
	DegreeLevel utils.DegreeLevel
}

func NewResearchCondition(hours int, degreeLevel utils.DegreeLevel) *ResearchCondition {
	return &ResearchCondition{
		Hours:       hours,
		DegreeLevel: degreeLevel,
	}
}

func (c *ResearchCondition) Fulfils(userInfo utils.UserInfo) (bool, error) {
	return false, nil
}

type NCoursesCondition struct {
	N       int
	Courses []utils.Course
}

func NewNCoursesCondition(n int, courses []utils.Course) *NCoursesCondition {
	return &NCoursesCondition{
		N:       n,
		Courses: courses,
	}
}

func (c *NCoursesCondition) Fulfils(userInfo utils.UserInfo) (bool, error) {
	return false, nil
}
