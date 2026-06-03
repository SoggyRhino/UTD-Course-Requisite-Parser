package rules

import (
	"encoding/json"
)

type LivingLearningRule struct {
	Prefixes []string `json:"prefixes,omitempty"`
	Degrees  []string `json:"degrees,omitempty"`
}

func (r *LivingLearningRule) MarshalJSON() ([]byte, error) {
	type Alias LivingLearningRule
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "living_learning",
		Alias: (*Alias)(r),
	})
}

func NewLivingLearningRuleFromPrefixes(prefixes []string) *LivingLearningRule {
	return &LivingLearningRule{Prefixes: prefixes}
}

func NewLivingLearningRuleFromDegrees(degrees []string) *LivingLearningRule {
	return &LivingLearningRule{Degrees: degrees}
}

func (r *LivingLearningRule) isRule() bool {
	return true
}
