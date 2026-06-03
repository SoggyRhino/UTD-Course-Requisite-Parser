package conditions

import (
	"encoding/json"
	"parser/constants"
)

type OrCondition struct {
	Conditions []Condition `json:"conditions,omitempty"`
}

func (o *OrCondition) MarshalJSON() ([]byte, error) {
	type Alias OrCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "or",
		Alias: (*Alias)(o),
	})
}

func (o *OrCondition) UnmarshalJSON(b []byte) error {
	var raw struct {
		Conditions []json.RawMessage `json:"conditions"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	o.Conditions = make([]Condition, len(raw.Conditions))
	for i, rawCond := range raw.Conditions {
		cond, err := UnmarshalCondition(rawCond)
		if err != nil {
			return err
		}
		o.Conditions[i] = cond
	}
	return nil
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

func (o *OrCondition) Fulfils(userInfo constants.UserInfo) (bool, error) {
	return false, nil
}

func (o *OrCondition) AppendGrade(grade constants.Grade) {
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
	Conditions []Condition `json:"conditions,omitempty"`
}

func (a *AndCondition) MarshalJSON() ([]byte, error) {
	type Alias AndCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "and",
		Alias: (*Alias)(a),
	})
}

func (a *AndCondition) UnmarshalJSON(b []byte) error {
	var raw struct {
		Conditions []json.RawMessage `json:"conditions"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	a.Conditions = make([]Condition, len(raw.Conditions))
	for i, rawCond := range raw.Conditions {
		cond, err := UnmarshalCondition(rawCond)
		if err != nil {
			return err
		}
		a.Conditions[i] = cond
	}
	return nil
}

func NewAndCondition(conditions ...Condition) Condition {
	if len(conditions) == 1 {
		return conditions[0]
	}
	return &AndCondition{Conditions: conditions}
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

func (a *AndCondition) Fulfils(userInfo constants.UserInfo) (bool, error) {
	//todo and condition
	return false, nil
}

func (a *AndCondition) AppendGrade(grade constants.Grade) {
	for _, condition := range a.Conditions {
		if gradedCondition, ok := condition.(GradedCondition); ok {
			gradedCondition.AppendGrade(grade)
		} else {
			//todo look into this, temporary solution for testing.
			panic("condition does not implement GradedCondition, probably a weird case")
		}
	}
}
