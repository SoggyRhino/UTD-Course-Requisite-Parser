package conditions

import (
	"parser/objects/constants"
	"strings"
	"testing"
)

type nilCondition struct{}

func (n nilCondition) Fulfils(constants.UserInfo, bool) *constants.Evaluation {
	return nil
}

func TestAlternativeConditionFulfils(t *testing.T) {
	testCases := map[string]struct {
		condition Condition
		userInfo  constants.UserInfo
		expected  constants.Evaluation
	}{
		"Passes through satisfied condition": {
			condition: NewCourseCondition("BIOL", "2311", ""),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{
					{Prefix: "BIOL", Number: "2311"}: "B",
				},
			},
			expected: constants.Evaluation{
				Name:    "BIOL 2311",
				Status:  constants.StatusPass,
				Summary: "Completed BIOL 2311 with grade B",
			},
		},
		"Unknown when standard path definitely fails": {
			condition: NewCourseCondition("BIOL", "2311", ""),
			userInfo:  constants.UserInfo{},
			expected: constants.Evaluation{
				Name:    "Alternative",
				Status:  constants.StatusUnknown,
				Summary: "Standard path not satisfied — an equivalent may also be accepted (contact adviser)",
				Children: []constants.Evaluation{
					{
						Name:    "BIOL 2311",
						Status:  constants.StatusDefiniteFail,
						Summary: "Have not taken BIOL 2311",
					},
				},
			},
		},
		"Unknown when standard path may fail": {
			condition: NewCourseCondition("BIOL", "2311", ""),
			userInfo: constants.UserInfo{
				CurrentEnrollment: []constants.Course{{Prefix: "BIOL", Number: "2311"}},
			},
			expected: constants.Evaluation{
				Name:    "Alternative",
				Status:  constants.StatusUnknown,
				Summary: "Standard path not satisfied — an equivalent may also be accepted (contact adviser)",
				Children: []constants.Evaluation{
					{
						Name:    "BIOL 2311",
						Status:  constants.StatusPossibleFail,
						Summary: "Currently enrolled in BIOL 2311 — awaiting final grade",
					},
				},
			},
		},
		"Wraps nil inner result as system error child": {
			condition: nilCondition{},
			userInfo:  constants.UserInfo{},
			expected: constants.Evaluation{
				Name:    "Alternative",
				Status:  constants.StatusUnknown,
				Summary: "Standard path not satisfied — an equivalent may also be accepted (contact adviser)",
				Children: []constants.Evaluation{
					{
						Name:    "Alternative",
						Status:  constants.StatusSystemError,
						Summary: "inner condition returned nil",
					},
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := NewAlternativeCondition(tc.condition).Fulfils(tc.userInfo, false)
			assertEval(t, tc.expected, got)
		})
	}
}

func TestAlternativeCondition_JSON(t *testing.T) {
	testCases := map[string]struct {
		condition Condition
	}{
		"Single course": {
			condition: NewCourseCondition("BIOL", "2311", ""),
		},
		"Course alternatives": {
			condition: NewOrCondition(
				NewCourseCondition("PHIL", "1301", ""),
				NewCourseCondition("PHIL", "1305", ""),
				NewCourseCondition("PHIL", "1306", ""),
			),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			original := NewAlternativeCondition(tc.condition)
			assertJSONRoundTrip[AlternativeCondition](t, original, "alternative")
		})
	}
}

func TestAlternativeCondition_UnmarshalJSON_Errors(t *testing.T) {
	testCases := map[string]struct {
		input       string
		expectedErr string
	}{
		"Malformed JSON syntax": {
			input:       `{"condition": `,
			expectedErr: "unexpected end of JSON input",
		},
		"Invalid inner condition type": {
			input:       `{"condition": {"type": "INVALID_UNKNOWN_CONDITION_TYPE"}}`,
			expectedErr: "unknown condition type",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			var decoded AlternativeCondition
			err := decoded.UnmarshalJSON([]byte(tc.input))

			if tc.expectedErr == "" {
				if err != nil {
					t.Fatalf("Expected no error, got: %v", err)
				}
				return
			}
			if err == nil {
				t.Fatalf("Expected error containing %q, got nil", tc.expectedErr)
			}
			if !strings.Contains(err.Error(), tc.expectedErr) {
				t.Errorf("Expected error containing %q, got: %q", tc.expectedErr, err.Error())
			}
		})
	}
}
