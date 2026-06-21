package objects

import (
	"encoding/json"
	"fmt"
	"parser/objects/conditions"
	"parser/objects/constants"
	"parser/objects/rules"
)

type RequirementsResult struct {
	Overall     constants.EvalStatus   `json:"overall"`
	PreReqs     *constants.Evaluation  `json:"pre_reqs,omitempty"`
	CoReqs      *constants.Evaluation  `json:"co_reqs,omitempty"`
	PreOrCoReqs *constants.Evaluation  `json:"pre_or_co_reqs,omitempty"`
	Rules       []constants.Evaluation `json:"rules,omitempty"`
	Notices     []constants.Evaluation `json:"notices,omitempty"`
}

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

func (r *Requirements) Evaluate(info constants.UserInfo) RequirementsResult {
	result := RequirementsResult{
		Overall: constants.StatusPass,
	}

	if r.PreReqs != nil {
		eval := r.PreReqs.Fulfils(info, false)
		result.PreReqs = eval
		result.Overall = constants.WorstStatus(result.Overall, eval.Status)
	}

	if r.CoReqs != nil {
		eval := r.CoReqs.Fulfils(info, true)
		result.CoReqs = eval
		result.Overall = constants.WorstStatus(result.Overall, eval.Status)
	}

	if r.PreOrCoReqs != nil {
		eval := r.PreOrCoReqs.Fulfils(info, true)
		result.PreOrCoReqs = eval
		result.Overall = constants.WorstStatus(result.Overall, eval.Status)
	}

	if len(r.Rules) > 0 {
		var evaluatedRules []constants.Evaluation
		for _, rule := range r.Rules {
			eval := rule.Fulfils(info)
			if eval != nil {
				evaluatedRules = append(evaluatedRules, *eval)
				result.Overall = constants.WorstStatus(result.Overall, eval.Status)
			}
		}
		result.Rules = evaluatedRules
	}

	if len(r.Notices) > 0 {
		result.Notices = make([]constants.Evaluation, len(r.Notices))
		for i, notice := range r.Notices {
			result.Notices[i] = constants.Evaluation{
				Name:    string(notice),
				Status:  constants.StatusPass,
				Summary: fmt.Sprintf("Notice: %s", notice),
			}
		}
	}

	return result
}
