package rules

import (
	"encoding/json"
	"fmt"
)

type Rule interface {
	isRule() bool //todo change
}

type ruleEnvelope struct {
	Type string `json:"type"`
}

func UnmarshalRule(b []byte) (Rule, error) {
	var env ruleEnvelope
	if err := json.Unmarshal(b, &env); err != nil {
		return nil, err
	}
	switch env.Type {
	case "repeat":
		var r RepeatRule
		return &r, json.Unmarshal(b, &r)
	case "gpa_repeat":
		var r GpaRepeatRule
		return &r, json.Unmarshal(b, &r)
	case "credit_for":
		var r CreditForRule
		return &r, json.Unmarshal(b, &r)
	case "degree_satisfaction":
		var r DegreeSatisfactionRule
		return &r, json.Unmarshal(b, &r)
	case "living_learning":
		var r LivingLearningRule
		return &r, json.Unmarshal(b, &r)
	case "school":
		var r SchoolRule
		return &r, json.Unmarshal(b, &r)
	case "same_as":
		var r SameAsRule
		return &r, json.Unmarshal(b, &r)
	default:
		return nil, fmt.Errorf("unknown rule type: %s", env.Type)
	}
}
