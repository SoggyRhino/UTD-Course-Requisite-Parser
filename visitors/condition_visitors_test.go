package visitors

import (
	"parser/conditions"
	"parser/parser"
	"testing"
)

func TestVisitSimpleGradeCondition(t *testing.T) {

	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"Basic": {
			Input: "ACCT 2301 with a minimum grade of C",
			Result: conditions.NewCourseCondition(
				"ACCT", "2301", "C"),
		},
		"Or Better": {
			Input:  "ANGM 3370 with a grade of B or higher",
			Result: conditions.NewCourseCondition("ANGM", "3370", "B"),
		},
		"Short Hand Course": {
			Input: "CS/CE 2305 with a minimum grade of C",
			Result: conditions.NewOrCondition(
				conditions.NewCourseCondition("CS", "2305", "C"),
				conditions.NewCourseCondition("CE", "2305", "C"),
			),
		},
		"GPA": {
			Input:  "MATH 2413 with a grade greater than or equal to C- (1.67)",
			Result: conditions.NewCourseCondition("MATH", "2413", "C-"),
		},
		"Grade List": {
			Input: "(CE 2336 or CS 2336 or CS 2337) with a grade of C or better",
			Result: conditions.NewOrCondition(
				conditions.NewCourseCondition("CE", "2336", "C"),
				conditions.NewCourseCondition("CS", "2336", "C"),
				conditions.NewCourseCondition("CS", "2337", "C"),
			),
		},
		"At least Grade in": {
			Input: "a grade of at least a C- in MATH 1314 and MATH 1316",
			Result: conditions.NewAndCondition(
				conditions.NewCourseCondition("MATH", "1314", "C-"),
				conditions.NewCourseCondition("MATH", "1316", "C-"),
			),
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Grade_condition), tc.Result)
		})
	}

}
