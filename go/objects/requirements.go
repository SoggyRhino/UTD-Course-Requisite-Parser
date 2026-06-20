package objects

import (
	"encoding/json"
	"parser/objects/conditions"
	"parser/objects/constants"
	"parser/objects/rules"
)

type Requirements struct {
	PreReqs     conditions.Condition `json:"pre_reqs,omitempty"`
	CoReqs      conditions.Condition `json:"co_reqs,omitempty"`
	PreOrCoReqs conditions.Condition `json:"pre_or_co_reqs,omitempty"`
	Rules       []rules.Rule         `json:"rules,omitempty"`
	Notices     []constants.Notice   `json:"notices,omitempty"`
}

func (r *Requirements) UnmarshalJSON(b []byte) error {
	type Alias Requirements
	raw := struct {
		PreReqs     json.RawMessage   `json:"pre_reqs,omitempty"`
		CoReqs      json.RawMessage   `json:"co_reqs,omitempty"`
		PreOrCoReqs json.RawMessage   `json:"pre_or_co_reqs,omitempty"`
		Rules       []json.RawMessage `json:"rules,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	if len(raw.PreReqs) > 0 {
		cond, err := conditions.UnmarshalCondition(raw.PreReqs)
		if err != nil {
			return err
		}
		r.PreReqs = cond
	}
	if len(raw.CoReqs) > 0 {
		cond, err := conditions.UnmarshalCondition(raw.CoReqs)
		if err != nil {
			return err
		}
		r.CoReqs = cond
	}
	if len(raw.PreOrCoReqs) > 0 {
		cond, err := conditions.UnmarshalCondition(raw.PreOrCoReqs)
		if err != nil {
			return err
		}
		r.PreOrCoReqs = cond
	}

	if len(raw.Rules) > 0 {
		r.Rules = make([]rules.Rule, len(raw.Rules))
		for i, rawRule := range raw.Rules {
			rule, err := rules.UnmarshalRule(rawRule)
			if err != nil {
				return err
			}
			r.Rules[i] = rule
		}
	}

	return nil
}
