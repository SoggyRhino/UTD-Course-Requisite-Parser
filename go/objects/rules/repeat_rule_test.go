package rules

import (
	"parser/objects/constants"
	"testing"
)

func TestRepeatRule_JSON(t *testing.T) {
	testCases := map[string]struct {
		rule *RepeatRule
	}{
		"Simple repeat": {
			rule: NewRepeatRule(2, 0, nil, ""),
		},
		"Repeat with hours": {
			rule: NewRepeatRule(0, 6, nil, ""),
		},
		"Repeat with courses": {
			rule: NewRepeatRule(1, 0, []constants.Course{{Prefix: "CS", Number: "1337"}}, ""),
		},
		"Repeat with internship": {
			rule: NewRepeatRule(1, 0, nil, "CS"),
		},
		"Course repeat rule": {
			rule: NewCourseRepeatRule([]constants.Course{{Prefix: "CS", Number: "1337"}}),
		},
		"Internship repeat rule": {
			rule: NewInternshipRepeatRule("CS"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assertJSONRoundTrip[RepeatRule](t, tc.rule, "repeat")
		})
	}
}

func TestGpaRepeatRule_JSON(t *testing.T) {
	testCases := map[string]struct {
		rule *GpaRepeatRule
	}{
		"Simple GPA repeat": {
			rule: NewGpaRepeatRule(constants.Course{Prefix: "CS", Number: "1337"}),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assertJSONRoundTrip[GpaRepeatRule](t, tc.rule, "gpa_repeat")
		})
	}
}
