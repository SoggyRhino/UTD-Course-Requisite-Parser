package main

import "parser/conditions"

func coursesToCondition(courses []conditions.Course, minGrade conditions.Grade) conditions.Condition {
	if len(courses) == 1 {
		return conditions.CourseCondition{Course: courses[0]}
	}
	or := make([]conditions.Condition, len(courses))
	for i, _ := range courses {
		or[i] = conditions.CourseCondition{}
	}
	return conditions.OrCondition{Conditions: or}
}
