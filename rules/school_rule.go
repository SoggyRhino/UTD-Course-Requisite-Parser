package rules

type SchoolRule struct {
	Schools []string
}

func (r *SchoolRule) isRule() bool {
	//TODO implement me
	panic("implement me")
}

func NewSchoolRule(schools []string) *SchoolRule {
	return &SchoolRule{
		Schools: schools,
	}
}

func (r *SchoolRule) IsRule() bool {
	return true
}
