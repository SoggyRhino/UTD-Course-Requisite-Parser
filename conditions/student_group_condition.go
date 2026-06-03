package conditions

import (
	"encoding/json"
	"parser/constants"
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

func (c *StudentGroupCondition) Fulfils(userInfo constants.UserInfo) (bool, error) {
	return false, nil
}
