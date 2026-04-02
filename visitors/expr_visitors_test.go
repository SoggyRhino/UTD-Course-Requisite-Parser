package visitors

import (
	"parser/conditions"
	"parser/parser"
	"parser/rules"
	"parser/utils"
	"testing"
)

func TestVisitExpr(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result any
	}{
		"ParenExpr": {
			Input:  "(ACN 6340)",
			Result: conditions.NewCourseCondition("ACN", "6340", ""),
		},
		"OrExpr": {
			Input: "MATH 1325 or MATH 2413",
			Result: conditions.NewOrCondition(
				conditions.NewCourseCondition("MATH", "1325", ""),
				conditions.NewCourseCondition("MATH", "2413", ""),
			),
		},
		"AndExpr": {
			Input: "ACCT 6301 and ACCT 6202",
			Result: conditions.NewAndConditionFromExpr(
				conditions.NewCourseCondition("ACCT", "6301", ""),
				conditions.NewCourseCondition("ACCT", "6202", ""),
			),
		},
		"AmpersandExpr": {
			Input: "RHET 1302 & junior standing",
			Result: conditions.NewAndConditionFromExpr(
				conditions.NewCourseCondition("RHET", "1302", ""),
				conditions.NewGradeLevelCondition(utils.Junior),
			),
		},
		"EquivalentExpr": {
			Input:  "ANGM 2310 with a grade of B or better or equivalent",
			Result: conditions.NewAlternativeCondition(conditions.NewCourseCondition("ANGM", "2310", "B")),
		},
		"ConsentExpr": {
			Input:  "instructor consent",
			Result: conditions.NewConsentCondition(utils.InstructorConsent),
		},
		"StandingExpr": {
			Input:  "good academic standing",
			Result: conditions.NewGenericStandingCondition(utils.GoodAcademicStanding),
		},
		"GpaExpr": {
			Input:  "A university grade point average of at least 2.750",
			Result: conditions.NewGpaCondition(2.750),
		},
		"GroupExpr": {
			Input:  "SCVG Student Group",
			Result: conditions.NewStudentGroupCondition(utils.SCVG),
		},
		"ConcurrentEnrollmentExpr": {
			Input: "concurrent enrollment in MATH 1314",
			Result: conditions.NewAndCondition(
				conditions.NewConcurrentEnrollmentCondition(utils.Course{Prefix: "MATH", Number: "1314"}),
			),
		},
		"GradeExpr": {
			Input:  "ACCT 2301 with a minimum grade of C",
			Result: conditions.NewCourseCondition("ACCT", "2301", "C"),
		},
		"AlternativeExpr": {
			Input:  "BIOL 2311 or equivalent",
			Result: conditions.NewAlternativeCondition(conditions.NewCourseCondition("BIOL", "2311", "")),
		},
		"GradeLevelStandingExpr": {
			Input:  "junior standing",
			Result: conditions.NewGradeLevelCondition(utils.Junior),
		},
		"GraduateStandingExpr": {
			Input:  "Graduate Level Standing",
			Result: conditions.NewGraduateStandingInCondition(),
		},
		"MajorExpr": {
			Input:  "Data Science major",
			Result: conditions.NewMajorCondition("Data Science"),
		},
		"DegreeExpr": {
			Input:  "an undergraduate degree in Accounting and adequate foundation/academic performance in a corresponding area",
			Result: conditions.NewDegreeCondition("Accounting"),
		},
		"CoreExpr": {
			Input:  "Completion of 040 Core",
			Result: conditions.NewCoreCondition("040", ""),
		},
		"AnyCoreExpr": {
			Input:  "any 3 semester credit hour 040 core course",
			Result: conditions.NewCoreConditionWithSemesterHours("040", "", 3),
		},
		"CompleteNExpr": {
			Input: "Completion of two of the following: ANGM 2303 or ANGM 2309 or ATCM 2345",
			Result: conditions.NewNCoursesCondition(2, []utils.Course{
				{Prefix: "ANGM", Number: "2303"},
				{Prefix: "ANGM", Number: "2309"},
				{Prefix: "ATCM", Number: "2345"},
			}),
		},
		"SemesterCreditHoursExpr": {
			Input:  "90 semester credit hours",
			Result: conditions.NewCreditHoursCondition(90),
		},
		"MinimumHoursExpr": {
			Input: "Minimum of 6 semester credit hours in any combination of DANC 2332 or DANC 2334",
			Result: conditions.NewCreditHoursFromCondition(6, []utils.Course{
				{Prefix: "DANC", Number: "2332"},
				{Prefix: "DANC", Number: "2334"},
			}),
		},
		"UpperDivisionHoursExpr": {
			Input:  "6 SCH of upper-division ARTS courses",
			Result: conditions.NewUpperDivisionCreditHoursCondition(6, "ARTS"),
		},
		"UpperDivisionClassesExpr": {
			Input:  "a 4000-level HIST course",
			Result: conditions.NewUpperDivisionCountCondition(1, "HIST"),
		},
		"ResearchExpr": {
			Input:  "at least 3 semester credit hours of undergraduate research",
			Result: conditions.NewResearchCondition(3, utils.Undergraduate),
		},
		"PlacementTestExpr": {
			Input:  "CS Placement Test 70 or higher",
			Result: conditions.NewPlacementTestScoreCondition("CS Placement Test", 70, 100),
		},
		"ApScoreExpr": {
			Input:  "AP score of at least 4",
			Result: conditions.NewAPScoreCondition(4),
		},
		"AleksScoreExpr": {
			Input:  "A score of 35% on ALEKS math placement exam",
			Result: conditions.NewAleksScoreCondition(35),
		},
		"ExactSectionExpr": {
			Input: "BIOL 2311.001",
			Result: conditions.NewAndCondition(
				conditions.NewConcurrentEnrollmentCondition(utils.Course{Prefix: "BIOL", Number: "2311", Section: "001"}),
			),
		},
		"AnyMajorCourseExpr": {
			Input:  "any previous PHIL course",
			Result: conditions.NewAnyPreviousMajorCourseCondition("PHIL"),
		},
		"LivingLearningExpr": {
			Input:  "Computer Science Living Learning Community",
			Result: rules.NewLivingLearningRuleFromDegrees([]string{"Computer Science"}),
		},
		"RepeatRuleExpr": {
			Input:  "Repeat Restriction",
			Result: rules.NewRepeatRule(1, 0, []utils.Course{}, ""),
		},
		"RepeatLimitHoursExpr": {
			Input:  "Repeat Limit - HLTL 4304 may only be repeated for a maximum of 6 semester credit hours",
			Result: rules.NewRepeatRule(0, 6, []utils.Course{{Prefix: "HLTL", Number: "4304"}}, ""),
		},
		"CourseExpr": {
			Input:  "ACCT 2301",
			Result: conditions.NewCourseCondition("ACCT", "2301", ""),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[any](t, tc.Input, rule((*parser.RequirementsParser).Expr), tc.Result)
		})
	}
}
