package conditions

import (
	"log"
	"parser/utils"
)

type AlternativeCondition struct {
	Condition Condition
}

func NewAlternativeCondition(condition Condition) *AlternativeCondition {
	return &AlternativeCondition{
		Condition: condition,
	}
}

func (a *AlternativeCondition) Fulfils(info utils.UserInfo) (bool, error) {
	log.Println("AlternativeCondition Fulfils not implemented")
	return a.Condition.Fulfils(info)
}
