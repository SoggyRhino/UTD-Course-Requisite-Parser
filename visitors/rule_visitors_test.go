package visitors

import (
	"parser/parser"
	"parser/rules"
	"parser/utils"
	"testing"
)

func TestVisitRepeatRule(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result any
	}{
		"Single course repeat restriction": {
			Input: "ACCT 2301 Repeat Restriction",
			Result: []*rules.RepeatRule{rules.NewCourseRepeatRule(utils.Course{
				Prefix: "ACCT",
				Number: "2301",
			})},
		},
		"Dual prefix course repeat restriction": {
			Input: "CE/CS 2305 Repeat Restriction",
			Result: []*rules.RepeatRule{
				rules.NewCourseRepeatRule(utils.Course{
					Prefix: "CE",
					Number: "2305",
				}),
				rules.NewCourseRepeatRule(utils.Course{
					Prefix: "CS",
					Number: "2305",
				}),
			},
		},
		"Internship repeat restriction": {
			Input:  "BBSC Internship Repeat Restriction",
			Result: rules.NewInternshipRepeatRule("BBSC"),
		},
		"Bare repeat restriction": {
			Input:  "Repeat Restriction",
			Result: rules.NewRepeatRule(0, 0, utils.Course{}, ""),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[any](t, tc.Input, rule((*parser.RequirementsParser).Repeat_rule), tc.Result)
		})
	}
}
