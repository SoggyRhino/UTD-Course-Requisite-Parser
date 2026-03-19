package conditions

type Course struct {
	Prefix  string
	Number  string
	Section string
}

type Grade string
type UserInfo struct {
	Taken map[Course]Grade
}

type Condition interface {
	Fulfils(UserInfo) (bool, error)
}

type GradedCondition interface {
	Condition
	AppendGrade(Grade)
}
