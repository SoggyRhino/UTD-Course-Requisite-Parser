package visitors

import (
	"parser/conditions"
	"parser/parser"
	"testing"
)

/*
course
    : '(' course ')'						# parenCourse
    | course '/' course						# crossListedCourse
    | PREFIX ('/' PREFIX)* COURSE_NUMBER	# simpleCourse
    ;
*/

func TestVisitCourse(t *testing.T) {

	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"Simple": {
			Input: "ACN 6340",
			Result: conditions.CourseCondition{
				Course: conditions.NewCourse("ACN", "6340"),
			},
		},
		"Simple: Expand Prefix": {
			Input: "CS/CE 6340",
			Result: conditions.OrCondition{
				Conditions: []conditions.Condition{
					conditions.CourseCondition{Course: conditions.NewCourse("CS", "6340")},
					conditions.CourseCondition{Course: conditions.NewCourse("CE", "6340")},
				},
			},
		},
		"Paren": {
			Input: "(ACN 6340)",
			Result: conditions.CourseCondition{
				Course: conditions.NewCourse("ACN", "6340"),
			},
		},
		"Cross-Listed": {
			Input: "ACN 6340/HCS 6340",
			Result: conditions.OrCondition{
				Conditions: []conditions.Condition{
					conditions.CourseCondition{Course: conditions.NewCourse("ACN", "6340")},
					conditions.CourseCondition{Course: conditions.NewCourse("HCS", "6340")},
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Course), tc.Result)
		})
	}
}
