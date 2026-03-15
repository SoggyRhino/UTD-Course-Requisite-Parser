package conditions

type CourseCondition struct {
	Course Course
}

func (c CourseCondition) Fulfils(userInfo UserInfo) (bool, error) {
	return false, nil
}
