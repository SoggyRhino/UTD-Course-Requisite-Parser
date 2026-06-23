package rules

import (
	"encoding/json"
	"fmt"
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

func (r *DegreeSatisfactionRule) Fulfils(userInfo constants.UserInfo) *constants.Evaluation {
	levelMatches := r.DegreeLevel == "" || r.DegreeLevel == userInfo.DegreeLevel

	if levelMatches {
		for _, prefix := range r.Prefixes {
			if userInfo.Major == prefix {
				return &constants.Evaluation{
					Name:    "Degree Satisfaction Rule",
					Status:  constants.StatusDefiniteFail,
					Summary: fmt.Sprintf("Course cannot be used to satisfy degree requirements for %s majors", prefix),
				}
			}
		}

		for _, degree := range r.Degrees {
			if userInfo.Major == degree {
				return &constants.Evaluation{
					Name:    "Degree Satisfaction Rule",
					Status:  constants.StatusDefiniteFail,
					Summary: fmt.Sprintf("Course cannot be used to satisfy degree requirements for the %s degree", degree),
				}
			}
		}

		for _, school := range r.Schools {
			if userInfo.School == school {
				if r.Math {
					return &constants.Evaluation{
						Name:    "Degree Satisfaction Rule",
						Status:  constants.StatusDefiniteFail,
						Summary: "Course cannot be used to satisfy mathematics requirements by students in the Mathematics school",
					}
				}
				return &constants.Evaluation{
					Name:    "Degree Satisfaction Rule",
					Status:  constants.StatusDefiniteFail,
					Summary: fmt.Sprintf("Course cannot be used to satisfy degree requirements for the %s school", school),
				}
			}
		}
	}

	return &constants.Evaluation{
		Name:    "Degree Satisfaction Rule",
		Status:  constants.StatusPass,
		Summary: "Course is eligible for the student's degree requirements",
	}
}
