package visitors

import (
	"parser/objects/constants"
	rules2 "parser/objects/rules"
	"parser/parser"
	"testing"
)

func TestVisitRepeatRule(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result any
	}{
		"Single course repeat restriction": {
			Input: "ACCT 2301 Repeat Restriction",
			Result: rules2.NewCourseRepeatRule([]constants.Course{
				{Prefix: "ACCT", Number: "2301"},
			}),
		},
		"Dual prefix course repeat restriction": {
			Input: "CE/CS 2305 Repeat Restriction",
			Result: rules2.NewCourseRepeatRule(
				[]constants.Course{{Prefix: "CE", Number: "2305"}, {Prefix: "CS", Number: "2305"}},
			),
		},
		"Internship repeat restriction": {
			Input:  "BBSC Internship Repeat Restriction",
			Result: rules2.NewInternshipRepeatRule("BBSC"),
		},
		"Bare repeat restriction": {
			Input:  "Repeat Restriction",
			Result: rules2.NewRepeatRule(1, 0, []constants.Course{}, ""),
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
		Result *rules2.RepeatRule
	}{
		"Repeat max hours": {
			Input:  "Repeat Limit - ACCT 7323 may only be repeated for a maximum of 9 semester credit hours",
			Result: rules2.NewRepeatRule(0, 9, []constants.Course{{Prefix: "ACCT", Number: "7323"}}, ""),
		},
		"Repeat hours max suffix": {
			Input:  "Repeat Limit - EEGR 6V88 may only be repeated for 9 semester credit hours maximum",
			Result: rules2.NewRepeatRule(0, 9, []constants.Course{{Prefix: "EEGR", Number: "6V88"}}, ""),
		},
		"Course repeat max hours may": {
			Input:  "AHST 6336 Repeat Limit - This course may only be repeated for a maximum of 6 semester credit hours",
			Result: rules2.NewRepeatRule(0, 6, []constants.Course{{Prefix: "AHST", Number: "6336"}}, ""),
		},
		"Course repeat max hours can": {
			Input:  "RHET 1302 Repeat Limit - This course can only be repeated for a maximum of 3 semester credit hours",
			Result: rules2.NewRepeatRule(0, 3, []constants.Course{{Prefix: "RHET", Number: "1302"}}, ""),
		},
		"Combined repeat max hours": {
			Input:  "Repeat Limit - CS 4V96 and CS 4V98 combined may only be repeated for a maximum of 6 semester credit hours",
			Result: rules2.NewRepeatRule(0, 6, []constants.Course{{Prefix: "CS", Number: "4V96"}, {Prefix: "CS", Number: "4V98"}}, ""),
		},
		"Course repeat limit": {
			Input:  "ANGM 2309 Repeat Limit",
			Result: rules2.NewCourseRepeatRule([]constants.Course{{Prefix: "ANGM", Number: "2309"}}),
		},
		"Course repeat limit variable number": {
			Input:  "HLTH 4V01 Repeat Limit",
			Result: rules2.NewCourseRepeatRule([]constants.Course{{Prefix: "HLTH", Number: "4V01"}}),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[*rules2.RepeatRule](t, tc.Input, rule((*parser.RequirementsParser).Repeat_limit_hours_rule), tc.Result)
		})
	}
}

func TestVisitRepeatLimitTimesRule(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result *rules2.RepeatRule
	}{
		"Repeat up to N times": {
			Input:  "Repeat Limit - UNIV 4074 may be repeated up to 3 times",
			Result: rules2.NewRepeatRule(3, 0, []constants.Course{{Prefix: "UNIV", Number: "4074"}}, ""),
		},
		"Repeat max N times": {
			Input:  "Repeat Limit - OPRE 7051 may only be repeated a maximum of 6 times",
			Result: rules2.NewRepeatRule(6, 0, []constants.Course{{Prefix: "OPRE", Number: "7051"}}, ""),
		},
		"Repeat max 3 times": {
			Input:  "Repeat Limit - UNIV 4076 may only be repeated a maximum of 3 times",
			Result: rules2.NewRepeatRule(3, 0, []constants.Course{{Prefix: "UNIV", Number: "4076"}}, ""),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[*rules2.RepeatRule](t, tc.Input, rule((*parser.RequirementsParser).Repeat_limit_times_rule), tc.Result)
		})
	}
}

func TestVisitGpaRepeatRule(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result *rules2.GpaRepeatRule
	}{
		"Basic": {
			Input:  "GPA Repeat Restriction - MIS 6309",
			Result: rules2.NewGpaRepeatRule(constants.Course{Prefix: "MIS", Number: "6309"}),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[*rules2.GpaRepeatRule](t, tc.Input, rule((*parser.RequirementsParser).Gpa_repeate_rule), tc.Result)
		})
	}
}

func TestVisitDegreeSatisfactionRule(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result *rules2.DegreeSatisfactionRule
	}{
		"Prefix degree satisfaction": {
			Input:  "May not be used to satisfy BS INTS degree requirements",
			Result: rules2.NewDegreeSatisfactionRuleFromPrefix([]string{"INTS"}, constants.Undergraduate),
		},
		"Named degree satisfaction": {
			Input:  "May not be used to fulfill degree requirements in MS Information Technology and Management",
			Result: rules2.NewDegreeSatisfactionRuleFromDegree([]string{"Information Technology and Management"}, constants.Graduate),
		},
		"Multi prefix for degree satisfaction": {
			Input:  "May not be used to satisfy degree requirements for MS CS or the MS SE degree plans",
			Result: rules2.NewDegreeSatisfactionRuleFromPrefix([]string{"CS", "SE"}, constants.Graduate),
		},
		"Of multi prefix degree satisfaction": {
			Input:  "May not be used to satisfy the degree requirements of the MS CS or the MS SE degree plans",
			Result: rules2.NewDegreeSatisfactionRuleFromPrefix([]string{"CS", "SE"}, constants.Graduate),
		},
		"School degree satisfaction": {
			Input:  "May not be used to satisfy degree requirements for the School of Engineering and Computer Science",
			Result: rules2.NewDegreeSatisfactionRuleFromSchool([]string{"Engineering and Computer Science"}, constants.AnyDegree),
		},
		"School degree satisfaction with majors in": {
			Input:  "May not be used to satisfy degree requirements for majors in the School of Engineering and Computer Science",
			Result: rules2.NewDegreeSatisfactionRuleFromSchool([]string{"Engineering and Computer Science"}, constants.AnyDegree),
		},
		"Schools degree satisfaction named list": {
			Input:  "May not be used to satisfy degree requirements for majors in Computer Engineering, Computer Science, and Software Engineering",
			Result: rules2.NewDegreeSatisfactionRuleFromSchool([]string{"Computer Engineering", "Computer Science", "Software Engineering"}, constants.AnyDegree),
		},
		"Schools degree satisfaction with level": {
			Input:  "May not be used to satisfy degree requirements for BS majors in Schools of Engineering and Computer Science or Natural Sciences and Mathematics",
			Result: rules2.NewDegreeSatisfactionRuleFromSchool([]string{"Engineering and Computer Science", "Natural Sciences and Mathematics"}, constants.Undergraduate),
		},
		"Student degree satisfaction": {
			Input:  "May not be used to satisfy degree requirements by students in Mathematics",
			Result: rules2.NewDegreeSatisfactionRuleFromDegree([]string{"Mathematics"}, constants.AnyDegree),
		},
		"Math degree satisfaction": {
			Input:  "May not be used to satisfy mathematics requirements by students in Mathematics",
			Result: rules2.NewMathDegreeSatisfactionRule(),
		},
		"Electives degree satisfaction": {
			Input:  "May not be used to satisfy mathematics requirements by students in Mathematics and may not be used to satisfy electives",
			Result: rules2.NewDegreeSatisfactionRuleFromElectives(rules2.NewMathDegreeSatisfactionRule()),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[*rules2.DegreeSatisfactionRule](t, tc.Input, rule((*parser.RequirementsParser).Degree_satisfaction_rule), tc.Result)
		})
	}
}

func TestVisitLivingLearningRule(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result *rules2.LivingLearningRule
	}{
		"Prefix": {
			Input:  "ARHM & ATEC Living Learning Community",
			Result: rules2.NewLivingLearningRuleFromPrefixes([]string{"ARHM", "ATEC"}),
		},
		"Degree": {
			Input:  "Engineering or Computer Science Living Learning Community",
			Result: rules2.NewLivingLearningRuleFromDegrees([]string{"Engineering", "Computer Science"}),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[*rules2.LivingLearningRule](t, tc.Input, rule((*parser.RequirementsParser).Living_learning_rule), tc.Result)
		})
	}
}

func TestVisitCreditForRule(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result *rules2.CreditForRule
	}{
		"Both simple courses": {
			Input: "Credit cannot be received for both FIN 3300 and FIN 3330",
			Result: rules2.NewCreditForRule(
				rules2.NewAndCourseCollection(
					rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "FIN", Number: "3300"}}),
					rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "FIN", Number: "3330"}}),
				),
			),
		},
		"Both courses with courses prefix": {
			Input: "Credit cannot be received for both courses, ENTP 6380 and MKT 6380",
			Result: rules2.NewCreditForRule(
				rules2.NewAndCourseCollection(
					rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "ENTP", Number: "6380"}}),
					rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "MKT", Number: "6380"}}),
				),
			),
		},
		"Single course and paren group": {
			Input: "Credit cannot be received for both CS 2337 and (CS 2336 or CE 2336)",
			Result: rules2.NewCreditForRule(
				rules2.NewAndCourseCollection(
					rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "CS", Number: "2337"}}),
					rules2.NewOrCourseCollection(
						rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "CS", Number: "2336"}}),
						rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "CE", Number: "2336"}}),
					),
				),
			),
		},
		"Both paren groups": {
			Input: "Credit cannot be received for both: (OPRE 6301 or SYSM 6303) and (OPRE 6359 or BUAN 6359)",
			Result: rules2.NewCreditForRule(
				rules2.NewAndCourseCollection(
					rules2.NewOrCourseCollection(
						rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "OPRE", Number: "6301"}}),
						rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "SYSM", Number: "6303"}}),
					),
					rules2.NewOrCourseCollection(
						rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "OPRE", Number: "6359"}}),
						rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "BUAN", Number: "6359"}}),
					),
				),
			),
		},
		"More than one of flat list": {
			Input: "Credit cannot be received for more than one of the following: BMEN 1100 or CE 1100 or CS 1200 or EE 1100 or MECH 1100",
			Result: rules2.NewCreditForRule(
				rules2.NewOrCourseCollection(
					rules2.NewOrCourseCollection(
						rules2.NewOrCourseCollection(
							rules2.NewOrCourseCollection(
								rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "BMEN", Number: "1100"}}),
								rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "CE", Number: "1100"}}),
							),
							rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "CS", Number: "1200"}}),
						),
						rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "EE", Number: "1100"}}),
					),
					rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "MECH", Number: "1100"}}),
				),
			),
		},
		"Multiple paren groups and bare course": {
			Input: "Credit cannot be received for more than one of the following: (ACCT 6320 or MIS 6320 or OPRE 6393) and (BUAN 6320 or ACCT 6321) and MIS 6326",
			Result: rules2.NewCreditForRule(
				rules2.NewAndCourseCollection(
					rules2.NewAndCourseCollection(
						rules2.NewOrCourseCollection(
							rules2.NewOrCourseCollection(
								rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "ACCT", Number: "6320"}}),
								rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "MIS", Number: "6320"}}),
							),
							rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "OPRE", Number: "6393"}}),
						),
						rules2.NewOrCourseCollection(
							rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "BUAN", Number: "6320"}}),
							rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "ACCT", Number: "6321"}}),
						),
					),
					rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "MIS", Number: "6326"}}),
				),
			),
		},
		"Ampersand instead of and": {
			Input: "Credit cannot be received for both courses, (CS 3341 or SE 3341 or STAT 3341) & ENGR 3341",
			Result: rules2.NewCreditForRule(
				rules2.NewAndCourseCollection(
					rules2.NewOrCourseCollection(
						rules2.NewOrCourseCollection(
							rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "CS", Number: "3341"}}),
							rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "SE", Number: "3341"}}),
						),
						rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "STAT", Number: "3341"}}),
					),
					rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "ENGR", Number: "3341"}}),
				),
			),
		},
		"Nested and inside paren": {
			Input: "Credit cannot be received for both courses, (CS 1336 and CS 1136) and CS 1436",
			Result: rules2.NewCreditForRule(
				rules2.NewAndCourseCollection(
					rules2.NewAndCourseCollection(
						rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "CS", Number: "1336"}}),
						rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "CS", Number: "1136"}}),
					),
					rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "CS", Number: "1436"}}),
				),
			),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[*rules2.CreditForRule](t, tc.Input, rule((*parser.RequirementsParser).Credit_for_rule), tc.Result)
		})
	}
}

func TestVisitSchoolRule(t *testing.T) {

	testCases := map[string]struct {
		Input  string
		Result *rules2.SchoolRule
	}{
		"Basic": {
			Input: "  Open to students in the School of Engineering and Computer Science, Actuarial Science, Data Science, and Cognitive Science only",
			Result: rules2.NewSchoolRule([]string{
				"Engineering and Computer Science",
				"Actuarial Science",
				"Data Science",
				"Cognitive Science",
			}),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[*rules2.SchoolRule](t, tc.Input, rule((*parser.RequirementsParser).School_rule), tc.Result)
		})
	}

}

func TestVisitSameAsRule(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result *rules2.SameAsRule
	}{
		"Basic": {
			Input:  "(Same as MATH 3335 and STAT 3335)",
			Result: rules2.NewSameAsRule([]constants.Course{{Prefix: "MATH", Number: "3335"}, {Prefix: "STAT", Number: "3335"}}),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[*rules2.SameAsRule](t, tc.Input, rule((*parser.RequirementsParser).Same_as_rule), tc.Result)
		})
	}
}
