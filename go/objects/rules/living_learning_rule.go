package rules

import (
	"encoding/json"
	"parser/objects/constants"
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

func (r *LivingLearningRule) Fulfils(userInfo constants.UserInfo) constants.Evaluation {
	return constants.Evaluation{
		Name:    "Living Learning Rule",
		Status:  constants.StatusDefiniteFail,
		Summary: "Not implemented",
	}
}
