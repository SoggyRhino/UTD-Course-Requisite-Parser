package visitors

import (
	"parser/conditions"
	"parser/constants"
	"parser/parser"
	"testing"
)

func TestVisitConsentCondition(t *testing.T) {
	//nothingburger but I want 100% coverage
	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"Instructor consent": {
			Input:  "instructor consent",
			Result: conditions.NewConsentCondition(constants.InstructorConsent),
		},
		"Instructor consent required": {
			Input:  "instructor consent required",
			Result: conditions.NewConsentCondition(constants.InstructorConsent),
		},
		"Department consent required": {
			Input:  "department consent required",
			Result: conditions.NewConsentCondition(constants.DepartmentConsent),
		},
		"UTeach advisor consent required": {
			Input:  "UTeach advisor consent required",
			Result: conditions.NewConsentCondition(constants.UTeachConsent),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Consent_condition), tc.Result)
		})
	}
}

func TestVisitStandingCondition(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"Upper-division standing": {
			Input:  "Upper-division standing",
			Result: conditions.NewGenericStandingCondition(constants.UpperDivisionStanding),
		},
		"Good academic standing": {
			Input:  "good academic standing",
			Result: conditions.NewGenericStandingCondition(constants.GoodAcademicStanding),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Standing_condition), tc.Result)
		})
	}
}

func TestVisitSimpleGradeCondition(t *testing.T) {

	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"Basic": {
			Input: "ACCT 2301 with a minimum grade of C",
			Result: conditions.NewCourseCondition(
				"ACCT", "2301", "C"),
		},
		"Or Better": {
			Input:  "ANGM 3370 with a grade of B or higher",
			Result: conditions.NewCourseCondition("ANGM", "3370", "B"),
		},
		"Short Hand Course": {
			Input: "CS/CE 2305 with a minimum grade of C",
			Result: conditions.NewOrCondition(
				conditions.NewCourseCondition("CS", "2305", "C"),
				conditions.NewCourseCondition("CE", "2305", "C"),
			),
		},
		"GPA": {
			Input:  "MATH 2413 with a grade greater than or equal to C- (1.67)",
			Result: conditions.NewCourseCondition("MATH", "2413", "C-"),
		},
		"Grade List": {
			Input: "(CE 2336 or CS 2336 or CS 2337) with a grade of C or better",
			Result: conditions.NewOrCondition(
				conditions.NewCourseCondition("CE", "2336", "C"),
				conditions.NewCourseCondition("CS", "2336", "C"),
				conditions.NewCourseCondition("CS", "2337", "C"),
			),
		},
		"At least Grade in": {
			Input: "a grade of at least a C- in MATH 1314 and MATH 1316",
			Result: conditions.NewAndCondition(
				conditions.NewCourseCondition("MATH", "1314", "C-"),
				conditions.NewCourseCondition("MATH", "1316", "C-"),
			),
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Grade_condition), tc.Result)
		})
	}

}

func TestVisitAlternateCondition(t *testing.T) {

	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"Single Course": {
			Input: "BIOL 2311 or equivalent",
			Result: conditions.NewAlternativeCondition(
				conditions.NewCourseCondition("BIOL", "2311", ""),
			),
		},

		"Multiple Courses": {
			Input: "PHIL 1301 or PHIL 1305 or PHIL 1306 or equivalent",
			Result: conditions.NewAlternativeCondition(
				conditions.NewOrCondition(
					conditions.NewCourseCondition("PHIL", "1301", ""),
					conditions.NewCourseCondition("PHIL", "1305", ""),
					conditions.NewCourseCondition("PHIL", "1306", ""),
				),
			),
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Alternative_condition), tc.Result)
		})
	}
}

func TestVisitGradeLevelStandingCondition(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"Basic": {
			Input: "Junior or Senior standing",
			Result: conditions.NewOrCondition(
				conditions.NewGradeLevelCondition(constants.Junior),
				conditions.NewGradeLevelCondition(constants.Senior),
			),
		},
		"Senior History Major standing": {
			Input:  "Senior History Major standing",
			Result: conditions.NewGradeLevelConditionWithDegree(constants.Senior, "History"),
		},
		"Minimum of Sophomore standing": {
			Input:  "Minimum of Sophomore standing",
			Result: conditions.NewGradeLevelCondition(constants.Sophomore),
		},
		"At least Senior-level Standing": {
			Input:  "At least Senior-level Standing",
			Result: conditions.NewGradeLevelCondition(constants.Senior),
		},
		"PSY Majors Only with Junior or Senior standing": {
			Input: "PSY Majors Only with Junior or Senior standing",
			Result: conditions.NewOrCondition(
				conditions.NewGradeLevelConditionWithDegree(constants.Junior, "PSY"),
				conditions.NewGradeLevelConditionWithDegree(constants.Senior, "PSY"),
			),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Grade_level_standing_condition), tc.Result)
		})
	}

}

func TestVisitGraduateStandingCondition(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"Graduate standing in Biology": {
			Input:  "Graduate standing in Biology",
			Result: conditions.NewGraduateStandingInConditionWithDegree("Biology"),
		},
		"Graduate standing in Chemistry": {
			Input:  "Graduate standing in chemistry",
			Result: conditions.NewGraduateStandingInConditionWithDegree("chemistry"),
		},
		"Graduate Level Standing": {
			Input:  "Graduate Level Standing",
			Result: conditions.NewGraduateStandingInCondition(),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Graduate_standing_condition), tc.Result)
		})
	}
}

func TestVisitGpaCondition(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"University GPA": {
			Input:  "A university grade point average of at least 2.750",
			Result: conditions.NewGpaCondition(2.750),
		},
		"Minimum GPA requirement": {
			Input:  "Minimum GPA requirement 3.200",
			Result: conditions.NewGpaCondition(3.200),
		},
		"GPA in course": {
			Input:  "a GPA of 3.000 or better in UTeach coursework",
			Result: conditions.NewGpaConditionWithDegree(3.000, "UTeach"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Gpa_condition), tc.Result)
		})
	}
}

func TestVisitMajorCondition(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		// prefixMajorCondition
		"Single prefix majors only": {
			Input:  "ENCS majors only",
			Result: conditions.NewMajorCondition("ENCS"),
		},
		"Multiple prefixes": {
			Input: "EE or CE or TE major",
			Result: conditions.NewOrCondition(
				conditions.NewMajorCondition("EE"),
				conditions.NewMajorCondition("CE"),
				conditions.NewMajorCondition("TE"),
			),
		},
		"Prefix with degree type": {
			Input:  "BBSC GRAD majors only",
			Result: conditions.NewMajorConditionWithDegreeLevel("BBSC", constants.Graduate),
		},
		"Prefixes with grade level": {
			Input: "CE or EE Freshman Majors only",
			Result: conditions.NewOrCondition(
				conditions.NewMajorConditionWithGradeLevel("CE", constants.Freshman),
				conditions.NewMajorConditionWithGradeLevel("EE", constants.Freshman),
			),
		},
		"Prefix with degree level": {
			Input:  "ENCS PhD majors only",
			Result: conditions.NewMajorConditionWithDegreeLevel("ENCS", constants.PhD),
		},
		"Grade level prefix major": {
			Input:  "Freshman ENCS Majors only",
			Result: conditions.NewMajorConditionWithGradeLevel("ENCS", constants.Freshman),
		},
		"Degree type prefix major": {
			Input:  "MS ITM Major only",
			Result: conditions.NewMajorConditionWithDegreeLevel("ITM", constants.Graduate),
		},
		"Named major": {
			Input:  "Data Science major",
			Result: conditions.NewMajorCondition("Data Science"),
		},
		"Named major only": {
			Input:  "Visual and Performing Arts Major Only",
			Result: conditions.NewMajorCondition("Visual and Performing Arts"),
		},
		"Named degree type major": {
			Input:  "International Management Studies PhD majors only",
			Result: conditions.NewMajorConditionWithDegreeLevel("International Management Studies", constants.PhD),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Major_condition), tc.Result)
		})
	}
}

func TestVisitDegreeCondition(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"Undergraduate degree in Accounting": {
			Input:  "an undergraduate degree in Accounting and adequate foundation/academic performance in a corresponding area",
			Result: conditions.NewDegreeCondition("Accounting"),
		},
		"Bachelors or Masters degree in subject list": {
			Input: "Bachelor's or Master's degree in psychology or computer science or neuroscience",
			Result: conditions.NewOrCondition(
				conditions.NewDegreeCondition("psychology"),
				conditions.NewDegreeCondition("computer science"),
				conditions.NewDegreeCondition("neuroscience"),
			),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Degree_condition), tc.Result)
		})
	}
}

func TestVisitCoreCondition(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		// coreCondition
		"Completion of core number": {
			Input:  "Completion of 050 Creative Arts Core",
			Result: conditions.NewCoreCondition("050", "Creative Arts"),
		},
		"Completion of bare core number": {
			Input:  "Completion of 040 Core",
			Result: conditions.NewCoreCondition("040", ""),
		},
		"Completion of core course": {
			Input:  "Completion of an 010 core course",
			Result: conditions.NewCoreCondition("010", ""),
		},
		"Completion of parenthesized core": {
			Input:  "Completion of a 060 (American History) core course",
			Result: conditions.NewCoreCondition("060", "American History"),
		},
		"Completion of multiword core": {
			Input:  "Completion of 040 Language, Philosophy and Culture Core",
			Result: conditions.NewCoreCondition("040", "Language, Philosophy and Culture"),
		},
		// anyCoreSCHCondition
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Core_condition), tc.Result)
		})
	}
}

func TestVisitAnyCoreSCHCondition(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"Any SCH core course": {
			Input:  "any 3 semester credit hour 040 core course",
			Result: conditions.NewCoreConditionWithSemesterHours("040", "", 3),
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Any_core_condition), tc.Result)
		})
	}
}

func TestVisitCreditHoursCondition(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		// semesterCreditHoursCondition
		"90 semester credit hours": {
			Input:  "90 semester credit hours",
			Result: conditions.NewCreditHoursCondition(90),
		},
		"45 semester credit hours": {
			Input:  "45 semester credit hours",
			Result: conditions.NewCreditHoursCondition(45),
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Semester_credit_hours_condition), tc.Result)
		})
	}
}

func TestVisitMinimumHoursCondition(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"Minimum of 6 SCH in any combination": {
			Input: "Minimum of 6 semester credit hours in any combination of DANC 2332 or DANC 2334",
			Result: conditions.NewCreditHoursFromCondition(6, []constants.Course{
				{Prefix: "DANC", Number: "2332"},
				{Prefix: "DANC", Number: "2334"},
			}),
		},
		"Minimum of 3 SCH of": {
			Input: "At least 3 semester credits of ECS 1192 or 2192 or 3292",
			Result: conditions.NewCreditHoursFromCondition(3, []constants.Course{
				{Prefix: "ECS", Number: "1192"},
				{Prefix: "ECS", Number: "2192"},
				{Prefix: "ECS", Number: "3292"},
			}),
		},
		"Minimum of 6 SCH from": {
			Input: "6 semester credit hours from the following: LIT 2320 or LIT 2321 or LIT 2322 or LIT 2331",
			Result: conditions.NewCreditHoursFromCondition(6, []constants.Course{
				{Prefix: "LIT", Number: "2320"},
				{Prefix: "LIT", Number: "2321"},
				{Prefix: "LIT", Number: "2322"},
				{Prefix: "LIT", Number: "2331"},
			}),
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Minimum_hours_condition), tc.Result)
		})
	}

}

func TestVisitUpperDivisionHoursCondition(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"SCH of upper-division prefix courses": {
			Input:  "6 SCH of upper-division ARTS courses",
			Result: conditions.NewUpperDivisionCreditHoursCondition(6, "ARTS"),
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Upper_division_hours_condition), tc.Result)
		})
	}
}

func TestVisitUpperDivisionClassesCondition(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"At least three CS 43XX classes": {
			Input:  "at least three CS 43XX classes",
			Result: conditions.NewUpperDivisionCountCondition(3, "CS"),
		},
		"A 4000-level course": {
			Input:  "a 4000-level HIST course",
			Result: conditions.NewUpperDivisionCountCondition(1, "HIST"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Upper_division_classes_condition), tc.Result)
		})
	}
}

func TestVisitResearchCondition(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"Undergraduate research": {
			Input:  "at least 3 semester credit hours of undergraduate research",
			Result: conditions.NewResearchCondition(3, constants.Undergraduate),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Research_condition), tc.Result)
		})
	}
}

func TestVisitCompleteNOfFollowingCondition(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"Completion of two of the following": {
			Input: "Completion of two of the following: ANGM 2303 or ANGM 2309 or ATCM 2345",
			Result: conditions.NewNCoursesCondition(2, []constants.Course{
				{Prefix: "ANGM", Number: "2303"},
				{Prefix: "ANGM", Number: "2309"},
				{Prefix: "ATCM", Number: "2345"},
			}),
		},
		"Two courses from the following": {
			Input: "Two courses from the following - BIOL 5410 or BIOL 5420 or BIOL 5440 or BIOL 5460",
			Result: conditions.NewNCoursesCondition(2, []constants.Course{
				{Prefix: "BIOL", Number: "5410"},
				{Prefix: "BIOL", Number: "5420"},
				{Prefix: "BIOL", Number: "5440"},
				{Prefix: "BIOL", Number: "5460"},
			}),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Complete_n_condition), tc.Result)
		})
	}
}

func TestVisitPlacementTestCondition(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"Placement score less than": {
			Input:  "Arabic Language Placement Test Score less than 20",
			Result: conditions.NewPlacementTestScoreCondition("Arabic Language Placement Test", 0, 20),
		},
		"Placement score greater than": {
			Input:  "Chinese Language Placement score greater than 80",
			Result: conditions.NewPlacementTestScoreCondition("Chinese Language Placement", 80, 100),
		},
		"Placement score range": {
			Input:  "Arabic Language Placement score of 20-39",
			Result: conditions.NewPlacementTestScoreCondition("Arabic Language Placement", 20, 39),
		},
		"Placement score range with Test in name": {
			Input:  "Japanese Language Place Test Score of 51-60",
			Result: conditions.NewPlacementTestScoreCondition("Japanese Language Place Test", 51, 60),
		},
		"Placement score minimum": {
			Input:  "CS Placement Test 70 or higher",
			Result: conditions.NewPlacementTestScoreCondition("CS Placement Test", 70, 100),
		},
		"Placement score minimum with 'a' and 'of'": {
			Input:  "a CS placement test of 70 or higher",
			Result: conditions.NewPlacementTestScoreCondition("CS placement test", 70, 100),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Placement_test_condition), tc.Result)
		})
	}
}

func TestVisitApScoreCondition(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"AP score of at least 4": {
			Input:  "AP score of at least 4",
			Result: conditions.NewAPScoreCondition(4),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Ap_score_condition), tc.Result)
		})
	}
}

func TestVisitAleksScoreCondition(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"Score of 35% on ALEKS": {
			Input:  "A score of 35% on ALEKS math placement exam",
			Result: conditions.NewAleksScoreCondition(35),
		},
		"Minimal placement score of 85% on ALEKS": {
			Input:  "A minimal placement score of 85% on ALEKS math placement exam",
			Result: conditions.NewAleksScoreCondition(85),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Aleks_score_condition), tc.Result)
		})
	}
}

func TestVisitStudentGroupCondition(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"Students in both groups": {
			Input: "Students in both Liberal Arts Honors /Collegium V Honors Group only",
			Result: conditions.NewAndCondition(
				conditions.NewStudentGroupCondition(constants.LiberalArtsHonors),
				conditions.NewStudentGroupCondition(constants.CollegeVHonors),
			),
		},
		"Multiple groups": {
			Input: "SCVG and DLAH student groups only",
			Result: conditions.NewOrCondition(
				conditions.NewStudentGroupCondition(constants.SCVG),
				conditions.NewStudentGroupCondition(constants.DLAH),
			),
		},
		"CV Honors students only": {
			Input:  "CV Honors Students only",
			Result: conditions.NewStudentGroupCondition(constants.CollegeVHonors),
		},
		"SCVG group": {
			Input:  "SCVG Student Group",
			Result: conditions.NewStudentGroupCondition(constants.SCVG),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Group_condition), tc.Result)
		})
	}
}

func TestVisitConcurrentEnrollmentCondition(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"Concurrent enrollment in single course": {
			Input: "concurrent enrollment in MATH 1314",
			Result: conditions.NewAndCondition(
				conditions.NewConcurrentEnrollmentCondition(constants.Course{Prefix: "MATH", Number: "1314"}),
			),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Concurrent_enrollment_condition), tc.Result)
		})
	}
}

func TestVisitExactSectionCondition(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"BIOL 2311.001": {
			Input: "BIOL 2311.001",
			Result: conditions.NewAndCondition(
				conditions.NewConcurrentEnrollmentCondition(constants.Course{Prefix: "BIOL", Number: "2311", Section: "001"}),
			),
		},
		"BIOL 2311.501": {
			Input: "BIOL 2311.501",
			Result: conditions.NewAndCondition(
				conditions.NewConcurrentEnrollmentCondition(constants.Course{Prefix: "BIOL", Number: "2311", Section: "501"}),
			),
		},
		"BIOL 2111.502": {
			Input: "BIOL 2111.502",
			Result: conditions.NewAndCondition(
				conditions.NewConcurrentEnrollmentCondition(constants.Course{Prefix: "BIOL", Number: "2111", Section: "502"}),
			),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Exact_section_condition), tc.Result)
		})
	}
}

func TestVisitWorkshopSectionCondition(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"Workshop section": {
			Input: "BIOL 2112 workshop 501",
			Result: conditions.NewAndCondition(
				conditions.NewConcurrentEnrollmentCondition(constants.Course{Prefix: "BIOL", Number: "2112", Section: "501"}),
			),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Exact_section_condition), tc.Result)
		})
	}
}

func TestVisitAnyPreviousMajorCourseCondition(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"Any previous course lowercase": {
			Input:  "any previous PHIL course",
			Result: conditions.NewAnyPreviousMajorCourseCondition("PHIL"),
		},
		"Any previous course uppercase": {
			Input:  "Any previous PHIL course",
			Result: conditions.NewAnyPreviousMajorCourseCondition("PHIL"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Any_major_course_condition), tc.Result)
		})
	}
}

func TestVisitAcademicPlanCondition(t *testing.T) {
	testCases := map[string]struct {
		Input  string
		Result conditions.Condition
	}{
		"Academic Plan Not Equal": {
			Input:  "Academic Plan Not Equal to BSANMSNF",
			Result: conditions.NewAcademicYearCondition("BSANMSNF", false),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Academic_plan_condition), tc.Result)
		})
	}
}
