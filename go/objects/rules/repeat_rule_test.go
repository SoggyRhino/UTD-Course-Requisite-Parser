package rules

import (
	"parser/objects/constants"
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestRepeatRule_Fulfils(t *testing.T) {
	testCases := map[string]struct {
		rule     *RepeatRule
		userInfo constants.UserInfo
		expected *constants.Evaluation
	}{
		"Taken fewer than count limit": {
			rule: NewCourseRepeatRule([]constants.Course{{Prefix: "CS", Number: "1337"}}),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{},
			},
			expected: &constants.Evaluation{
				Name:    "Repeat Rule",
				Status:  constants.StatusPass,
				Summary: "Course repeat limits have not been exceeded",
			},
		},
		"Taken exactly count limit": {
			rule: NewCourseRepeatRule([]constants.Course{{Prefix: "CS", Number: "1337"}}),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{
					{Prefix: "CS", Number: "1337"}: "A",
				},
			},
			expected: &constants.Evaluation{
				Name:    "Repeat Rule",
				Status:  constants.StatusDefiniteFail,
				Summary: "Course has been repeated 1 times, which meets or exceeds the limit of 1",
			},
		},
		"Taken hours under limit": {
			rule: NewRepeatRule(0, 6, []constants.Course{{Prefix: "CS", Number: "1337"}}, ""),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{
					{Prefix: "CS", Number: "1337"}: "A",
				},
			},
			expected: &constants.Evaluation{
				Name:    "Repeat Rule",
				Status:  constants.StatusPass,
				Summary: "Course repeat limits have not been exceeded",
			},
		},
		"Taken hours over limit": {
			rule: NewRepeatRule(0, 3, []constants.Course{{Prefix: "CS", Number: "1337"}}, ""),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{
					{Prefix: "CS", Number: "1337"}: "A",
				},
			},
			expected: &constants.Evaluation{
				Name:    "Repeat Rule",
				Status:  constants.StatusDefiniteFail,
				Summary: "Course has been repeated for 3 hours, which meets or exceeds the limit of 3 hours",
			},
		},
		"Enrolled in course exceeding limit": {
			rule: NewCourseRepeatRule([]constants.Course{{Prefix: "CS", Number: "1337"}, {Prefix: "CS", Number: "1336"}}),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{
					{Prefix: "CS", Number: "1337"}: "A",
				},
				CurrentEnrollment: []constants.Course{{Prefix: "CS", Number: "1336"}},
			},
			expected: &constants.Evaluation{
				Name:    "Repeat Rule",
				Status:  constants.StatusDefiniteFail,
				Summary: "Course has been repeated 2 times, which meets or exceeds the limit of 1",
			},
		},
		"Internship case": {
			rule:     NewInternshipRepeatRule("CS"),
			userInfo: constants.UserInfo{},
			expected: &constants.Evaluation{
				Name:    "Repeat Rule (Major Internship)",
				Status:  constants.StatusPossibleFail,
				Summary: "Major internship repeat rule not implemented",
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := tc.rule.Fulfils(tc.userInfo)
			if diff := cmp.Diff(tc.expected, got); diff != "" {
				t.Errorf("Unexpected result (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGpaRepeatRule_Fulfils(t *testing.T) {
	testCases := map[string]struct {
		rule     *GpaRepeatRule
		userInfo constants.UserInfo
		expected *constants.Evaluation
	}{
		"Matches plan and taken": {
			rule: &GpaRepeatRule{
				Course:       constants.Course{Prefix: "CS", Number: "1337"},
				AcademicPlan: "CS",
			},
			userInfo: constants.UserInfo{
				AcademicPlan: "CS",
				Taken: map[constants.Course]constants.Grade{
					{Prefix: "CS", Number: "1337"}: "B",
				},
			},
			expected: &constants.Evaluation{
				Name:    "GPA Repeat Rule",
				Status:  constants.StatusDefiniteFail,
				Summary: "Cannot repeat course CS 1337 to improve GPA",
			},
		},
		"Different plan but taken": {
			rule: &GpaRepeatRule{
				Course:       constants.Course{Prefix: "CS", Number: "1337"},
				AcademicPlan: "CS",
			},
			userInfo: constants.UserInfo{
				AcademicPlan: "SE",
				Taken: map[constants.Course]constants.Grade{
					{Prefix: "CS", Number: "1337"}: "B",
				},
			},
			expected: &constants.Evaluation{
				Name:    "GPA Repeat Rule",
				Status:  constants.StatusPass,
				Summary: "Course can be repeated to improve GPA",
			},
		},
		"Matches plan but not taken": {
			rule: &GpaRepeatRule{
				Course:       constants.Course{Prefix: "CS", Number: "1337"},
				AcademicPlan: "CS",
			},
			userInfo: constants.UserInfo{
				AcademicPlan: "CS",
			},
			expected: &constants.Evaluation{
				Name:    "GPA Repeat Rule",
				Status:  constants.StatusPass,
				Summary: "Course can be repeated to improve GPA",
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := tc.rule.Fulfils(tc.userInfo)
			if diff := cmp.Diff(tc.expected, got); diff != "" {
				t.Errorf("Unexpected result (-want +got):\n%s", diff)
			}
		})
	}
}
