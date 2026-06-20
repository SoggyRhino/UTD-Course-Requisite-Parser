package rules

import (
	"parser/objects/constants"
	"testing"
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
