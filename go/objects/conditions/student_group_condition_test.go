package conditions

import (
	"parser/objects/constants"
	"testing"
)

func TestStudentGroupConditionFulfils(t *testing.T) {
	testCases := map[string]struct {
		condition StudentGroupCondition
		userInfo  constants.UserInfo
		expected  constants.Evaluation
	}{
		"Group match": {
			condition: *NewStudentGroupCondition(constants.ComputerScholarsProgram),
			userInfo:  constants.UserInfo{Groups: []constants.StudentGroup{constants.ComputerScholarsProgram}},
			expected: constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: `Student is a member of group "Computer Scholars Program"`,
			},
		},
		"Group mismatch": {
			condition: *NewStudentGroupCondition(constants.ComputerScholarsProgram),
			userInfo:  constants.UserInfo{Groups: []constants.StudentGroup{constants.CollegeVHonors}},
			expected: constants.Evaluation{
				Status:  constants.StatusDefiniteFail,
				Summary: `Student is not a member of group "Computer Scholars Program"`,
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

func TestStudentGroupCondition_JSON(t *testing.T) {
	cond := NewStudentGroupCondition("Honors")
	assertJSONRoundTrip[StudentGroupCondition](t, cond, "student_group")
}
