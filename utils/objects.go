package utils

type Course struct {
	Prefix  string
	Number  string
	Section string
}

type Grade string
type UserInfo struct {
	Taken map[Course]Grade
}
