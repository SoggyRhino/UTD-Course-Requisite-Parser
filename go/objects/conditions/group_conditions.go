package conditions

import (
	"encoding/json"
	"fmt"
	"parser/objects/constants"
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

type rawOrCondition struct {
	Conditions []any `json:"conditions"`
}

func (o *OrCondition) UnmarshalJSON(b []byte) error {
	var raw rawOrCondition
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	o.Conditions = make([]Condition, len(raw.Conditions))
	for i, rawCondAny := range raw.Conditions {
		rawCond, _ := json.Marshal(rawCondAny)
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
	var flat []Condition
	if or1, ok := cond1.(*OrCondition); ok {
		flat = append(flat, or1.Conditions...)
	} else {
		flat = append(flat, cond1)
	}
	if or2, ok := cond2.(*OrCondition); ok {
		flat = append(flat, or2.Conditions...)
	} else {
		flat = append(flat, cond2)
	}
	return &OrCondition{Conditions: flat}
}

func (o *OrCondition) AppendGrade(grade constants.Grade) {
	for _, c := range o.Conditions {
		if gc, ok := c.(GradedCondition); ok {
			gc.AppendGrade(grade)
		}
	}
}

func (o *OrCondition) Fulfils(info constants.UserInfo, allowCoReq bool) *constants.Evaluation {

	children := make([]constants.Evaluation, 0, len(o.Conditions))
	var bestNonPass constants.EvalStatus
	initialized := false

	for _, c := range o.Conditions {
		evaluation := c.Fulfils(info, allowCoReq)
		if evaluation == nil {
			return &constants.Evaluation{Name: "Or", Status: constants.StatusSystemError, Summary: "condition returned nil"}
		}
		children = append(children, *evaluation)

		if evaluation.Status == constants.StatusPass {
			return &constants.Evaluation{
				Name:     "Or",
				Status:   constants.StatusPass,
				Summary:  fmt.Sprintf("At least one of %d conditions satisfied", len(o.Conditions)),
				Children: children,
			}
		}

		if !initialized || evaluation.Status.Priority() > bestNonPass.Priority() {
			bestNonPass = evaluation.Status
			initialized = true
		}
	}

	return &constants.Evaluation{
		Name:     "Or",
		Status:   bestNonPass,
		Summary:  fmt.Sprintf("None of %d conditions satisfied", len(o.Conditions)),
		Children: children,
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

type rawAndCondition struct {
	Conditions []any `json:"conditions"`
}

func (a *AndCondition) UnmarshalJSON(b []byte) error {
	var raw rawAndCondition
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	a.Conditions = make([]Condition, len(raw.Conditions))
	for i, rawCondAny := range raw.Conditions {
		rawCond, _ := json.Marshal(rawCondAny)
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
	var flat []Condition
	if and1, ok := cond1.(*AndCondition); ok {
		flat = append(flat, and1.Conditions...)
	} else {
		flat = append(flat, cond1)
	}
	if and2, ok := cond2.(*AndCondition); ok {
		flat = append(flat, and2.Conditions...)
	} else {
		flat = append(flat, cond2)
	}
	return &AndCondition{Conditions: flat}
}

func (a *AndCondition) AppendGrade(grade constants.Grade) {
	for _, c := range a.Conditions {
		if gc, ok := c.(GradedCondition); ok {
			gc.AppendGrade(grade)
		}
	}
}

func (a *AndCondition) Fulfils(info constants.UserInfo, allowCoReq bool) *constants.Evaluation {
	children := make([]constants.Evaluation, 0, len(a.Conditions))

	worstStatus := constants.StatusPass
	for _, c := range a.Conditions {
		evaluation := c.Fulfils(info, allowCoReq)
		if evaluation == nil {
			evaluation = &constants.Evaluation{Name: "And", Status: constants.StatusSystemError, Summary: "condition returned nil"}
		}
		children = append(children, *evaluation)

		if evaluation.Status.Priority() < worstStatus.Priority() {
			worstStatus = evaluation.Status
		}
	}

	if worstStatus == constants.StatusPass {
		return &constants.Evaluation{
			Name:     "And",
			Status:   constants.StatusPass,
			Summary:  fmt.Sprintf("All %d conditions satisfied", len(a.Conditions)),
			Children: children,
		}
	}
	return &constants.Evaluation{
		Name:     "And",
		Status:   worstStatus,
		Summary:  fmt.Sprintf("Not all conditions satisfied (%d total)", len(a.Conditions)),
		Children: children,
	}
}
