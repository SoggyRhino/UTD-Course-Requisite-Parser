package conditions

import (
	"parser/objects/constants"
	"testing"
)

func TestGPAConditionFulfils(t *testing.T) {
	testCases := map[string]struct {
		condition GPACondition
		userInfo  constants.UserInfo
		expected  constants.Evaluation
	}{
		"GPA match": {
			condition: *NewGpaCondition(3.0),
			userInfo:  constants.UserInfo{GPA: 3.5},
			expected: constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: "GPA is 3.50 (requires 3.00)",
			},
		},
		"GPA exact match": {
			condition: *NewGpaCondition(3.0),
			userInfo:  constants.UserInfo{GPA: 3.0},
			expected: constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: "GPA is 3.00 (requires 3.00)",
			},
		},
		"GPA mismatch": {
			condition: *NewGpaCondition(3.0),
			userInfo:  constants.UserInfo{GPA: 2.5},
			expected: constants.Evaluation{
				Status:  constants.StatusDefiniteFail,
				Summary: "GPA is 2.50 but requires 3.00",
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := tc.condition.Fulfils(tc.userInfo, false)
			assertEval(t, tc.expected, got)
		})
	}
}

func TestGPACondition_JSON(t *testing.T) {
	cond := NewGpaCondition(3.0)
	assertJSONRoundTrip[GPACondition](t, cond, "gpa")
}

func TestNewGpaConditionWithDegree(t *testing.T) {
	cond := NewGpaConditionWithDegree(3.0, "CS")
	if cond.Degree != "CS" {
		t.Errorf("Expected degree CS, got %s", cond.Degree)
	}
}
