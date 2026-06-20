package rules

import (
	"testing"
)

func TestLivingLearningRule_JSON(t *testing.T) {
	testCases := map[string]struct {
		rule *LivingLearningRule
	}{
		"From prefixes": {
			rule: NewLivingLearningRuleFromPrefixes([]string{"CS"}),
		},
		"From degrees": {
			rule: NewLivingLearningRuleFromDegrees([]string{"BS"}),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assertJSONRoundTrip[LivingLearningRule](t, tc.rule, "living_learning")
		})
	}
}
