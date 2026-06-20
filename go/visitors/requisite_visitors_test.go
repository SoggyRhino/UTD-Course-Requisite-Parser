package visitors

import (
	"parser/objects"
	conditions2 "parser/objects/conditions"
	"parser/objects/constants"
	rules2 "parser/objects/rules"
	"parser/parser"
	"testing"
)

func TestRequisiteVisitors(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result objects.Requirements
	}{
		"First": {
			Input: "Prerequisite: ACN 6340 or HCS 6340.",
			Result: objects.Requirements{
				PreReqs: conditions2.NewOrCondition(
					conditions2.NewCourseCondition("ACN", "6340", ""),
					conditions2.NewCourseCondition("HCS", "6340", ""),
				),
			},
		},
		"Exclude Notice Req": {
			Input: "BLAW 2301 Repeat Restriction and non-DMHP/non-LLC (DMLC, DFLC, DHLC) Student Group Only",
			Result: objects.Requirements{
				Rules: []rules2.Rule{
					rules2.NewCourseRepeatRule([]constants.Course{
						{Prefix: "BLAW", Number: "2301"},
					}),
				},
				Notices: []constants.Notice{constants.ExcludeDMHPLLCNotice},
			},
		},
		"Append Academic Plan Req": {
			Input: "Prerequisite: BUAN 6341 and Academic Plan Not Equal to BSANMSNF",
			Result: objects.Requirements{
				PreReqs: conditions2.NewAndCondition(
					conditions2.NewAcademicYearCondition("BSANMSNF", false),
					conditions2.NewCourseCondition("BUAN", "6341", ""),
				),
			},
		},
		"Academic Plan Req": {
			Input: "Academic Plan Not Equal to BSANMSNF",
			Result: objects.Requirements{
				PreReqs: conditions2.NewAcademicYearCondition("BSANMSNF", false),
			},
		},
		"Exact Coreq Notice Req": {
			Input: "Check class notes to make sure you are selecting the matching corequisite section",
			Result: objects.Requirements{
				Notices: []constants.Notice{constants.ExactCoReqNotice},
			},
		},
		"Computer Scholars Req": {
			Input: "Computing Scholars Program",
			Result: objects.Requirements{
				Notices: []constants.Notice{constants.ComputerScholarsProgramNotice},
			},
		},
		"GPA Repeat Req": {
			Input: "GPA Repeat Restriction - MIS 6309",
			Result: objects.Requirements{
				Rules: []rules2.Rule{rules2.NewGpaRepeatRule(constants.Course{Prefix: "MIS", Number: "6309"})},
			},
		},
		"Repeat Limit Hours Req": {
			Input: "Repeat Limit - ACCT 7323 may only be repeated for a maximum of 9 semester credit hours",
			Result: objects.Requirements{
				Rules: []rules2.Rule{rules2.NewRepeatRule(0, 9, []constants.Course{{Prefix: "ACCT", Number: "7323"}}, "")},
			},
		},
		"Repeat Limit Times Req": {
			Input: "Repeat Limit - OPRE 7051 may only be repeated a maximum of 6 times",
			Result: objects.Requirements{
				Rules: []rules2.Rule{
					rules2.NewRepeatRule(6, 0, []constants.Course{{Prefix: "OPRE", Number: "7051"}}, ""),
				},
			},
		},
		"Repeat Req": {
			Input: "ACCT 2301 Repeat Restriction",
			Result: objects.Requirements{
				Rules: []rules2.Rule{
					rules2.NewCourseRepeatRule([]constants.Course{{Prefix: "ACCT", Number: "2301"}}),
				},
			},
		},
		"Degree Satisfaction Req": {
			Input: "May not be used to satisfy BS INTS degree requirements",
			Result: objects.Requirements{
				Rules: []rules2.Rule{
					rules2.NewDegreeSatisfactionRuleFromPrefix([]string{"INTS"}, constants.Undergraduate),
				},
			},
		},
		"Credit For Req": {
			Input: "Credit cannot be received for both BPS 6310 and ENTP 6310",
			Result: objects.Requirements{
				Rules: []rules2.Rule{rules2.NewCreditForRule(
					rules2.NewAndCourseCollection(
						rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "BPS", Number: "6310"}}),
						rules2.NewSimpleCourseCollection([]constants.Course{{Prefix: "ENTP", Number: "6310"}}),
					))},
			},
		},
		"Living Learning Req": {
			Input: "ARHM & ATEC Living Learning Community",
			Result: objects.Requirements{
				Rules: []rules2.Rule{rules2.NewLivingLearningRuleFromPrefixes([]string{"ARHM", "ATEC"})},
			},
		},
		"School Req": {
			Input: "Open to students in the School of Engineering and Computer Science, Actuarial Science, Data Science, and Cognitive Science only",
			Result: objects.Requirements{
				Rules: []rules2.Rule{rules2.NewSchoolRule([]string{
					"Engineering and Computer Science",
					"Actuarial Science",
					"Data Science",
					"Cognitive Science",
				})},
			},
		},
		"Major Req": {
			Input: "ENCS majors only",
			Result: objects.Requirements{
				PreReqs: conditions2.NewMajorCondition("ENCS"),
			},
		},
		"Prereq Req": {
			Input: "Prerequisite: ACCT 2301",
			Result: objects.Requirements{
				PreReqs: conditions2.NewCourseCondition("ACCT", "2301", ""),
			},
		},
		"Coreq Req": {
			Input: "Corequisite: BIOL 2311.001",
			Result: objects.Requirements{
				CoReqs: conditions2.NewAndCondition(
					conditions2.NewExactSectionCondition(constants.Course{Prefix: "BIOL", Number: "2311", Section: "001"}),
				),
			},
		},
		"Prereq and Coreq Req": {
			Input: "Prerequisite: Collegium V Honors Student Group and Corequisite: CHEM 1315",
			Result: objects.Requirements{
				PreReqs: conditions2.NewStudentGroupCondition(constants.CollegeVHonors),
				CoReqs:  conditions2.NewCourseCondition("CHEM", "1315", ""),
			},
		},
		"Pre or Co Req": {
			Input: "Prerequisite or Corequisite: ACCT 3332",
			Result: objects.Requirements{
				PreOrCoReqs: conditions2.NewCourseCondition("ACCT", "3332", ""),
			},
		},
		"Same As Req": {
			Input: "(Same as MATH 3335 and STAT 3335)",
			Result: objects.Requirements{
				Rules: []rules2.Rule{rules2.NewSameAsRule([]constants.Course{{Prefix: "MATH", Number: "3335"}, {Prefix: "STAT", Number: "3335"}})},
			},
		},
		"Expr Req": {
			Input: "MECO 3340",
			Result: objects.Requirements{
				PreReqs: conditions2.NewCourseCondition("MECO", "3340", ""),
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[objects.Requirements](t, tc.Input, rule((*parser.RequirementsParser).Prog), tc.Result)
		})
	}
}
