package conditions

import "fmt"

type OrCondition struct {
	Conditions []Condition
}

func (o OrCondition) Fulfils(userInfo UserInfo) (bool, error) {
	for _, condition := range o.Conditions {
		if pass, _ := condition.Fulfils(userInfo); pass {
			return true, nil
		}
	}
	return false, fmt.Errorf("no conditions met")
}
