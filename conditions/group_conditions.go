package conditions

import (
	"fmt"
	"parser/utils"
)

type OrCondition struct {
	Conditions []Condition
}

func NewOrCondition(conditions ...Condition) Condition {
	if len(conditions) == 1 {
		return conditions[0]
	}
	return &OrCondition{Conditions: conditions}
}

func (o *OrCondition) Fulfils(userInfo utils.UserInfo) (bool, error) {
	//todo or condition
	for _, condition := range o.Conditions {
		if pass, _ := condition.Fulfils(userInfo); pass {
			return true, nil
		}
	}
	return false, fmt.Errorf("no conditions met")
}

func (o *OrCondition) AppendGrade(grade utils.Grade) {
	for _, condition := range o.Conditions {
		if gradedCondition, ok := condition.(GradedCondition); ok {
			gradedCondition.AppendGrade(grade)
		} else {
			//todo look into this, temporary solution for testing.
			panic("condition does not implement GradedCondition, probably a weird case")
		}
	}
}

type AndCondition struct {
	Conditions []Condition
}

func NewAndCondition(conditions ...Condition) Condition {
	if len(conditions) == 1 {
		return conditions[0]
	}
	return &OrCondition{Conditions: conditions}
}

func (a *AndCondition) Fulfils(userInfo utils.UserInfo) (bool, error) {
	//todo and condition
	return false, fmt.Errorf("no conditions met")
}

func (a *AndCondition) AppendGrade(grade utils.Grade) {
	for _, condition := range a.Conditions {
		if gradedCondition, ok := condition.(GradedCondition); ok {
			gradedCondition.AppendGrade(grade)
		} else {
			//todo look into this, temporary solution for testing.
			panic("condition does not implement GradedCondition, probably a weird case")
		}
	}
}
