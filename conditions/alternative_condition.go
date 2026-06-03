package conditions

import (
	"encoding/json"
	"log"
	"parser/constants"
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

func (a *AlternativeCondition) UnmarshalJSON(b []byte) error {
	var raw struct {
		Condition json.RawMessage `json:"condition"`
	}
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
	return &AlternativeCondition{
		Condition: condition,
	}
}

func (a *AlternativeCondition) Fulfils(info constants.UserInfo) (bool, error) {
	log.Println("AlternativeCondition Fulfils not implemented")
	return a.Condition.Fulfils(info)
}
