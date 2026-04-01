package conditions

import "parser/utils"

type ConsentCondition struct {
	Consent utils.Consent
}

func NewConsentCondition(consent utils.Consent) *ConsentCondition {
	return &ConsentCondition{
		Consent: consent,
	}
}

func (c *ConsentCondition) Fulfils(userInfo utils.UserInfo) (bool, error) {
	return false, nil
}
