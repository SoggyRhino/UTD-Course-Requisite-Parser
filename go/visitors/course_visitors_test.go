package visitors

import (
	"parser/objects/conditions"
	"parser/parser"
	"testing"
)

func TestVisitCourse(t *testing.T) {

	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"Simple": {
			Input:  "ACN 6340",
			Result: conditions.NewCourseCondition("ACN", "6340", ""),
		},
		"Simple: Expand Prefix": {
			Input: "CS/CE 6340",
			Result: conditions.NewOrCondition(
				conditions.NewCourseCondition("CS", "6340", ""),
				conditions.NewCourseCondition("CE", "6340", ""),
			),
		},
		"Paren": {
			Input:  "(ACN 6340)",
			Result: conditions.NewCourseCondition("ACN", "6340", ""),
		},
		"Cross-Listed": {
			Input: "ACN 6340/HCS 6340",
			Result: conditions.NewOrCondition(
				conditions.NewCourseCondition("ACN", "6340", ""),
				conditions.NewCourseCondition("HCS", "6340", ""),
			),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Course), tc.Result)
		})
	}
}

func TestVisitCourseList(t *testing.T) {

	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"Full Course List": {
			Input: "ACN 6340 or HCS 6340",
			Result: conditions.NewOrCondition(
				conditions.NewCourseCondition("ACN", "6340", ""),
				conditions.NewCourseCondition("HCS", "6340", ""),
			),
		},
		"Shorthand Course List": {
			Input: "ACN 6340 or 7340",
			Result: conditions.NewOrCondition(
				conditions.NewCourseCondition("ACN", "6340", ""),
				conditions.NewCourseCondition("ACN", "7340", ""),
			),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Course_list), tc.Result)
		})
	}
}

func TestVisitGradeCourseList(t *testing.T) {

	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"Either": {
			Input: "either MATH 2414 or in MATH 2418 or MATH 2419",
			Result: conditions.NewOrCondition(
				conditions.NewCourseCondition("MATH", "2414", ""),
				conditions.NewCourseCondition("MATH", "2418", ""),
				conditions.NewCourseCondition("MATH", "2419", ""),
			),
		},
		"Parenthesis": {
			Input: "(CE 2336 or CS 2336 or CS 2337)",
			Result: conditions.NewOrCondition(
				conditions.NewCourseCondition("CE", "2336", ""),
				conditions.NewCourseCondition("CS", "2336", ""),
				conditions.NewCourseCondition("CS", "2337", ""),
			),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Grade_course_list), tc.Result)
		})
	}
}
