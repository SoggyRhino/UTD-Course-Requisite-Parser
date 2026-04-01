package rules

type SchoolRule struct {
	Schools []string
}

func NewSchoolRule(schools []string) *SchoolRule {
	return &SchoolRule{
		Schools: schools,
	}
}
