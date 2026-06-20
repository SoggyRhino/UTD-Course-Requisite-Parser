package conditions

import (
	"parser/objects/constants"
	"testing"
)

func TestMajorConditionFulfils(t *testing.T) {
	testCases := map[string]struct {
		condition MajorCondition
		userInfo  constants.UserInfo
		expected  constants.Evaluation
	}{
		"Major match": {
			condition: *NewMajorCondition("Computer Science"),
			userInfo:  constants.UserInfo{Major: "Computer Science"},
			expected: constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: "Major and level requirements satisfied",
			},
		},
		"Major mismatch": {
			condition: *NewMajorCondition("Computer Science"),
			userInfo:  constants.UserInfo{Major: "Biology"},
			expected: constants.Evaluation{
				Status:  constants.StatusDefiniteFail,
				Summary: `Major is "Biology"; requires "Computer Science"`,
			},
		},
		"Degree level match": {
			condition: *NewMajorConditionWithDegreeLevel("CS", constants.Undergraduate),
			userInfo:  constants.UserInfo{Major: "CS", DegreeLevel: constants.Undergraduate},
			expected: constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: "Major and level requirements satisfied",
			},
		},
		"Degree level mismatch": {
			condition: *NewMajorConditionWithDegreeLevel("CS", constants.Graduate),
			userInfo:  constants.UserInfo{Major: "CS", DegreeLevel: constants.Undergraduate},
			expected: constants.Evaluation{
				Status:  constants.StatusDefiniteFail,
				Summary: `Degree level is "Undergraduate"; requires "Graduate"`,
			},
		},
		"Grade level match": {
			condition: *NewMajorConditionWithGradeLevel("CS", constants.Senior),
			userInfo:  constants.UserInfo{Major: "CS", GradeLevel: constants.Senior},
			expected: constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: "Major and level requirements satisfied",
			},
		},
		"Grade level mismatch": {
			condition: *NewMajorConditionWithGradeLevel("CS", constants.Senior),
			userInfo:  constants.UserInfo{Major: "CS", GradeLevel: constants.Freshman},
			expected: constants.Evaluation{
				Status:  constants.StatusDefiniteFail,
				Summary: `Grade level is "Freshman"; requires "Senior"`,
			},
		},
		"Complex major/level match": {
			condition: *NewMajorConditionWithDegreeAndGradeLevel("Computer Science", constants.Undergraduate, constants.Junior),
			userInfo:  constants.UserInfo{Major: "BS Computer Science", DegreeLevel: constants.Undergraduate, GradeLevel: constants.Junior},
			expected: constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: "Major and level requirements satisfied",
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

func TestDegreeConditionFulfils(t *testing.T) {
	testCases := map[string]struct {
		condition DegreeCondition
		userInfo  constants.UserInfo
		expected  constants.Evaluation
	}{
		"Degree match": {
			condition: *NewDegreeCondition("Computer Science"),
			userInfo:  constants.UserInfo{Major: "BS Computer Science"},
			expected: constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: `Degree requirement "Computer Science" satisfied`,
			},
		},
		"Degree mismatch": {
			condition: *NewDegreeCondition("Computer Science"),
			userInfo:  constants.UserInfo{Major: "Biology"},
			expected: constants.Evaluation{
				Status:  constants.StatusDefiniteFail,
				Summary: `Major is "Biology"; requires degree in "Computer Science"`,
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

func TestMajorCondition_JSON(t *testing.T) {
	cond := NewMajorCondition("Computer Science")
	assertJSONRoundTrip[MajorCondition](t, cond, "major")
}

func TestDegreeCondition_JSON(t *testing.T) {
	cond := NewDegreeCondition("Bachelor of Science")
	assertJSONRoundTrip[DegreeCondition](t, cond, "degree")
}
