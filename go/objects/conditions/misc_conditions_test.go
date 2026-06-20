package conditions

import (
	"parser/objects/constants"
	"testing"
)

func TestConcurrentEnrollmentConditionFulfils(t *testing.T) {
	testCases := map[string]struct {
		condition ConcurrentEnrollmentCondition
		userInfo  constants.UserInfo
		expected  constants.Evaluation
	}{
		"Already taken": {
			condition: *NewConcurrentEnrollmentCondition(constants.Course{Prefix: "CS", Number: "1337"}),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{{Prefix: "CS", Number: "1337"}: "A"},
			},
			expected: constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: "Already completed CS 1337",
			},
		},
		"Currently enrolled": {
			condition: *NewConcurrentEnrollmentCondition(constants.Course{Prefix: "CS", Number: "1337"}),
			userInfo: constants.UserInfo{
				CurrentEnrollment: []constants.Course{{Prefix: "CS", Number: "1337"}},
			},
			expected: constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: "Currently enrolled in CS 1337",
			},
		},
		"Missing": {
			condition: *NewConcurrentEnrollmentCondition(constants.Course{Prefix: "CS", Number: "1337"}),
			userInfo:  constants.UserInfo{},
			expected: constants.Evaluation{
				Status:  constants.StatusPossibleFail,
				Summary: "Requires concurrent enrollment in CS 1337",
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

func TestExactSectionConditionFulfils(t *testing.T) {
	testCases := map[string]struct {
		condition ExactSectionCondition
		userInfo  constants.UserInfo
		expected  constants.Evaluation
	}{
		"Section match": {
			condition: *NewExactSectionCondition(constants.Course{Prefix: "CS", Number: "1337", Section: "001"}),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{{Prefix: "CS", Number: "1337", Section: "001"}: "A"},
			},
			expected: constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: "Completed CS 1337 section 001",
			},
		},
		"Section mismatch": {
			condition: *NewExactSectionCondition(constants.Course{Prefix: "CS", Number: "1337", Section: "001"}),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{{Prefix: "CS", Number: "1337", Section: "501"}: "A"},
			},
			expected: constants.Evaluation{
				Status:  constants.StatusDefiniteFail,
				Summary: "Must be in CS 1337 section 001",
			},
		},
		"Section enrollment match": {
			condition: *NewExactSectionCondition(constants.Course{Prefix: "CS", Number: "1337", Section: "001"}),
			userInfo: constants.UserInfo{
				CurrentEnrollment: []constants.Course{{Prefix: "CS", Number: "1337", Section: "001"}},
			},
			expected: constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: "Currently enrolled in CS 1337 section 001",
			},
		},
		"Empty section match": {
			condition: *NewExactSectionCondition(constants.Course{Prefix: "CS", Number: "1337"}),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{{Prefix: "CS", Number: "1337", Section: "501"}: "A"},
			},
			expected: constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: "Completed CS 1337 section ",
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

func TestAnyPreviousMajorCourseConditionFulfils(t *testing.T) {
	testCases := map[string]struct {
		condition AnyPreviousMajorCourseCondition
		userInfo  constants.UserInfo
		expected  constants.Evaluation
	}{
		"Found": {
			condition: *NewAnyPreviousMajorCourseCondition("CS"),
			userInfo: constants.UserInfo{
				Taken: map[constants.Course]constants.Grade{{Prefix: "CS", Number: "1337"}: "A"},
			},
			expected: constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: "Completed at least one course with prefix CS (CS 1337)",
			},
		},
		"Not found": {
			condition: *NewAnyPreviousMajorCourseCondition("CS"),
			userInfo:  constants.UserInfo{},
			expected: constants.Evaluation{
				Status:  constants.StatusDefiniteFail,
				Summary: "No previous courses with prefix CS found",
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

func TestAcademicYearConditionFulfils(t *testing.T) {
	testCases := map[string]struct {
		condition AcademicYearCondition
		userInfo  constants.UserInfo
		expected  constants.Evaluation
	}{
		"Plan match (equal=true)": {
			condition: *NewAcademicYearCondition("2023", true),
			userInfo:  constants.UserInfo{AcademicPlan: "2023"},
			expected: constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: "Academic plan 2023 matches required 2023",
			},
		},
		"Plan mismatch (equal=true)": {
			condition: *NewAcademicYearCondition("2023", true),
			userInfo:  constants.UserInfo{AcademicPlan: "2022"},
			expected: constants.Evaluation{
				Status:  constants.StatusDefiniteFail,
				Summary: "Academic plan 2022 does not match required 2023",
			},
		},
		"Plan match (equal=false)": {
			condition: *NewAcademicYearCondition("2023", false),
			userInfo:  constants.UserInfo{AcademicPlan: "2023"},
			expected: constants.Evaluation{
				Status:  constants.StatusDefiniteFail,
				Summary: "Academic plan 2023 matches prohibited 2023",
			},
		},
		"Plan mismatch (equal=false)": {
			condition: *NewAcademicYearCondition("2023", false),
			userInfo:  constants.UserInfo{AcademicPlan: "2022"},
			expected: constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: "Academic plan 2022 is not 2023",
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

func TestConcurrentEnrollmentCondition_JSON(t *testing.T) {
	cond := NewConcurrentEnrollmentCondition(constants.Course{Prefix: "CS", Number: "1337"})
	assertJSONRoundTrip[ConcurrentEnrollmentCondition](t, cond, "concurrent_enrollment")
}

func TestExactSectionCondition_JSON(t *testing.T) {
	cond := NewExactSectionCondition(constants.Course{Prefix: "CS", Number: "1337"})
	assertJSONRoundTrip[ExactSectionCondition](t, cond, "exact_section")
}

func TestAnyPreviousMajorCourseCondition_JSON(t *testing.T) {
	cond := NewAnyPreviousMajorCourseCondition("CS")
	assertJSONRoundTrip[AnyPreviousMajorCourseCondition](t, cond, "any_previous_major_course")
}

func TestAcademicYearCondition_JSON(t *testing.T) {
	cond := NewAcademicYearCondition("2023", true)
	assertJSONRoundTrip[AcademicYearCondition](t, cond, "academic_year")
}
