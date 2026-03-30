package visitors

import (
	"parser/conditions"
	"parser/parser"
	"testing"
)

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
				conditions.NewGradeLevelCondition(conditions.Junior),
				conditions.NewGradeLevelCondition(conditions.Senior),
			),
		},
		"Senior History Major standing": {
			Input:  "Senior History Major standing",
			Result: conditions.NewGradeLevelConditionWithDegree(conditions.Senior, "History"),
		},
		"Minimum of Sophomore standing": {
			Input:  "Minimum of Sophomore standing",
			Result: conditions.NewGradeLevelCondition(conditions.Sophomore),
		},
		"At least Senior-level Standing": {
			Input:  "At least Senior-level Standing",
			Result: conditions.NewGradeLevelCondition(conditions.Senior),
		},
		"PSY Majors Only with Junior or Senior standing": {
			Input: "PSY Majors Only with Junior or Senior standing",
			Result: conditions.NewOrCondition(
				conditions.NewGradeLevelConditionWithDegree(conditions.Junior, "PSY"),
				conditions.NewGradeLevelConditionWithDegree(conditions.Senior, "PSY"),
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
			Result: conditions.NewMajorConditionWithDegreeLevel("BBSC", conditions.Graduate),
		},
		"Prefixes with grade level": {
			Input: "CE or EE Freshman Majors only",
			Result: conditions.NewOrCondition(
				conditions.NewMajorConditionWithGradeLevel("CE", conditions.Freshman),
				conditions.NewMajorConditionWithGradeLevel("EE", conditions.Freshman),
			),
		},
		"Prefix with degree level": {
			Input:  "ENCS PhD majors only",
			Result: conditions.NewMajorConditionWithDegreeLevel("ENCS", conditions.PhD),
		},
		"Grade level prefix major": {
			Input:  "Freshman ENCS Majors only",
			Result: conditions.NewMajorConditionWithGradeLevel("ENCS", conditions.Freshman),
		},
		"Degree type prefix major": {
			Input:  "MS ITM Major only",
			Result: conditions.NewMajorConditionWithDegreeLevel("ITM", conditions.Graduate),
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
			Result: conditions.NewMajorConditionWithDegreeLevel("International Management Studies", conditions.PhD),
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
			Result: conditions.NewCreditHoursFromCondition(6, []conditions.Course{
				{Prefix: "DANC", Number: "2332"},
				{Prefix: "DANC", Number: "2334"},
			}),
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testTree[conditions.Condition](t, tc.Input, rule((*parser.RequirementsParser).Minimum_hours_condition), tc.Result)
		})
	}

}
