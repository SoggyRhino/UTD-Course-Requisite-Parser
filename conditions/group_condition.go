package conditions

import "parser/utils"

type StudentGroupCondition struct {
	Groups utils.StudentGroup
}

func NewStudentGroupCondition(groups utils.StudentGroup) *StudentGroupCondition {
	return &StudentGroupCondition{
		Groups: groups,
	}
}

func (c *StudentGroupCondition) Fulfils(userInfo utils.UserInfo) (bool, error) {
	return false, nil
}
