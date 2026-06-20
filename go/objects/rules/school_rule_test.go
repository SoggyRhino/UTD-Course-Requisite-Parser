package rules

import (
	"testing"
)

func TestSchoolRule_JSON(t *testing.T) {
	testCases := map[string]struct {
		rule *SchoolRule
	}{
		"Simple school rule": {
			rule: NewSchoolRule([]string{"ECS"}),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assertJSONRoundTrip[SchoolRule](t, tc.rule, "school")
		})
	}
}
