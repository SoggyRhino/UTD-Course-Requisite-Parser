package rules

import "parser/utils"

type DegreeSatisfactionRule struct {
	Prefixes    []string
	Degrees     []string
	Schools     []string
	DegreeLevel utils.DegreeLevel
	Elective    bool
	Math        bool
}

func NewDegreeSatisfactionRuleFromPrefix(prefixes []string, level utils.DegreeLevel) *DegreeSatisfactionRule {
	return &DegreeSatisfactionRule{Prefixes: prefixes, DegreeLevel: level}
}

func NewDegreeSatisfactionRuleFromDegree(degrees []string, level utils.DegreeLevel) *DegreeSatisfactionRule {
	return &DegreeSatisfactionRule{Degrees: degrees, DegreeLevel: level}
}

func NewDegreeSatisfactionRuleFromSchool(schools []string, level utils.DegreeLevel) *DegreeSatisfactionRule {
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
