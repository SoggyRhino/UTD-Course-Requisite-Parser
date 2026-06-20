package rules

import (
	"parser/objects/constants"
	"testing"
)

func TestSameAsRule_JSON(t *testing.T) {
	testCases := map[string]struct {
		rule *SameAsRule
	}{
		"Simple same as": {
			rule: NewSameAsRule([]constants.Course{{Prefix: "CS", Number: "1337"}}),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assertJSONRoundTrip[SameAsRule](t, tc.rule, "same_as")
		})
	}
}
