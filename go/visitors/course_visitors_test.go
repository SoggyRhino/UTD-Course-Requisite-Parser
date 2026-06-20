package visitors

import (
	conditions2 "parser/objects/conditions"
	"parser/parser"
	"testing"
)

func TestVisitCourse(t *testing.T) {

	testCases := map[string]struct {
		Input  string
		Result conditions2.Condition
	}{
		"Simple": {
			Input:  "ACN 6340",
			Result: conditions2.NewCourseCondition("ACN", "6340", ""),
		},
		"Simple: Expand Prefix": {
			Input: "CS/CE 6340",
			Result: conditions2.NewOrCondition(
				conditions2.NewCourseCondition("CS", "6340", ""),
				conditions2.NewCourseCondition("CE", "6340", ""),
			),
		},
		"Paren": {
			Input:  "(ACN 6340)",
			Result: conditions2.NewCourseCondition("ACN", "6340", ""),
		},
		"Cross-Listed": {
			Input: "ACN 6340/HCS 6340",
			Result: conditions2.NewOrCondition(
				conditions2.NewCourseCondition("ACN", "6340", ""),
				conditions2.NewCourseCondition("HCS", "6340", ""),
			),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions2.Condition](t, tc.Input, rule((*parser.RequirementsParser).Course), tc.Result)
		})
	}
}

func TestVisitCourseList(t *testing.T) {

	testCases := map[string]struct {
		Input  string
		Result conditions2.Condition
	}{
		"Full Course List": {
			Input: "ACN 6340 or HCS 6340",
			Result: conditions2.NewOrCondition(
				conditions2.NewCourseCondition("ACN", "6340", ""),
				conditions2.NewCourseCondition("HCS", "6340", ""),
			),
		},
		"Shorthand Course List": {
			Input: "ACN 6340 or 7340",
			Result: conditions2.NewOrCondition(
				conditions2.NewCourseCondition("ACN", "6340", ""),
				conditions2.NewCourseCondition("ACN", "7340", ""),
			),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions2.Condition](t, tc.Input, rule((*parser.RequirementsParser).Course_list), tc.Result)
		})
	}
}

func TestVisitGradeCourseList(t *testing.T) {

	testCases := map[string]struct {
		Input  string
		Result conditions2.Condition
	}{
		"Either": {
			Input: "either MATH 2414 or in MATH 2418 or MATH 2419",
			Result: conditions2.NewOrCondition(
				conditions2.NewCourseCondition("MATH", "2414", ""),
				conditions2.NewCourseCondition("MATH", "2418", ""),
				conditions2.NewCourseCondition("MATH", "2419", ""),
			),
		},
		"Parenthesis": {
			Input: "(CE 2336 or CS 2336 or CS 2337)",
			Result: conditions2.NewOrCondition(
				conditions2.NewCourseCondition("CE", "2336", ""),
				conditions2.NewCourseCondition("CS", "2336", ""),
				conditions2.NewCourseCondition("CS", "2337", ""),
			),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions2.Condition](t, tc.Input, rule((*parser.RequirementsParser).Grade_course_list), tc.Result)
		})
	}
}
