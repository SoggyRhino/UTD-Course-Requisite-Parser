package conditions

import (
	"encoding/json"
	"parser/objects/constants"
)

type ConsentCondition struct {
	Consent constants.Consent `json:"consent"`
}

func (c *ConsentCondition) MarshalJSON() ([]byte, error) {
	type Alias ConsentCondition
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "consent",
		Alias: (*Alias)(c),
	})
}

func NewConsentCondition(consent constants.Consent) *ConsentCondition {
	return &ConsentCondition{
		Consent: consent,
	}
}

func (c *ConsentCondition) Fulfils(userInfo constants.UserInfo) *constants.Evaluation {
	return &constants.Evaluation{
		Status:  constants.StatusPossibleFail,
		Summary: "Requires manual consent/approval",
	}
}
