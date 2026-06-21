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

type rawRequirements struct {
	PreReqs     json.RawMessage    `json:"pre_reqs,omitempty"`
	CoReqs      json.RawMessage    `json:"co_reqs,omitempty"`
	PreOrCoReqs json.RawMessage    `json:"pre_or_co_reqs,omitempty"`
	Rules       []any              `json:"rules,omitempty"`
	Notices     []constants.Notice `json:"notices,omitempty"`
}

func (r *Requirements) UnmarshalJSON(b []byte) error {
	var raw rawRequirements
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
		for i, rawRuleAny := range raw.Rules {
			rawRule, _ := json.Marshal(rawRuleAny)
			rule, err := rules.UnmarshalRule(rawRule)
			if err != nil {
				return err
			}
			r.Rules[i] = rule
		}
	}
	r.Notices = raw.Notices

	return nil
}

func (r *Requirements) Evaluate(info constants.UserInfo) RequirementsResult {
	result := RequirementsResult{
		Overall: constants.StatusPass,
	}

	if r.PreReqs != nil {
		result.PreReqs = r.PreReqs.Fulfils(info, false)
		result.Overall = constants.WorstStatus(result.Overall, result.PreReqs.Status)
	}

	if r.CoReqs != nil {
		result.CoReqs = r.CoReqs.Fulfils(info, true)
		result.Overall = constants.WorstStatus(result.Overall, result.CoReqs.Status)
	}

	if r.PreOrCoReqs != nil {
		result.PreOrCoReqs = r.PreOrCoReqs.Fulfils(info, true)
		result.Overall = constants.WorstStatus(result.Overall, result.PreOrCoReqs.Status)
	}

	if len(r.Rules) > 0 {
		var evaluatedRules []constants.Evaluation
		for _, rule := range r.Rules {
			if res := rule.Fulfils(info); res != nil {
				evaluatedRules = append(evaluatedRules, *res)
				result.Overall = constants.WorstStatus(result.Overall, res.Status)
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
