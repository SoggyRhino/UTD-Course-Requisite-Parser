package visitors

import (
	"parser/parser"
	"parser/rules"
	"parser/utils"
	"testing"
)

func TestVisitRepeatRule(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result any
	}{
		"Single course repeat restriction": {
			Input: "ACCT 2301 Repeat Restriction",
			Result: rules.NewCourseRepeatRule([]utils.Course{
				{Prefix: "ACCT", Number: "2301"},
			}),
		},
		"Dual prefix course repeat restriction": {
			Input: "CE/CS 2305 Repeat Restriction",
			Result: rules.NewCourseRepeatRule(
				[]utils.Course{{Prefix: "CE", Number: "2305"}, {Prefix: "CS", Number: "2305"}},
			),
		},
		"Internship repeat restriction": {
			Input:  "BBSC Internship Repeat Restriction",
			Result: rules.NewInternshipRepeatRule("BBSC"),
		},
		"Bare repeat restriction": {
			Input:  "Repeat Restriction",
			Result: rules.NewRepeatRule(0, 0, []utils.Course{}, ""),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[any](t, tc.Input, rule((*parser.RequirementsParser).Repeat_rule), tc.Result)
		})
	}
}

func TestVisitRepeatLimitHoursRule(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result *rules.RepeatRule
	}{
		"Repeat max hours": {
			Input:  "Repeat Limit - ACCT 7323 may only be repeated for a maximum of 9 semester credit hours",
			Result: rules.NewRepeatRule(0, 9, []utils.Course{{Prefix: "ACCT", Number: "7323"}}, ""),
		},
		"Repeat hours max suffix": {
			Input:  "Repeat Limit - EEGR 6V88 may only be repeated for 9 semester credit hours maximum",
			Result: rules.NewRepeatRule(0, 9, []utils.Course{{Prefix: "EEGR", Number: "6V88"}}, ""),
		},
		"Course repeat max hours may": {
			Input:  "AHST 6336 Repeat Limit - This course may only be repeated for a maximum of 6 semester credit hours",
			Result: rules.NewRepeatRule(0, 6, []utils.Course{{Prefix: "AHST", Number: "6336"}}, ""),
		},
		"Course repeat max hours can": {
			Input:  "RHET 1302 Repeat Limit - This course can only be repeated for a maximum of 3 semester credit hours",
			Result: rules.NewRepeatRule(0, 3, []utils.Course{{Prefix: "RHET", Number: "1302"}}, ""),
		},
		"Combined repeat max hours": {
			Input:  "Repeat Limit - CS 4V96 and CS 4V98 combined may only be repeated for a maximum of 6 semester credit hours",
			Result: rules.NewRepeatRule(0, 6, []utils.Course{{Prefix: "CS", Number: "4V96"}, {Prefix: "CS", Number: "4V98"}}, ""),
		},
		"Course repeat limit": {
			Input:  "ANGM 2309 Repeat Limit",
			Result: rules.NewCourseRepeatRule([]utils.Course{{Prefix: "ANGM", Number: "2309"}}),
		},
		"Course repeat limit variable number": {
			Input:  "HLTH 4V01 Repeat Limit",
			Result: rules.NewCourseRepeatRule([]utils.Course{{Prefix: "HLTH", Number: "4V01"}}),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[*rules.RepeatRule](t, tc.Input, rule((*parser.RequirementsParser).Repeat_limit_hours_rule), tc.Result)
		})
	}
}

func TestVisitRepeatLimitTimesRule(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result *rules.RepeatRule
	}{
		"Repeat up to N times": {
			Input:  "Repeat Limit - UNIV 4074 may be repeated up to 3 times",
			Result: rules.NewRepeatRule(3, 0, []utils.Course{{Prefix: "UNIV", Number: "4074"}}, ""),
		},
		"Repeat max N times": {
			Input:  "Repeat Limit - OPRE 7051 may only be repeated a maximum of 6 times",
			Result: rules.NewRepeatRule(6, 0, []utils.Course{{Prefix: "OPRE", Number: "7051"}}, ""),
		},
		"Repeat max 3 times": {
			Input:  "Repeat Limit - UNIV 4076 may only be repeated a maximum of 3 times",
			Result: rules.NewRepeatRule(3, 0, []utils.Course{{Prefix: "UNIV", Number: "4076"}}, ""),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[*rules.RepeatRule](t, tc.Input, rule((*parser.RequirementsParser).Repeat_limit_times_rule), tc.Result)
		})
	}
}

func TestVisitGpaRepeatRule(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result *rules.GpaRepeatRule
	}{
		"Basic": {
			Input:  "GPA Repeat Restriction - MIS 6309",
			Result: rules.NewGpaRepeatRule(utils.Course{Prefix: "MIS", Number: "6309"}),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[*rules.GpaRepeatRule](t, tc.Input, rule((*parser.RequirementsParser).Gpa_repeate_rule), tc.Result)
		})
	}
}

func TestVisitDegreeSatisfactionRule(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result *rules.DegreeSatisfactionRule
	}{
		"Prefix degree satisfaction": {
			Input:  "May not be used to satisfy BS INTS degree requirements",
			Result: rules.NewDegreeSatisfactionRuleFromPrefix([]string{"INTS"}, utils.Undergraduate),
		},
		"Named degree satisfaction": {
			Input:  "May not be used to fulfill degree requirements in MS Information Technology and Management",
			Result: rules.NewDegreeSatisfactionRuleFromDegree([]string{"Information Technology and Management"}, utils.Graduate),
		},
		"Multi prefix for degree satisfaction": {
			Input:  "May not be used to satisfy degree requirements for MS CS or the MS SE degree plans",
			Result: rules.NewDegreeSatisfactionRuleFromPrefix([]string{"CS", "SE"}, utils.Graduate),
		},
		"Of multi prefix degree satisfaction": {
			Input:  "May not be used to satisfy the degree requirements of the MS CS or the MS SE degree plans",
			Result: rules.NewDegreeSatisfactionRuleFromPrefix([]string{"CS", "SE"}, utils.Graduate),
		},
		"School degree satisfaction": {
			Input:  "May not be used to satisfy degree requirements for the School of Engineering and Computer Science",
			Result: rules.NewDegreeSatisfactionRuleFromSchool([]string{"Engineering and Computer Science"}, utils.AnyDegree),
		},
		"School degree satisfaction with majors in": {
			Input:  "May not be used to satisfy degree requirements for majors in the School of Engineering and Computer Science",
			Result: rules.NewDegreeSatisfactionRuleFromSchool([]string{"Engineering and Computer Science"}, utils.AnyDegree),
		},
		"Schools degree satisfaction named list": {
			Input:  "May not be used to satisfy degree requirements for majors in Computer Engineering, Computer Science, and Software Engineering",
			Result: rules.NewDegreeSatisfactionRuleFromSchool([]string{"Computer Engineering", "Computer Science", "Software Engineering"}, utils.AnyDegree),
		},
		"Schools degree satisfaction with level": {
			Input:  "May not be used to satisfy degree requirements for BS majors in Schools of Engineering and Computer Science or Natural Sciences and Mathematics",
			Result: rules.NewDegreeSatisfactionRuleFromSchool([]string{"Engineering and Computer Science", "Natural Sciences and Mathematics"}, utils.Undergraduate),
		},
		"Student degree satisfaction": {
			Input:  "May not be used to satisfy degree requirements by students in Mathematics",
			Result: rules.NewDegreeSatisfactionRuleFromDegree([]string{"Mathematics"}, utils.AnyDegree),
		},
		"Math degree satisfaction": {
			Input:  "May not be used to satisfy mathematics requirements by students in Mathematics",
			Result: rules.NewMathDegreeSatisfactionRule(),
		},
		"Electives degree satisfaction": {
			Input:  "May not be used to satisfy mathematics requirements by students in Mathematics and may not be used to satisfy electives",
			Result: rules.NewDegreeSatisfactionRuleFromElectives(rules.NewMathDegreeSatisfactionRule()),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[*rules.DegreeSatisfactionRule](t, tc.Input, rule((*parser.RequirementsParser).Degree_satisfaction_rule), tc.Result)
		})
	}
}

func TestVisitLivingLearningRule(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result *rules.LivingLearningRule
	}{
		"Prefix": {
			Input:  "ARHM & ATEC Living Learning Community",
			Result: rules.NewLivingLearningRuleFromPrefixes([]string{"ARHM", "ATEC"}),
		},
		"Degree": {
			Input:  "Engineering or Computer Science Living Learning Community",
			Result: rules.NewLivingLearningRuleFromDegrees([]string{"Engineering", "Computer Science"}),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[*rules.LivingLearningRule](t, tc.Input, rule((*parser.RequirementsParser).Living_learning_rule), tc.Result)
		})
	}
}

func TestVisitSchoolRule(t *testing.T) {

	testCases := map[string]struct {
		Input  string
		Result *rules.SchoolRule
	}{
		"Basic": {
			Input: "  Open to students in the School of Engineering and Computer Science, Actuarial Science, Data Science, and Cognitive Science only",
			Result: rules.NewSchoolRule([]string{
				"Engineering and Computer Science",
				"Actuarial Science",
				"Data Science",
				"Cognitive Science",
			}),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[*rules.SchoolRule](t, tc.Input, rule((*parser.RequirementsParser).School_rule), tc.Result)
		})
	}

}
