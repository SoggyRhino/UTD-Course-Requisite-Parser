package rules

import (
	"parser/objects/constants"
	"testing"
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
