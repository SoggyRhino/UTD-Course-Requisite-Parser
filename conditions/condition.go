package conditions

import "parser/utils"

type Condition interface {
	Fulfils(utils.UserInfo) (bool, error)
}

type GradedCondition interface {
	Condition
	AppendGrade(utils.Grade)
}
