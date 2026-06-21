package rules

import (
	"parser/objects/constants"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreditForRule_JSON(t *testing.T) {
	testCases := map[string]struct {
		rule *CreditForRule
	}{
		"Simple collection": {
			rule: NewCreditForRule(NewSimpleCourseCollection([]constants.Course{{Prefix: "CS", Number: "1337"}})),
		},
		"And collection": {
			rule: NewCreditForRule(NewAndCourseCollection(
				NewSimpleCourseCollection([]constants.Course{{Prefix: "CS", Number: "1337"}}),
				NewSimpleCourseCollection([]constants.Course{{Prefix: "CS", Number: "2337"}}),
			)),
		},
		"Or collection": {
			rule: NewCreditForRule(NewOrCourseCollection(
				NewSimpleCourseCollection([]constants.Course{{Prefix: "CS", Number: "1337"}}),
				NewSimpleCourseCollection([]constants.Course{{Prefix: "CS", Number: "2337"}}),
			)),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assertJSONRoundTrip[CreditForRule](t, tc.rule, "credit_for")
		})
	}
}

func TestAndCourseCollection_JSON(t *testing.T) {
	col := NewAndCourseCollection(
		NewSimpleCourseCollection([]constants.Course{{Prefix: "CS", Number: "1337"}}),
		NewSimpleCourseCollection([]constants.Course{{Prefix: "CS", Number: "2337"}}),
	)
	assertJSONRoundTrip[AndCourseCollection](t, col, "and")
}

func TestOrCourseCollection_JSON(t *testing.T) {
	col := NewOrCourseCollection(
		NewSimpleCourseCollection([]constants.Course{{Prefix: "CS", Number: "1337"}}),
		NewSimpleCourseCollection([]constants.Course{{Prefix: "CS", Number: "2337"}}),
	)
	assertJSONRoundTrip[OrCourseCollection](t, col, "or")
}

func TestSimpleCourseCollection_JSON(t *testing.T) {
	col := NewSimpleCourseCollection([]constants.Course{{Prefix: "CS", Number: "1337"}})
	assertJSONRoundTrip[SimpleCourseCollection](t, col, "simple")
}

func TestCreditForRule_Fulfils(t *testing.T) {
	testCases := map[string]struct {
		rule     *CreditForRule
		userInfo constants.UserInfo
		expected *constants.Evaluation
	}{
		"Passes when no courses taken": {
			rule: NewCreditForRule(NewSimpleCourseCollection([]constants.Course{{Prefix: "CS", Number: "1337"}})),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{},
			},
			expected: &constants.Evaluation{
				Name:    "Credit For Rule",
				Status:  constants.StatusPass,
				Summary: "Student satisfies credit for rule",
			},
		},
		"Fails when simple course taken": {
			rule: NewCreditForRule(NewSimpleCourseCollection([]constants.Course{{Prefix: "CS", Number: "1337"}})),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{
					{Prefix: "CS", Number: "1337"}: "A",
				},
			},
			expected: &constants.Evaluation{
				Name:    "Credit For Rule",
				Status:  constants.StatusDefiniteFail,
				Summary: "Student violates credit for rule: CS 1337",
			},
		},
		"Fails when OR course taken": {
			rule: NewCreditForRule(NewOrCourseCollection(
				NewSimpleCourseCollection([]constants.Course{{Prefix: "CS", Number: "1337"}}),
				NewSimpleCourseCollection([]constants.Course{{Prefix: "SE", Number: "1337"}}),
			)),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{
					{Prefix: "SE", Number: "1337"}: "A",
				},
			},
			expected: &constants.Evaluation{
				Name:    "Credit For Rule",
				Status:  constants.StatusDefiniteFail,
				Summary: "Student violates credit for rule: (CS 1337 OR SE 1337)",
			},
		},
		"Passes when only one of AND taken": {
			rule: NewCreditForRule(NewAndCourseCollection(
				NewSimpleCourseCollection([]constants.Course{{Prefix: "CS", Number: "1337"}}),
				NewSimpleCourseCollection([]constants.Course{{Prefix: "SE", Number: "1337"}}),
			)),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{
					{Prefix: "CS", Number: "1337"}: "A",
				},
			},
			expected: &constants.Evaluation{
				Name:    "Credit For Rule",
				Status:  constants.StatusPass,
				Summary: "Student satisfies credit for rule",
			},
		},
		"Fails when both of AND taken": {
			rule: NewCreditForRule(NewAndCourseCollection(
				NewSimpleCourseCollection([]constants.Course{{Prefix: "CS", Number: "1337"}}),
				NewSimpleCourseCollection([]constants.Course{{Prefix: "SE", Number: "1337"}}),
			)),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{
					{Prefix: "CS", Number: "1337"}: "A",
					{Prefix: "SE", Number: "1337"}: "B",
				},
			},
			expected: &constants.Evaluation{
				Name:    "Credit For Rule",
				Status:  constants.StatusDefiniteFail,
				Summary: "Student violates credit for rule: (CS 1337 AND SE 1337)",
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
