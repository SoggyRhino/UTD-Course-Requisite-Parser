package rules

import (
	"parser/objects/constants"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDegreeSatisfactionRule_JSON(t *testing.T) {
	testCases := map[string]struct {
		rule *DegreeSatisfactionRule
	}{
		"From prefix": {
			rule: NewDegreeSatisfactionRuleFromPrefix([]string{"CS"}, constants.Undergraduate),
		},
		"From degree": {
			rule: NewDegreeSatisfactionRuleFromDegree([]string{"BS"}, constants.Undergraduate),
		},
		"From school": {
			rule: NewDegreeSatisfactionRuleFromSchool([]string{"ECS"}, constants.Undergraduate),
		},
		"Math rule": {
			rule: NewMathDegreeSatisfactionRule(),
		},
		"Elective rule": {
			rule: NewDegreeSatisfactionRuleFromElectives(NewDegreeSatisfactionRuleFromPrefix([]string{"CS"}, constants.Undergraduate)),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assertJSONRoundTrip[DegreeSatisfactionRule](t, tc.rule, "degree_satisfaction")
		})
	}
}

func TestDegreeSatisfactionRule_Fulfils(t *testing.T) {
	testCases := map[string]struct {
		rule     *DegreeSatisfactionRule
		userInfo constants.UserInfo
		expected *constants.Evaluation
	}{
		"Matches prefix and level": {
			rule: NewDegreeSatisfactionRuleFromPrefix([]string{"CS"}, constants.Undergraduate),
			userInfo: constants.UserInfo{
				Major:       "CS",
				DegreeLevel: constants.Undergraduate,
			},
			expected: &constants.Evaluation{
				Name:    "Degree Satisfaction Rule",
				Status:  constants.StatusDefiniteFail,
				Summary: "Course cannot be used to satisfy degree requirements for CS majors",
			},
		},
		"Matches degree but different level": {
			rule: NewDegreeSatisfactionRuleFromDegree([]string{"BS"}, constants.Graduate),
			userInfo: constants.UserInfo{
				Major:       "BS",
				DegreeLevel: constants.Undergraduate,
			},
			expected: &constants.Evaluation{
				Name:    "Degree Satisfaction Rule",
				Status:  constants.StatusPass,
				Summary: "Student satisfies degree requirements rule",
			},
		},
		"Matches school with Math true": {
			rule: NewMathDegreeSatisfactionRule(),
			userInfo: constants.UserInfo{
				School: "Mathematics",
			},
			expected: &constants.Evaluation{
				Name:    "Degree Satisfaction Rule",
				Status:  constants.StatusDefiniteFail,
				Summary: "Course cannot be used to satisfy mathematics requirements by students in Mathematics",
			},
		},
		"Matches school with Math false": {
			rule: NewDegreeSatisfactionRuleFromSchool([]string{"ECS"}, ""),
			userInfo: constants.UserInfo{
				School: "ECS",
			},
			expected: &constants.Evaluation{
				Name:    "Degree Satisfaction Rule",
				Status:  constants.StatusDefiniteFail,
				Summary: "Course cannot be used to satisfy degree requirements for the ECS school",
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
