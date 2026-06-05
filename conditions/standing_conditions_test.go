package conditions

import (
	"parser/constants"
	"testing"
)

func TestGradeLevelConditionFulfils(t *testing.T) {
	testCases := map[string]struct {
		condition GradeLevelCondition
		userInfo  constants.UserInfo
		expected  constants.Evaluation
	}{
		"Level match": {
			condition: *NewGradeLevelCondition(constants.Senior),
			userInfo:  constants.UserInfo{GradeLevel: constants.Senior},
			expected: constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: "Grade level requirement Senior satisfied",
			},
		},
		"Level mismatch": {
			condition: *NewGradeLevelCondition(constants.Senior),
			userInfo:  constants.UserInfo{GradeLevel: constants.Freshman},
			expected: constants.Evaluation{
				Status:  constants.StatusDefiniteFail,
				Summary: "Grade level is Freshman; requires Senior",
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := tc.condition.Fulfils(tc.userInfo)
			assertEval(t, tc.expected, got)
		})
	}
}

func TestGraduateStandingInConditionFulfils(t *testing.T) {
	testCases := map[string]struct {
		condition GraduateStandingInCondition
		userInfo  constants.UserInfo
		expected  constants.Evaluation
	}{
		"Graduate match": {
			condition: *NewGraduateStandingInCondition(),
			userInfo:  constants.UserInfo{DegreeLevel: constants.Graduate},
			expected: constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: "Graduate standing requirement satisfied",
			},
		},
		"PhD match": {
			condition: *NewGraduateStandingInCondition(),
			userInfo:  constants.UserInfo{DegreeLevel: constants.PhD},
			expected: constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: "Graduate standing requirement satisfied",
			},
		},
		"Undergraduate mismatch": {
			condition: *NewGraduateStandingInCondition(),
			userInfo:  constants.UserInfo{DegreeLevel: constants.Undergraduate},
			expected: constants.Evaluation{
				Status:  constants.StatusDefiniteFail,
				Summary: "Degree level is Undergraduate; requires Graduate standing",
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := tc.condition.Fulfils(tc.userInfo)
			assertEval(t, tc.expected, got)
		})
	}
}

func TestGenericStandingConditionFulfils(t *testing.T) {
	testCases := map[string]struct {
		condition GenericStandingCondition
		userInfo  constants.UserInfo
		expected  constants.Evaluation
	}{
		"Standing match": {
			condition: *NewGenericStandingCondition(constants.GoodAcademicStanding),
			userInfo:  constants.UserInfo{Standing: constants.GoodAcademicStanding},
			expected: constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: `Academic standing "Good Academic Standing" satisfied`,
			},
		},
		"Standing mismatch": {
			condition: *NewGenericStandingCondition(constants.GoodAcademicStanding),
			userInfo:  constants.UserInfo{Standing: constants.Standing("Probation")},
			expected: constants.Evaluation{
				Status:  constants.StatusDefiniteFail,
				Summary: `Academic standing is "Probation"; requires "Good Academic Standing"`,
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := tc.condition.Fulfils(tc.userInfo)
			assertEval(t, tc.expected, got)
		})
	}
}

func TestGradeLevelCondition_JSON(t *testing.T) {
	cond := NewGradeLevelCondition(constants.Senior)
	assertJSONRoundTrip[GradeLevelCondition](t, cond, "grade_level")
}

func TestGraduateStandingInCondition_JSON(t *testing.T) {
	cond := NewGraduateStandingInCondition()
	assertJSONRoundTrip[GraduateStandingInCondition](t, cond, "graduate_standing_in")
}

func TestGenericStandingCondition_JSON(t *testing.T) {
	cond := NewGenericStandingCondition(constants.Standing("Good"))
	assertJSONRoundTrip[GenericStandingCondition](t, cond, "generic_standing")
}

func TestStandingConditionConstructors(t *testing.T) {
	NewGradeLevelConditionWithDegree(constants.Senior, "CS")
	NewGraduateStandingInConditionWithDegree("CS")
}
