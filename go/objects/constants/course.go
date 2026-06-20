package constants

import (
	"strconv"
	"strings"
)

type Grade string

func (g Grade) GPA() float64 {
	switch g {
	case "A+", "A":
		return 4.0
	case "A-":
		return 3.67
	case "B+":
		return 3.33
	case "B":
		return 3.0
	case "B-":
		return 2.67
	case "C+":
		return 2.33
	case "C":
		return 2.0
	default:
		return 0

	}
}

type Course struct {
	Prefix  string `json:"prefix,omitempty"`
	Number  string `json:"number,omitempty"`
	Section string `json:"section,omitempty"`
}

func (c Course) Hours() int {
	if h, err := strconv.Atoi(c.Number[1:2]); err == nil {
		return h
	}
	return 3
}

func (c Course) CourseNumberInt() int {
	n, err := strconv.Atoi(strings.TrimSpace(c.Number))
	if err != nil {
		return -1
	}
	return n
}

func (c Course) IsUpperDivision() bool {
	return c.CourseNumberInt() >= 3000
}
