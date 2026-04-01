package rules

import (
	"parser/utils"
)

type CreditForRule struct {
	Courses CourseCollection
}

func NewCreditForRule(courses CourseCollection) *CreditForRule {
	return &CreditForRule{Courses: courses}
}

type CourseCollection interface {
	isCourseCollection() bool //todo change this actual work when evaluating
}

type AndCourseCollection struct {
	Courses []CourseCollection
}

func NewAndCourseCollection(col1, col2 CourseCollection) *AndCourseCollection {
	var flattenedCourses []CourseCollection

	if and1, isAnd1 := col1.(*AndCourseCollection); isAnd1 {
		flattenedCourses = append(flattenedCourses, and1.Courses...)
	} else {
		flattenedCourses = append(flattenedCourses, col1)
	}

	if and2, isAnd2 := col2.(*AndCourseCollection); isAnd2 {
		flattenedCourses = append(flattenedCourses, and2.Courses...)
	} else {
		flattenedCourses = append(flattenedCourses, col2)
	}

	return &AndCourseCollection{
		Courses: flattenedCourses,
	}
}

func (a *AndCourseCollection) isCourseCollection() bool {
	return true
}

type OrCourseCollection struct {
	Courses []CourseCollection
}

func NewOrCourseCollection(col1, col2 CourseCollection) *OrCourseCollection {
	var flattenedCourses []CourseCollection

	if or1, isOr1 := col1.(*OrCourseCollection); isOr1 {
		flattenedCourses = append(flattenedCourses, or1.Courses...)
	} else {
		flattenedCourses = append(flattenedCourses, col1)
	}

	if or2, isOr2 := col2.(*OrCourseCollection); isOr2 {
		flattenedCourses = append(flattenedCourses, or2.Courses...)
	} else {
		flattenedCourses = append(flattenedCourses, col2)
	}

	return &OrCourseCollection{
		Courses: flattenedCourses,
	}
}

func (o *OrCourseCollection) isCourseCollection() bool {
	return true
}

type SimpleCourseCollection struct {
	Course []utils.Course
}

func NewSimpleCourseCollection(course []utils.Course) *SimpleCourseCollection {
	return &SimpleCourseCollection{Course: course}
}

func (s *SimpleCourseCollection) isCourseCollection() bool {
	return true
}
