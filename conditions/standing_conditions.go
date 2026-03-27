package conditions

type GradeLevel int

const (
	Senior GradeLevel = iota
	Junior
	Sophomore
	Freshman
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
