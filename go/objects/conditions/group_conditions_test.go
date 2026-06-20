package conditions

import (
	"parser/objects/constants"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type staticCondition struct {
	evaluation *constants.Evaluation
}

func (s staticCondition) Fulfils(constants.UserInfo, bool) *constants.Evaluation {
	return s.evaluation
}

func TestOrConditionFulfils(t *testing.T) {
	testCases := map[string]struct {
		condition OrCondition
		expected  constants.Evaluation
	}{
		"Returns pass when first condition passes": {
			condition: OrCondition{Conditions: []Condition{
				staticCondition{evaluation: &constants.Evaluation{Status: constants.StatusPass, Summary: "first passes"}},
				staticCondition{evaluation: &constants.Evaluation{Status: constants.StatusDefiniteFail, Summary: "second fails"}},
			}},
			expected: constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: "At least one of 2 conditions satisfied",
				Children: []constants.Evaluation{
					{Status: constants.StatusPass, Summary: "first passes"},
				},
			},
		},
		"Returns pass when later condition passes": {
			condition: OrCondition{Conditions: []Condition{
				staticCondition{evaluation: &constants.Evaluation{Status: constants.StatusDefiniteFail, Summary: "first fails"}},
				staticCondition{evaluation: &constants.Evaluation{Status: constants.StatusPass, Summary: "second passes"}},
				staticCondition{evaluation: &constants.Evaluation{Status: constants.StatusDefiniteFail, Summary: "third not evaluated"}},
			}},
			expected: constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: "At least one of 3 conditions satisfied",
				Children: []constants.Evaluation{
					{Status: constants.StatusDefiniteFail, Summary: "first fails"},
					{Status: constants.StatusPass, Summary: "second passes"},
				},
			},
		},
		"Returns definite fail when no condition passes": {
			condition: OrCondition{Conditions: []Condition{
				staticCondition{evaluation: &constants.Evaluation{Status: constants.StatusDefiniteFail, Summary: "first fails"}},
				staticCondition{evaluation: &constants.Evaluation{Status: constants.StatusPossibleFail, Summary: "second may fail"}},
			}},
			expected: constants.Evaluation{
				Status:  constants.StatusDefiniteFail,
				Summary: "None of 2 conditions satisfied",
				Children: []constants.Evaluation{
					{Status: constants.StatusDefiniteFail, Summary: "first fails"},
					{Status: constants.StatusPossibleFail, Summary: "second may fail"},
				},
			},
		},
		"Wraps nil child evaluation": {
			condition: OrCondition{Conditions: []Condition{
				staticCondition{},
			}},
			expected: constants.Evaluation{
				Status:  constants.StatusSystemError,
				Summary: "condition returned nil",
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := tc.condition.Fulfils(constants.UserInfo{}, false)
			assertEval(t, tc.expected, got)
		})
	}
}

func TestOrCondition_Constructors(t *testing.T) {
	course := NewCourseCondition("BIOL", "2311", "")

	t.Run("Single condition returns condition", func(t *testing.T) {
		got := NewOrCondition(course)
		if diff := cmp.Diff(course, got); diff != "" {
			t.Errorf("Unexpected condition (-want +got):\n%s", diff)
		}
	})

	t.Run("Multiple conditions creates OrCondition", func(t *testing.T) {
		first := NewCourseCondition("BIOL", "2311", "")
		second := NewCourseCondition("CHEM", "1311", "")
		expected := &OrCondition{Conditions: []Condition{first, second}}

		got := NewOrCondition(first, second)
		if diff := cmp.Diff(expected, got); diff != "" {
			t.Errorf("Unexpected condition (-want +got):\n%s", diff)
		}
	})

	t.Run("Flattens expression operands", func(t *testing.T) {
		first := NewCourseCondition("BIOL", "2311", "")
		second := NewCourseCondition("CHEM", "1311", "")
		third := NewCourseCondition("PHYS", "2325", "")
		fourth := NewCourseCondition("MATH", "2413", "")
		expected := &OrCondition{Conditions: []Condition{first, second, third, fourth}}

		got := NewOrConditionFromExpr(
			&OrCondition{Conditions: []Condition{first, second}},
			&OrCondition{Conditions: []Condition{third, fourth}},
		)
		if diff := cmp.Diff(expected, got); diff != "" {
			t.Errorf("Unexpected condition (-want +got):\n%s", diff)
		}
	})

	t.Run("Keeps non-Or expression operands", func(t *testing.T) {
		first := NewCourseCondition("BIOL", "2311", "")
		second := NewCourseCondition("CHEM", "1311", "")
		expected := &OrCondition{Conditions: []Condition{first, second}}

		got := NewOrConditionFromExpr(first, second)
		if diff := cmp.Diff(expected, got); diff != "" {
			t.Errorf("Unexpected condition (-want +got):\n%s", diff)
		}
	})
}

func TestOrConditionAppendGrade(t *testing.T) {
	first := NewCourseCondition("BIOL", "2311", "")
	second := NewCreditHoursCondition(30)
	condition := OrCondition{Conditions: []Condition{first, second}}

	condition.AppendGrade("C")

	if first.MinGrade != "C" {
		t.Errorf("Expected first minimum grade %q, got %q", constants.Grade("C"), first.MinGrade)
	}
}

func TestOrCondition_JSON(t *testing.T) {
	assertJSONRoundTrip[OrCondition](t, &OrCondition{Conditions: []Condition{
		NewCourseCondition("BIOL", "2311", ""),
		NewCourseCondition("CHEM", "1311", ""),
	}}, "or")
}

func TestOrCondition_UnmarshalJSON_Errors(t *testing.T) {
	testCases := map[string]struct {
		input       string
		expectedErr string
	}{
		"Malformed JSON syntax": {
			input:       `{"conditions": [`,
			expectedErr: "unexpected end of JSON input",
		},
		"Invalid inner condition type": {
			input:       `{"conditions": [{"type": "INVALID_UNKNOWN_CONDITION_TYPE"}]}`,
			expectedErr: "unknown condition type",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			var decoded OrCondition
			err := decoded.UnmarshalJSON([]byte(tc.input))

			if err == nil {
				t.Fatalf("Expected error containing %q, got nil", tc.expectedErr)
			}
			if !strings.Contains(err.Error(), tc.expectedErr) {
				t.Errorf("Expected error containing %q, got: %q", tc.expectedErr, err.Error())
			}
		})
	}
}

func TestAndConditionFulfils(t *testing.T) {
	testCases := map[string]struct {
		condition AndCondition
		expected  constants.Evaluation
	}{
		"Returns pass when all conditions pass": {
			condition: AndCondition{Conditions: []Condition{
				staticCondition{evaluation: &constants.Evaluation{Status: constants.StatusPass, Summary: "first passes"}},
				staticCondition{evaluation: &constants.Evaluation{Status: constants.StatusPass, Summary: "second passes"}},
			}},
			expected: constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: "All 2 conditions satisfied",
				Children: []constants.Evaluation{
					{Status: constants.StatusPass, Summary: "first passes"},
					{Status: constants.StatusPass, Summary: "second passes"},
				},
			},
		},
		"Returns lowest priority failing status": {
			condition: AndCondition{Conditions: []Condition{
				staticCondition{evaluation: &constants.Evaluation{Status: constants.StatusDefiniteFail, Summary: "first fails"}},
				staticCondition{evaluation: &constants.Evaluation{Status: constants.StatusUnknown, Summary: "second unknown"}},
			}},
			expected: constants.Evaluation{
				Status:  constants.StatusUnknown,
				Summary: "Not all conditions satisfied (2 total)",
				Children: []constants.Evaluation{
					{Status: constants.StatusDefiniteFail, Summary: "first fails"},
					{Status: constants.StatusUnknown, Summary: "second unknown"},
				},
			},
		},
		"Wraps nil child evaluation": {
			condition: AndCondition{Conditions: []Condition{
				staticCondition{},
			}},
			expected: constants.Evaluation{
				Status:  constants.StatusSystemError,
				Summary: "Not all conditions satisfied (1 total)",
				Children: []constants.Evaluation{
					{Status: constants.StatusSystemError, Summary: "condition returned nil"},
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := tc.condition.Fulfils(constants.UserInfo{}, false)
			assertEval(t, tc.expected, got)
		})
	}
}

func TestAndCondition_Constructors(t *testing.T) {
	course := NewCourseCondition("BIOL", "2311", "")

	t.Run("Single condition returns condition", func(t *testing.T) {
		got := NewAndCondition(course)
		if diff := cmp.Diff(course, got); diff != "" {
			t.Errorf("Unexpected condition (-want +got):\n%s", diff)
		}
	})

	t.Run("Multiple conditions creates AndCondition", func(t *testing.T) {
		first := NewCourseCondition("BIOL", "2311", "")
		second := NewCourseCondition("CHEM", "1311", "")
		expected := &AndCondition{Conditions: []Condition{first, second}}

		got := NewAndCondition(first, second)
		if diff := cmp.Diff(expected, got); diff != "" {
			t.Errorf("Unexpected condition (-want +got):\n%s", diff)
		}
	})

	t.Run("Flattens expression operands", func(t *testing.T) {
		first := NewCourseCondition("BIOL", "2311", "")
		second := NewCourseCondition("CHEM", "1311", "")
		third := NewCourseCondition("PHYS", "2325", "")
		fourth := NewCourseCondition("MATH", "2413", "")
		expected := &AndCondition{Conditions: []Condition{first, second, third, fourth}}

		got := NewAndConditionFromExpr(
			&AndCondition{Conditions: []Condition{first, second}},
			&AndCondition{Conditions: []Condition{third, fourth}},
		)
		if diff := cmp.Diff(expected, got); diff != "" {
			t.Errorf("Unexpected condition (-want +got):\n%s", diff)
		}
	})

	t.Run("Keeps non-And expression operands", func(t *testing.T) {
		first := NewCourseCondition("BIOL", "2311", "")
		second := NewCourseCondition("CHEM", "1311", "")
		expected := &AndCondition{Conditions: []Condition{first, second}}

		got := NewAndConditionFromExpr(first, second)
		if diff := cmp.Diff(expected, got); diff != "" {
			t.Errorf("Unexpected condition (-want +got):\n%s", diff)
		}
	})
}

func TestAndConditionAppendGrade(t *testing.T) {
	first := NewCourseCondition("BIOL", "2311", "")
	second := NewCreditHoursCondition(30)
	condition := AndCondition{Conditions: []Condition{first, second}}

	condition.AppendGrade("C")

	if first.MinGrade != "C" {
		t.Errorf("Expected first minimum grade %q, got %q", constants.Grade("C"), first.MinGrade)
	}
}

func TestAndCondition_JSON(t *testing.T) {
	assertJSONRoundTrip[AndCondition](t, &AndCondition{Conditions: []Condition{
		NewCourseCondition("BIOL", "2311", ""),
		NewCourseCondition("CHEM", "1311", ""),
	}}, "and")
}

func TestOrConditionConstructors(t *testing.T) {
	c1 := NewCourseCondition("CS", "1", "")
	c2 := NewCourseCondition("CS", "2", "")

	t.Run("Single condition", func(t *testing.T) {
		if NewOrCondition(c1) != c1 {
			t.Error("NewOrCondition with 1 condition should return the condition itself")
		}
	})

	t.Run("Flattening", func(t *testing.T) {
		or1 := &OrCondition{Conditions: []Condition{c1}}
		or2 := &OrCondition{Conditions: []Condition{c2}}
		flat := NewOrConditionFromExpr(or1, or2)
		if len(flat.Conditions) != 2 {
			t.Errorf("Expected 2 conditions after flattening, got %d", len(flat.Conditions))
		}
	})

	t.Run("Non-flattening", func(t *testing.T) {
		flat := NewOrConditionFromExpr(c1, c2)
		if len(flat.Conditions) != 2 {
			t.Errorf("Expected 2 conditions, got %d", len(flat.Conditions))
		}
	})
}

func TestAndConditionConstructors(t *testing.T) {
	c1 := NewCourseCondition("CS", "1", "")
	c2 := NewCourseCondition("CS", "2", "")

	t.Run("Single condition", func(t *testing.T) {
		if NewAndCondition(c1) != c1 {
			t.Error("NewAndCondition with 1 condition should return the condition itself")
		}
	})

	t.Run("Flattening", func(t *testing.T) {
		and1 := &AndCondition{Conditions: []Condition{c1}}
		and2 := &AndCondition{Conditions: []Condition{c2}}
		flat := NewAndConditionFromExpr(and1, and2)
		if len(flat.Conditions) != 2 {
			t.Errorf("Expected 2 conditions after flattening, got %d", len(flat.Conditions))
		}
	})

	t.Run("Non-flattening", func(t *testing.T) {
		flat := NewAndConditionFromExpr(c1, c2)
		if len(flat.Conditions) != 2 {
			t.Errorf("Expected 2 conditions, got %d", len(flat.Conditions))
		}
	})
}

func TestGroupCondition_AppendGrade(t *testing.T) {
	cc := NewCourseCondition("CS", "1337", "C")
	or := NewOrCondition(cc, NewCourseCondition("CS", "2337", "")).(*OrCondition)
	or.AppendGrade("B")
	if cc.MinGrade != "B" {
		t.Errorf("Expected grade B, got %s", cc.MinGrade)
	}

	cc2 := NewCourseCondition("CS", "1337", "C")
	and := NewAndCondition(cc2, NewCourseCondition("CS", "2337", "")).(*AndCondition)
	and.AppendGrade("B")
	if cc2.MinGrade != "B" {
		t.Errorf("Expected grade B, got %s", cc2.MinGrade)
	}
}

func TestAndCondition_UnmarshalJSON_Errors(t *testing.T) {
	testCases := map[string]struct {
		input       string
		expectedErr string
	}{
		"Malformed JSON syntax": {
			input:       `{"conditions": [`,
			expectedErr: "unexpected end of JSON input",
		},
		"Invalid inner condition type": {
			input:       `{"conditions": [{"type": "INVALID_UNKNOWN_CONDITION_TYPE"}]}`,
			expectedErr: "unknown condition type",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			var decoded AndCondition
			err := decoded.UnmarshalJSON([]byte(tc.input))

			if err == nil {
				t.Fatalf("Expected error containing %q, got nil", tc.expectedErr)
			}
			if !strings.Contains(err.Error(), tc.expectedErr) {
				t.Errorf("Expected error containing %q, got: %q", tc.expectedErr, err.Error())
			}
		})
	}
}
