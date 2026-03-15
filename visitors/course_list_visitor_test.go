package visitors

import (
	"parser/conditions"
	"parser/parser"
	"testing"
)

/*
	 course_list
		: course (OR course)+                   # fullCourseList
		| course (OR COURSE_NUMBER)+            # shorthandCourseList
		;
*/
func TestVisitCourseList(t *testing.T) {

	testCases := map[string]struct {
		Input  string
		Result conditions.OrCondition
	}{
		"Full Course List": {
			Input: "ACN 6340 or HCS 6340",
			Result: conditions.OrCondition{
				Conditions: []conditions.Condition{
					conditions.CourseCondition{Course: conditions.NewCourse("ACN", "6340")},
					conditions.CourseCondition{Course: conditions.NewCourse("HCS", "6340")},
				},
			},
		},
		"Shorthand Course List": {
			Input: "ACN 6340 or 7340",
			Result: conditions.OrCondition{
				Conditions: []conditions.Condition{
					conditions.CourseCondition{Course: conditions.NewCourse("ACN", "6340")},
					conditions.CourseCondition{Course: conditions.NewCourse("ACN", "7340")},
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Course_list), tc.Result)
		})
	}
}
