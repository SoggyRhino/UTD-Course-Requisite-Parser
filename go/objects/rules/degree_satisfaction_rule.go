package rules

import (
	"encoding/json"
	"parser/objects/constants"
)

type DegreeSatisfactionRule struct {
	Prefixes    []string              `json:"prefixes,omitempty"`
	Degrees     []string              `json:"degrees,omitempty"`
	Schools     []string              `json:"schools,omitempty"`
	DegreeLevel constants.DegreeLevel `json:"degree_level,omitempty"`
	Elective    bool                  `json:"elective,omitempty"`
	Math        bool                  `json:"math,omitempty"`
}

func (r *DegreeSatisfactionRule) MarshalJSON() ([]byte, error) {
	type Alias DegreeSatisfactionRule
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "degree_satisfaction",
		Alias: (*Alias)(r),
	})
}

func NewDegreeSatisfactionRuleFromPrefix(prefixes []string, level constants.DegreeLevel) *DegreeSatisfactionRule {
	return &DegreeSatisfactionRule{Prefixes: prefixes, DegreeLevel: level}
}

func NewDegreeSatisfactionRuleFromDegree(degrees []string, level constants.DegreeLevel) *DegreeSatisfactionRule {
	return &DegreeSatisfactionRule{Degrees: degrees, DegreeLevel: level}
}

func NewDegreeSatisfactionRuleFromSchool(schools []string, level constants.DegreeLevel) *DegreeSatisfactionRule {
	return &DegreeSatisfactionRule{Schools: schools, DegreeLevel: level}
}

func NewMathDegreeSatisfactionRule() *DegreeSatisfactionRule {
	return &DegreeSatisfactionRule{
		Schools: []string{"Mathematics"},
		Math:    true,
	}
}

func NewDegreeSatisfactionRuleFromElectives(rule *DegreeSatisfactionRule) *DegreeSatisfactionRule {
	rule.Elective = true
	return rule
}

func (r *DegreeSatisfactionRule) isRule() bool {
	return true
}
