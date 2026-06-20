package rules

import (
	"encoding/json"
	"fmt"
	"parser/objects/constants"
	"strings"
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

func (r *LivingLearningRule) Fulfils(userInfo constants.UserInfo) *constants.Evaluation {
	if len(r.Prefixes) > 0 {
		return &constants.Evaluation{
			Name:    "Living Learning Rule",
			Status:  constants.StatusPossibleFail,
			Summary: "Living learning prefixes mapping not implemented yet",
		}
	}

	for _, degree := range r.Degrees {
		if userInfo.Major == degree {
			return &constants.Evaluation{
				Name:    "Living Learning Rule",
				Status:  constants.StatusPass,
				Summary: fmt.Sprintf("Student's major matches living learning degree: %s", degree),
			}
		}
	}

	return &constants.Evaluation{
		Name:    "Living Learning Rule",
		Status:  constants.StatusDefiniteFail,
		Summary: fmt.Sprintf("Student's major %q is not in the required living learning degrees: %s", userInfo.Major, strings.Join(r.Degrees, ", ")),
	}
}
