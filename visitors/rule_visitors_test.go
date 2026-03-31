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
			Result: rules.NewCourseRepeatRule([]utils.Course{
				{Prefix: "ACCT", Number: "2301"},
			}),
		},
		"Dual prefix course repeat restriction": {
			Input: "CE/CS 2305 Repeat Restriction",
			Result: rules.NewCourseRepeatRule(
				[]utils.Course{{Prefix: "CE", Number: "2305"}, {Prefix: "CS", Number: "2305"}},
			),
		},
		"Internship repeat restriction": {
			Input:  "BBSC Internship Repeat Restriction",
			Result: rules.NewInternshipRepeatRule("BBSC"),
		},
		"Bare repeat restriction": {
			Input:  "Repeat Restriction",
			Result: rules.NewRepeatRule(0, 0, []utils.Course{}, ""),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[any](t, tc.Input, rule((*parser.RequirementsParser).Repeat_rule), tc.Result)
		})
	}
}

func TestVisitRepeatLimitHoursRule(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result *rules.RepeatRule
	}{
		"Repeat max hours": {
			Input:  "Repeat Limit - ACCT 7323 may only be repeated for a maximum of 9 semester credit hours",
			Result: rules.NewRepeatRule(0, 9, []utils.Course{{Prefix: "ACCT", Number: "7323"}}, ""),
		},
		"Repeat hours max suffix": {
			Input:  "Repeat Limit - EEGR 6V88 may only be repeated for 9 semester credit hours maximum",
			Result: rules.NewRepeatRule(0, 9, []utils.Course{{Prefix: "EEGR", Number: "6V88"}}, ""),
		},
		"Course repeat max hours may": {
			Input:  "AHST 6336 Repeat Limit - This course may only be repeated for a maximum of 6 semester credit hours",
			Result: rules.NewRepeatRule(0, 6, []utils.Course{{Prefix: "AHST", Number: "6336"}}, ""),
		},
		"Course repeat max hours can": {
			Input:  "RHET 1302 Repeat Limit - This course can only be repeated for a maximum of 3 semester credit hours",
			Result: rules.NewRepeatRule(0, 3, []utils.Course{{Prefix: "RHET", Number: "1302"}}, ""),
		},
		"Combined repeat max hours": {
			Input:  "Repeat Limit - CS 4V96 and CS 4V98 combined may only be repeated for a maximum of 6 semester credit hours",
			Result: rules.NewRepeatRule(0, 6, []utils.Course{{Prefix: "CS", Number: "4V96"}, {Prefix: "CS", Number: "4V98"}}, ""),
		},
		"Course repeat limit": {
			Input:  "ANGM 2309 Repeat Limit",
			Result: rules.NewCourseRepeatRule([]utils.Course{{Prefix: "ANGM", Number: "2309"}}),
		},
		"Course repeat limit variable number": {
			Input:  "HLTH 4V01 Repeat Limit",
			Result: rules.NewCourseRepeatRule([]utils.Course{{Prefix: "HLTH", Number: "4V01"}}),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[*rules.RepeatRule](t, tc.Input, rule((*parser.RequirementsParser).Repeat_limit_hours_rule), tc.Result)
		})
	}
}

func TestVisitRepeatLimitTimesRule(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result *rules.RepeatRule
	}{
		"Repeat up to N times": {
			Input:  "Repeat Limit - UNIV 4074 may be repeated up to 3 times",
			Result: rules.NewRepeatRule(3, 0, []utils.Course{{Prefix: "UNIV", Number: "4074"}}, ""),
		},
		"Repeat max N times": {
			Input:  "Repeat Limit - OPRE 7051 may only be repeated a maximum of 6 times",
			Result: rules.NewRepeatRule(6, 0, []utils.Course{{Prefix: "OPRE", Number: "7051"}}, ""),
		},
		"Repeat max 3 times": {
			Input:  "Repeat Limit - UNIV 4076 may only be repeated a maximum of 3 times",
			Result: rules.NewRepeatRule(3, 0, []utils.Course{{Prefix: "UNIV", Number: "4076"}}, ""),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[*rules.RepeatRule](t, tc.Input, rule((*parser.RequirementsParser).Repeat_limit_times_rule), tc.Result)
		})
	}
}

func TestVisitGpaRepeatRule(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result *rules.GpaRepeatRule
	}{
		"Basic": {
			Input:  "GPA Repeat Restriction - MIS 6309",
			Result: rules.NewGpaRepeatRule(utils.Course{Prefix: "MIS", Number: "6309"}),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[*rules.GpaRepeatRule](t, tc.Input, rule((*parser.RequirementsParser).Gpa_repeate_rule), tc.Result)
		})
	}
}
