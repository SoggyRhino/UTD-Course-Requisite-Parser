package conditions

type CreditHoursCondition struct {
	Hours int
}

func NewCreditHoursCondition(hours int) *CreditHoursCondition {
	return &CreditHoursCondition{
		Hours: hours,
	}
}

func (c *CreditHoursCondition) Fulfils(userInfo UserInfo) (bool, error) {
	return false, nil
}

type CreditHoursFromCondition struct {
	Hours   int
	Courses []Course
}

func NewCreditHoursFromCondition(hours int, courses []Course) *CreditHoursFromCondition {
	return &CreditHoursFromCondition{
		Hours:   hours,
		Courses: courses,
	}
}

func (c *CreditHoursFromCondition) Fulfils(userInfo UserInfo) (bool, error) {
	return false, nil
}
