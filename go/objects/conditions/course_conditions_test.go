package conditions

import (
	"parser/objects/constants"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGradeAtLeast(t *testing.T) {
	testCases := map[string]struct {
		actual   constants.Grade
		min      constants.Grade
		expected bool
	}{
		"Above minimum": {
			actual:   "A",
			min:      "B",
			expected: true,
		},
		"Equal to minimum": {
			actual:   "C",
			min:      "C",
			expected: true,
		},
		"Below minimum": {
			actual:   "C-",
			min:      "C",
			expected: false,
		},
		"Unknown actual grade": {
			actual:   "P",
			min:      "C",
			expected: false,
		},
		"Unknown minimum grade": {
			actual:   "A",
			min:      "P",
			expected: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := gradeAtLeast(tc.actual, tc.min)
			if got != tc.expected {
				t.Errorf("Unexpected grade comparison: got %t, want %t", got, tc.expected)
			}
		})
	}

	t.Run("Unknown actual grade", func(t *testing.T) {
		if gradeAtLeast("P", "A") != false {
			t.Error("Expected false for unknown actual grade")
		}
	})

	t.Run("Unknown minimum grade", func(t *testing.T) {
		if gradeAtLeast("A", "P") != false {
			t.Error("Expected false for unknown minimum grade")
		}
	})
}

func TestCourseConditionConstructors(t *testing.T) {
	NewCoreConditionWithSemesterHours("010", "Title", 3)
	NewUpperDivisionCountCondition(2, "CS")
}

func TestCourseConditionFulfils(t *testing.T) {
	testCases := map[string]struct {
		condition CourseCondition
		userInfo  constants.UserInfo
		expected  constants.Evaluation
	}{
		"Completed without minimum grade": {
			condition: *NewCourseCondition("BIOL", "2311", ""),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{
					{Prefix: "BIOL", Number: "2311"}: "B",
				},
			},
			expected: constants.Evaluation{
				Name:    "BIOL 2311",
				Status:  constants.StatusPass,
				Summary: "Completed BIOL 2311 with grade B",
			},
		},
		"Completed without minimum grade, below C": {
			condition: *NewCourseCondition("BIOL", "2311", ""),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{
					{Prefix: "BIOL", Number: "2311"}: "D",
				},
			},
			expected: constants.Evaluation{
				Name:    "BIOL 2311",
				Status:  constants.StatusDefiniteFail,
				Summary: "Completed BIOL 2311 but grade D does not meet minimum C",
			},
		},
		"Completed with sufficient grade": {
			condition: *NewCourseCondition("BIOL", "2311", "C"),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{
					{Prefix: "BIOL", Number: "2311"}: "B",
				},
			},
			expected: constants.Evaluation{
				Name:    "BIOL 2311",
				Status:  constants.StatusPass,
				Summary: "Completed BIOL 2311 with grade B",
			},
		},
		"Completed with insufficient grade": {
			condition: *NewCourseCondition("BIOL", "2311", "C"),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{
					{Prefix: "BIOL", Number: "2311"}: "D",
				},
			},
			expected: constants.Evaluation{
				Name:    "BIOL 2311",
				Status:  constants.StatusDefiniteFail,
				Summary: "Completed BIOL 2311 but grade D does not meet minimum C",
			},
		},
		"Currently enrolled": {
			condition: *NewCourseCondition("BIOL", "2311", ""),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{
					{Prefix: "CHEM", Number: "1311"}: "A",
				},
				CurrentEnrollment: []constants.Course{{Prefix: "BIOL", Number: "2311"}},
			},
			expected: constants.Evaluation{
				Name:    "BIOL 2311",
				Status:  constants.StatusPossibleFail,
				Summary: "Currently enrolled in BIOL 2311 — awaiting final grade",
			},
		},
		"Not completed": {
			condition: *NewCourseCondition("BIOL", "2311", ""),
			userInfo:  constants.UserInfo{},
			expected: constants.Evaluation{
				Name:    "BIOL 2311",
				Status:  constants.StatusDefiniteFail,
				Summary: "Have not taken BIOL 2311",
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := tc.condition.Fulfils(tc.userInfo, false)
			assertEval(t, tc.expected, got)
		})
	}
}

func TestCourseCondition_ConstructorsAndJSON(t *testing.T) {
	condition := NewCourseCondition("BIOL", "2311", "")
	expected := &CourseCondition{
		Course: constants.Course{
			Prefix: "BIOL",
			Number: "2311",
		},
	}

	if diff := cmp.Diff(expected, condition); diff != "" {
		t.Errorf("Unexpected course condition (-want +got):\n%s", diff)
	}

	condition.AppendGrade("C")
	if condition.MinGrade != "C" {
		t.Errorf("Expected minimum grade %q, got %q", constants.Grade("C"), condition.MinGrade)
	}

	assertJSONRoundTrip[CourseCondition](t, condition, "course")
}

func TestCoreConditionFulfils(t *testing.T) {
	testCases := map[string]struct {
		condition CoreCondition
		expected  constants.Evaluation
	}{
		"Core title": {
			condition: *NewCoreCondition("050", "Creative Arts"),
			expected: constants.Evaluation{
				Name:    "Core Creative Arts",
				Status:  constants.StatusUnknown,
				Summary: "Cannot verify core requirement \"Creative Arts\" automatically",
			},
		},
		"Core number fallback": {
			condition: *NewCoreCondition("040", ""),
			expected: constants.Evaluation{
				Name:    "Core 040",
				Status:  constants.StatusUnknown,
				Summary: "Cannot verify core requirement \"040\" automatically",
			},
		},
		"Core semester hours": {
			condition: *NewCoreConditionWithSemesterHours("040", "Language", 3),
			expected: constants.Evaluation{
				Name:    "Core Language",
				Status:  constants.StatusUnknown,
				Summary: "Cannot verify 3 SCH of core requirement \"Language\" automatically",
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := tc.condition.Fulfils(constants.UserInfo{}, false)
			assertEval(t, tc.expected, got)
		})
	}
}

func TestCoreCondition_JSON(t *testing.T) {
	assertJSONRoundTrip[CoreCondition](t, NewCoreCondition("050", "Creative Arts"), "core")
}

func TestCreditHoursConditionFulfils(t *testing.T) {
	testCases := map[string]struct {
		condition CreditHoursCondition
		userInfo  constants.UserInfo
		expected  constants.Evaluation
	}{
		"Sufficient hours": {
			condition: *NewCreditHoursCondition(45),
			userInfo: constants.UserInfo{
				TotalSCH: 60,
			},
			expected: constants.Evaluation{
				Name:    "Credit Hours",
				Status:  constants.StatusPass,
				Summary: "Has 60 SCH (requires 45)",
			},
		},
		"Insufficient hours": {
			condition: *NewCreditHoursCondition(45),
			userInfo: constants.UserInfo{
				TotalSCH: 30,
			},
			expected: constants.Evaluation{
				Name:    "Credit Hours",
				Status:  constants.StatusDefiniteFail,
				Summary: "Has 30 SCH but requires 45",
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := tc.condition.Fulfils(tc.userInfo, false)
			assertEval(t, tc.expected, got)
		})
	}
}

func TestCreditHoursCondition_JSON(t *testing.T) {
	assertJSONRoundTrip[CreditHoursCondition](t, NewCreditHoursCondition(45), "credit_hours")
}

func TestCreditHoursFromConditionFulfils(t *testing.T) {
	testCases := map[string]struct {
		condition CreditHoursFromCondition
		userInfo  constants.UserInfo
		expected  constants.Evaluation
	}{
		"Earned required hours": {
			condition: *NewCreditHoursFromCondition(6, []constants.Course{
				{Prefix: "BIOL", Number: "2311"},
				{Prefix: "CHEM", Number: "1311"},
			}),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{
					{Prefix: "BIOL", Number: "2311"}: "A",
					{Prefix: "CHEM", Number: "1311"}: "B",
				},
			},
			expected: constants.Evaluation{
				Name:    "Credit Hours From",
				Status:  constants.StatusPass,
				Summary: "Earned 6 SCH from specified courses (requires 6)",
			},
		},
		"Projected hours may satisfy requirement": {
			condition: *NewCreditHoursFromCondition(6, []constants.Course{
				{Prefix: "BIOL", Number: "2311"},
				{Prefix: "CHEM", Number: "1311"},
			}),
			userInfo: constants.UserInfo{
				CurrentEnrollment: []constants.Course{
					{Prefix: "BIOL", Number: "2311"},
					{Prefix: "CHEM", Number: "1311"},
				},
			},
			expected: constants.Evaluation{
				Name:    "Credit Hours From",
				Status:  constants.StatusPossibleFail,
				Summary: "Earned 0 SCH; 6 more in progress — may reach required 6 SCH",
			},
		},
		"Insufficient specified hours": {
			condition: *NewCreditHoursFromCondition(6, []constants.Course{
				{Prefix: "BIOL", Number: "2311"},
				{Prefix: "CHEM", Number: "1311"},
			}),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{
					{Prefix: "BIOL", Number: "2311"}: "A",
				},
			},
			expected: constants.Evaluation{
				Name:    "Credit Hours From",
				Status:  constants.StatusDefiniteFail,
				Summary: "Only 3 SCH earned from specified courses (requires 6)",
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := tc.condition.Fulfils(tc.userInfo, false)
			assertEval(t, tc.expected, got)
		})
	}
}

func TestCreditHoursFromCondition_JSON(t *testing.T) {
	assertJSONRoundTrip[CreditHoursFromCondition](t, NewCreditHoursFromCondition(6, []constants.Course{{Prefix: "BIOL", Number: "2311"}}), "credit_hours_from")
}

func TestUpperDivisionCoursesConditionFulfils(t *testing.T) {
	testCases := map[string]struct {
		condition UpperDivisionCoursesCondition
		userInfo  constants.UserInfo
		expected  constants.Evaluation
	}{
		"Enough upper-division hours with prefix": {
			condition: *NewUpperDivisionCreditHoursCondition(6, "CS"),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{
					{Prefix: "CS", Number: "3345"}:   "A",
					{Prefix: "CS", Number: "4337"}:   "B",
					{Prefix: "MATH", Number: "3310"}: "A",
				},
			},
			expected: constants.Evaluation{
				Name:    "Upper Division Courses",
				Status:  constants.StatusPass,
				Summary: "Has 6 upper-division SCH in CS (requires 6)",
			},
		},
		"Not enough upper-division hours": {
			condition: *NewUpperDivisionCreditHoursCondition(6, "CS"),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{
					{Prefix: "CS", Number: "3345"}: "A",
					{Prefix: "CS", Number: "2336"}: "B",
				},
			},
			expected: constants.Evaluation{
				Name:    "Upper Division Courses",
				Status:  constants.StatusDefiniteFail,
				Summary: "Has 3 upper-division SCH in CS (requires 6)",
			},
		},
		"Enough upper-division count with any prefix": {
			condition: *NewUpperDivisionCountCondition(2, ""),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{
					{Prefix: "CS", Number: "3345"}:   "A",
					{Prefix: "MATH", Number: "3310"}: "A",
				},
			},
			expected: constants.Evaluation{
				Name:    "Upper Division Courses",
				Status:  constants.StatusPass,
				Summary: "Has completed 2 upper-division courses in any prefix (requires 2)",
			},
		},
		"Not enough upper-division count": {
			condition: *NewUpperDivisionCountCondition(2, "CS"),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{
					{Prefix: "CS", Number: "3345"}: "A",
				},
			},
			expected: constants.Evaluation{
				Name:    "Upper Division Courses",
				Status:  constants.StatusDefiniteFail,
				Summary: "Has completed 1 upper-division courses in CS (requires 2)",
			},
		},
		"Invalid upper-division condition": {
			condition: UpperDivisionCoursesCondition{},
			userInfo:  constants.UserInfo{},
			expected: constants.Evaluation{
				Name:    "Upper Division Courses",
				Status:  constants.StatusInvalidRule,
				Summary: "UpperDivisionCoursesCondition has neither Hours nor Count set",
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := tc.condition.Fulfils(tc.userInfo, false)
			assertEval(t, tc.expected, got)
		})
	}
}

func TestUpperDivisionCoursesCondition_JSON(t *testing.T) {
	assertJSONRoundTrip[UpperDivisionCoursesCondition](t, NewUpperDivisionCreditHoursCondition(6, "CS"), "upper_division_courses")
}

func TestResearchConditionFulfils(t *testing.T) {
	testCases := map[string]struct {
		condition ResearchCondition
		userInfo  constants.UserInfo
		expected  constants.Evaluation
	}{
		"Wrong degree level": {
			condition: *NewResearchCondition(3, constants.Graduate),
			userInfo: constants.UserInfo{
				DegreeLevel: constants.Undergraduate,
			},
			expected: constants.Evaluation{
				Name:    "Research",
				Status:  constants.StatusDefiniteFail,
				Summary: "Requires Graduate standing; student is Undergraduate",
			},
		},
		"Matching degree level cannot be verified": {
			condition: *NewResearchCondition(3, constants.Graduate),
			userInfo: constants.UserInfo{
				DegreeLevel: constants.Graduate,
			},
			expected: constants.Evaluation{
				Name:    "Research",
				Status:  constants.StatusUnknown,
				Summary: "Cannot automatically verify 3 research hours for Graduate students",
			},
		},
		"Any degree level cannot be verified": {
			condition: *NewResearchCondition(3, constants.AnyDegree),
			userInfo:  constants.UserInfo{},
			expected: constants.Evaluation{
				Name:    "Research",
				Status:  constants.StatusUnknown,
				Summary: "Cannot automatically verify 3 research hours for Any Degree students",
			},
		},
		"No degree level cannot be verified": {
			condition: ResearchCondition{Hours: 3},
			userInfo:  constants.UserInfo{},
			expected: constants.Evaluation{
				Name:    "Research",
				Status:  constants.StatusUnknown,
				Summary: "Cannot automatically verify 3 research hours for  students",
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := tc.condition.Fulfils(tc.userInfo, false)
			assertEval(t, tc.expected, got)
		})
	}
}

func TestResearchCondition_JSON(t *testing.T) {
	assertJSONRoundTrip[ResearchCondition](t, NewResearchCondition(3, constants.Undergraduate), "research")
}

func TestNCoursesConditionFulfils(t *testing.T) {
	testCases := map[string]struct {
		condition NCoursesCondition
		userInfo  constants.UserInfo
		expected  constants.Evaluation
	}{
		"Completed enough courses": {
			condition: *NewNCoursesCondition(2, []constants.Course{
				{Prefix: "BIOL", Number: "2311"},
				{Prefix: "CHEM", Number: "1311"},
				{Prefix: "PHYS", Number: "2325"},
			}),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{
					{Prefix: "BIOL", Number: "2311"}: "A",
					{Prefix: "CHEM", Number: "1311"}: "B",
				},
			},
			expected: constants.Evaluation{
				Name:    "N Courses",
				Status:  constants.StatusPass,
				Summary: "Completed 2 of required 2 courses",
				Children: []constants.Evaluation{
					{Name: "BIOL 2311", Status: constants.StatusPass, Summary: "Completed BIOL 2311"},
					{Name: "CHEM 1311", Status: constants.StatusPass, Summary: "Completed CHEM 1311"},
					{Name: "PHYS 2325", Status: constants.StatusDefiniteFail, Summary: "Have not taken PHYS 2325"},
				},
			},
		},
		"In progress courses may satisfy requirement": {
			condition: *NewNCoursesCondition(2, []constants.Course{
				{Prefix: "BIOL", Number: "2311"},
				{Prefix: "CHEM", Number: "1311"},
			}),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{
					{Prefix: "BIOL", Number: "2311"}: "A",
				},
				CurrentEnrollment: []constants.Course{{Prefix: "CHEM", Number: "1311"}},
			},
			expected: constants.Evaluation{
				Name:    "N Courses",
				Status:  constants.StatusPossibleFail,
				Summary: "Completed 1 courses; 1 in progress — may satisfy requirement of 2",
				Children: []constants.Evaluation{
					{Name: "BIOL 2311", Status: constants.StatusPass, Summary: "Completed BIOL 2311"},
					{Name: "CHEM 1311", Status: constants.StatusPossibleFail, Summary: "Currently enrolled in CHEM 1311"},
				},
			},
		},
		"Not enough courses": {
			condition: *NewNCoursesCondition(2, []constants.Course{
				{Prefix: "BIOL", Number: "2311"},
				{Prefix: "CHEM", Number: "1311"},
			}),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{
					{Prefix: "BIOL", Number: "2311"}: "A",
				},
			},
			expected: constants.Evaluation{
				Name:    "N Courses",
				Status:  constants.StatusDefiniteFail,
				Summary: "Completed only 1 of required 2 courses from the specified list",
				Children: []constants.Evaluation{
					{Name: "BIOL 2311", Status: constants.StatusPass, Summary: "Completed BIOL 2311"},
					{Name: "CHEM 1311", Status: constants.StatusDefiniteFail, Summary: "Have not taken CHEM 1311"},
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := tc.condition.Fulfils(tc.userInfo, false)
			assertEval(t, tc.expected, got)
		})
	}
}

func TestNCoursesCondition_JSON(t *testing.T) {
	assertJSONRoundTrip[NCoursesCondition](t, NewNCoursesCondition(1, []constants.Course{{Prefix: "BIOL", Number: "2311"}}), "n_courses")
}
