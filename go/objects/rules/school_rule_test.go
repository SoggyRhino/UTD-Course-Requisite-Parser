package rules

import (
	"parser/objects/constants"
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestSchoolRule_Fulfils(t *testing.T) {
	testCases := map[string]struct {
		rule     *SchoolRule
		userInfo constants.UserInfo
		expected *constants.Evaluation
	}{
		"Matches school": {
			rule: NewSchoolRule([]string{"ECS", "NSM"}),
			userInfo: constants.UserInfo{
				School: "ECS",
			},
			expected: &constants.Evaluation{
				Name:    "School Rule",
				Status:  constants.StatusPass,
				Summary: "Student is in required school: ECS",
			},
		},
		"Does not match school": {
			rule: NewSchoolRule([]string{"ECS", "NSM"}),
			userInfo: constants.UserInfo{
				School: "JSOM",
			},
			expected: &constants.Evaluation{
				Name:    "School Rule",
				Status:  constants.StatusDefiniteFail,
				Summary: "Student is not in any of the required schools: ECS, NSM",
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
