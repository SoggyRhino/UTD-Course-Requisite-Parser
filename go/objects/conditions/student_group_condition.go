package conditions

import (
	"encoding/json"
	"fmt"
	"parser/objects/constants"
)

type StudentGroupCondition struct {
	Groups constants.StudentGroup `json:"groups,omitempty"`
}

func (c *StudentGroupCondition) MarshalJSON() ([]byte, error) {
	type Alias StudentGroupCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "student_group",
		Alias: (*Alias)(c),
	})
}

func NewStudentGroupCondition(groups constants.StudentGroup) *StudentGroupCondition {
	return &StudentGroupCondition{
		Groups: groups,
	}
}

func (c *StudentGroupCondition) Fulfils(userInfo constants.UserInfo) *constants.Evaluation {
	for _, group := range userInfo.Groups {
		if group == c.Groups {
			return &constants.Evaluation{
				Status:  constants.StatusPass,
				Summary: fmt.Sprintf("Student is a member of group %q", c.Groups),
			}
		}
	}
	return &constants.Evaluation{
		Status:  constants.StatusDefiniteFail,
		Summary: fmt.Sprintf("Student is not a member of group %q", c.Groups),
	}
}
