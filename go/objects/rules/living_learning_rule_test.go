package rules

import (
	"parser/objects/constants"
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestLivingLearningRule_Fulfils(t *testing.T) {
	testCases := map[string]struct {
		rule     *LivingLearningRule
		userInfo constants.UserInfo
		expected *constants.Evaluation
	}{
		"Matches degree": {
			rule: NewLivingLearningRuleFromDegrees([]string{"CS", "SE"}),
			userInfo: constants.UserInfo{
				Major: "CS",
			},
			expected: &constants.Evaluation{
				Name:    "Living Learning Rule",
				Status:  constants.StatusPass,
				Summary: "Student's major matches living learning degree: CS",
			},
		},
		"Does not match degree": {
			rule: NewLivingLearningRuleFromDegrees([]string{"CS", "SE"}),
			userInfo: constants.UserInfo{
				Major: "MATH",
			},
			expected: &constants.Evaluation{
				Name:    "Living Learning Rule",
				Status:  constants.StatusDefiniteFail,
				Summary: "Student's major \"MATH\" is not in the required living learning degrees: CS, SE",
			},
		},
		"Prefixes not implemented": {
			rule: NewLivingLearningRuleFromPrefixes([]string{"CS"}),
			userInfo: constants.UserInfo{
				Major: "CS",
			},
			expected: &constants.Evaluation{
				Name:    "Living Learning Rule",
				Status:  constants.StatusPossibleFail,
				Summary: "Living learning prefixes mapping not implemented yet",
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
