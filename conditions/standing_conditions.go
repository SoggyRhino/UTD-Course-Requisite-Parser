package conditions

type GradeLevel string

const (
	Senior    GradeLevel = "Senior"
	Junior    GradeLevel = "Junior"
	Sophomore GradeLevel = "Sophomore"
	Freshman  GradeLevel = "Freshman"
)

type GradeLevelCondition struct {
	GradeLevel GradeLevel
	Degree     string
	School     string
}

func NewGradeLevelCondition(level GradeLevel) *GradeLevelCondition {
	return &GradeLevelCondition{
		GradeLevel: level,
	}
}

func NewGradeLevelConditionWithDegree(level GradeLevel, degree string) *GradeLevelCondition {
	return &GradeLevelCondition{
		GradeLevel: level,
		Degree:     degree,
	}
}

func (g *GradeLevelCondition) Fulfils(userInfo UserInfo) (bool, error) {
	return false, nil
}

type GraduateStandingInCondition struct {
	Degree string
}

func NewGraduateStandingInCondition() *GraduateStandingInCondition {
	return &GraduateStandingInCondition{}
}

func NewGraduateStandingInConditionWithDegree(degree string) *GraduateStandingInCondition {
	return &GraduateStandingInCondition{
		Degree: degree}
}

func (g *GraduateStandingInCondition) Fulfils(userInfo UserInfo) (bool, error) {
	return false, nil
}
