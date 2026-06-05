package visitors

import (
	"parser/conditions"
	"parser/constants"
	"parser/parser"
	"parser/rules"
	"testing"
)

func TestRequisiteVisitors(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result Requirements
	}{
		"First": {
			Input: "Prerequisite: ACN 6340 or HCS 6340.",
			Result: Requirements{
				PreReqs: conditions.NewOrCondition(
					conditions.NewCourseCondition("ACN", "6340", ""),
					conditions.NewCourseCondition("HCS", "6340", ""),
				),
			},
		},
		"Exclude Notice Req": {
			Input: "BLAW 2301 Repeat Restriction and non-DMHP/non-LLC (DMLC, DFLC, DHLC) Student Group Only",
			Result: Requirements{
				Rules: []rules.Rule{
					rules.NewCourseRepeatRule([]constants.Course{
						{Prefix: "BLAW", Number: "2301"},
					}),
				},
				Notices: []constants.Notice{constants.ExcludeDMHPLLCNotice},
			},
		},
		"Append Academic Plan Req": {
			Input: "Prerequisite: BUAN 6341 and Academic Plan Not Equal to BSANMSNF",
			Result: Requirements{
				PreReqs: conditions.NewAndCondition(
					conditions.NewAcademicYearCondition("BSANMSNF", false),
					conditions.NewCourseCondition("BUAN", "6341", ""),
				),
			},
		},
		"Academic Plan Req": {
			Input: "Academic Plan Not Equal to BSANMSNF",
			Result: Requirements{
				PreReqs: conditions.NewAcademicYearCondition("BSANMSNF", false),
			},
		},
		"Exact Coreq Notice Req": {
			Input: "Check class notes to make sure you are selecting the matching corequisite section",
			Result: Requirements{
				Notices: []constants.Notice{constants.ExactCoReqNotice},
			},
		},
		"Computer Scholars Req": {
			Input: "Computing Scholars Program",
			Result: Requirements{
				Notices: []constants.Notice{constants.ComputerScholarsProgramNotice},
			},
		},
		"GPA Repeat Req": {
			Input: "GPA Repeat Restriction - MIS 6309",
			Result: Requirements{
				Rules: []rules.Rule{rules.NewGpaRepeatRule(constants.Course{Prefix: "MIS", Number: "6309"})},
			},
		},
		"Repeat Limit Hours Req": {
			Input: "Repeat Limit - ACCT 7323 may only be repeated for a maximum of 9 semester credit hours",
			Result: Requirements{
				Rules: []rules.Rule{rules.NewRepeatRule(0, 9, []constants.Course{{Prefix: "ACCT", Number: "7323"}}, "")},
			},
		},
		"Repeat Limit Times Req": {
			Input: "Repeat Limit - OPRE 7051 may only be repeated a maximum of 6 times",
			Result: Requirements{
				Rules: []rules.Rule{
					rules.NewRepeatRule(6, 0, []constants.Course{{Prefix: "OPRE", Number: "7051"}}, ""),
				},
			},
		},
		"Repeat Req": {
			Input: "ACCT 2301 Repeat Restriction",
			Result: Requirements{
				Rules: []rules.Rule{
					rules.NewCourseRepeatRule([]constants.Course{{Prefix: "ACCT", Number: "2301"}}),
				},
			},
		},
		"Degree Satisfaction Req": {
			Input: "May not be used to satisfy BS INTS degree requirements",
			Result: Requirements{
				Rules: []rules.Rule{
					rules.NewDegreeSatisfactionRuleFromPrefix([]string{"INTS"}, constants.Undergraduate),
				},
			},
		},
		"Credit For Req": {
			Input: "Credit cannot be received for both BPS 6310 and ENTP 6310",
			Result: Requirements{
				Rules: []rules.Rule{rules.NewCreditForRule(
					rules.NewAndCourseCollection(
						rules.NewSimpleCourseCollection([]constants.Course{{Prefix: "BPS", Number: "6310"}}),
						rules.NewSimpleCourseCollection([]constants.Course{{Prefix: "ENTP", Number: "6310"}}),
					))},
			},
		},
		"Living Learning Req": {
			Input: "ARHM & ATEC Living Learning Community",
			Result: Requirements{
				Rules: []rules.Rule{rules.NewLivingLearningRuleFromPrefixes([]string{"ARHM", "ATEC"})},
			},
		},
		"School Req": {
			Input: "Open to students in the School of Engineering and Computer Science, Actuarial Science, Data Science, and Cognitive Science only",
			Result: Requirements{
				Rules: []rules.Rule{rules.NewSchoolRule([]string{
					"Engineering and Computer Science",
					"Actuarial Science",
					"Data Science",
					"Cognitive Science",
				})},
			},
		},
		"Major Req": {
			Input: "ENCS majors only",
			Result: Requirements{
				PreReqs: conditions.NewMajorCondition("ENCS"),
			},
		},
		"Prereq Req": {
			Input: "Prerequisite: ACCT 2301",
			Result: Requirements{
				PreReqs: conditions.NewCourseCondition("ACCT", "2301", ""),
			},
		},
		"Coreq Req": {
			Input: "Corequisite: BIOL 2311.001",
			Result: Requirements{
				CoReqs: conditions.NewAndCondition(
					conditions.NewExactSectionCondition(constants.Course{Prefix: "BIOL", Number: "2311", Section: "001"}),
				),
			},
		},
		"Prereq and Coreq Req": {
			Input: "Prerequisite: Collegium V Honors Student Group and Corequisite: CHEM 1315",
			Result: Requirements{
				PreReqs: conditions.NewStudentGroupCondition(constants.CollegeVHonors),
				CoReqs:  conditions.NewCourseCondition("CHEM", "1315", ""),
			},
		},
		"Pre or Co Req": {
			Input: "Prerequisite or Corequisite: ACCT 3332",
			Result: Requirements{
				PreOrCoReqs: conditions.NewCourseCondition("ACCT", "3332", ""),
			},
		},
		"Same As Req": {
			Input: "(Same as MATH 3335 and STAT 3335)",
			Result: Requirements{
				Rules: []rules.Rule{rules.NewSameAsRule([]constants.Course{{Prefix: "MATH", Number: "3335"}, {Prefix: "STAT", Number: "3335"}})},
			},
		},
		"Expr Req": {
			Input: "MECO 3340",
			Result: Requirements{
				PreReqs: conditions.NewCourseCondition("MECO", "3340", ""),
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[Requirements](t, tc.Input, rule((*parser.RequirementsParser).Prog), tc.Result)
		})
	}
}
