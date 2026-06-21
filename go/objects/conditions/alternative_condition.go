package conditions

import (
	"encoding/json"
	"parser/objects/constants"
)

type AlternativeCondition struct {
	Condition Condition `json:"condition,omitempty"`
}

func (a *AlternativeCondition) MarshalJSON() ([]byte, error) {
	type Alias AlternativeCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "alternative",
		Alias: (*Alias)(a),
	})
}

type rawAlternativeCondition struct {
	Condition json.RawMessage `json:"condition"`
}

func (a *AlternativeCondition) UnmarshalJSON(b []byte) error {
	var raw rawAlternativeCondition
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	cond, err := UnmarshalCondition(raw.Condition)
	if err != nil {
		return err
	}
	a.Condition = cond
	return nil
}

func NewAlternativeCondition(condition Condition) *AlternativeCondition {
	return &AlternativeCondition{Condition: condition}
}

func (a *AlternativeCondition) Fulfils(info constants.UserInfo, allowCoReq bool) *constants.Evaluation {
	inner := a.Condition.Fulfils(info, allowCoReq)
	if inner == nil {
		//todo remove ability to return nil
		inner = &constants.Evaluation{Name: "Alternative", Status: constants.StatusSystemError, Summary: "inner condition returned nil"}
	}

	if inner.Status == constants.StatusPass {
		return inner
	}

	return &constants.Evaluation{
		Name:     "Alternative",
		Status:   constants.StatusUnknown,
		Summary:  "Standard path not satisfied — an equivalent may also be accepted (contact adviser)",
		Children: []constants.Evaluation{*inner},
	}
}
