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

func NewOrConditionFromExpr(cond1, cond2 Condition) *OrCondition {
	var flattenedConditions []Condition

	if or1, isOr1 := cond1.(*OrCondition); isOr1 {
		flattenedConditions = append(flattenedConditions, or1.Conditions...)
	} else {
		flattenedConditions = append(flattenedConditions, cond1)
	}

	if or2, isOr2 := cond2.(*OrCondition); isOr2 {
		flattenedConditions = append(flattenedConditions, or2.Conditions...)
	} else {
		flattenedConditions = append(flattenedConditions, cond2)
	}

	return &OrCondition{Conditions: flattenedConditions}
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

func NewAndConditionFromExpr(cond1, cond2 Condition) *AndCondition {
	var flattenedConditions []Condition

	if and1, isAnd1 := cond1.(*AndCondition); isAnd1 {
		flattenedConditions = append(flattenedConditions, and1.Conditions...)
	} else {
		flattenedConditions = append(flattenedConditions, cond1)
	}

	if and2, isAnd2 := cond2.(*AndCondition); isAnd2 {
		flattenedConditions = append(flattenedConditions, and2.Conditions...)
	} else {
		flattenedConditions = append(flattenedConditions, cond2)
	}

	return &AndCondition{Conditions: flattenedConditions}
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
