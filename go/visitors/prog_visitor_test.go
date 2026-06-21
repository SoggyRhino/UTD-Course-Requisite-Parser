package visitors

import (
	"parser/objects"
	"parser/objects/conditions"
	"parser/objects/constants"
	"parser/objects/rules"
	"parser/parser"
	"testing"
)

func TestVisitProg(t *testing.T) {

	testCases := map[string]struct {
		Input  string
		Result objects.Requirements
	}{
		"Simple": {
			Input: "Prerequisite: ACCT 2301. Repeat Restriction.",
			Result: objects.Requirements{
				PreReqs: conditions.NewCourseCondition("ACCT", "2301", ""),
				Rules: []rules.Rule{
					rules.NewRepeatRule(1, 0, []constants.Course{}, ""),
				},
			},
		},

		"Longer": {
			Input: "Prerequisites: (ACCT 2301 with a minimum grade of C) and (ACCT 2302 with a minimum grade of C) and ITSS 3300 and (MATH 1325 or MATH 2413 or MATH 2417). May not be used to satisfy BS INTS degree requirements.",
			Result: objects.Requirements{
				PreReqs: conditions.NewAndCondition(
					conditions.NewCourseCondition("ACCT", "2301", "C"),
					conditions.NewCourseCondition("ACCT", "2302", "C"),
					conditions.NewCourseCondition("ITSS", "3300", ""),
					conditions.NewOrCondition(
						conditions.NewCourseCondition("MATH", "1325", ""),
						conditions.NewCourseCondition("MATH", "2413", ""),
						conditions.NewCourseCondition("MATH", "2417", ""),
					),
				),
				Rules: []rules.Rule{
					rules.NewDegreeSatisfactionRuleFromPrefix([]string{"INTS"}, constants.Undergraduate),
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[objects.Requirements](t, tc.Input, rule((*parser.RequirementsParser).Prog), tc.Result)
		})
	}
}
