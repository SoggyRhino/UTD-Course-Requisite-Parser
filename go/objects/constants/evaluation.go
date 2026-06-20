package constants

type Evaluation struct {
	Name     string       `json:"name,omitempty"`
	Status   EvalStatus   `json:"status"`
	Summary  string       `json:"summary,omitempty"`
	Children []Evaluation `json:"children,omitempty"`
}

type EvalStatus string

const (
	StatusPass         EvalStatus = "pass"
	StatusDefiniteFail EvalStatus = "definite_fail"
	StatusPossibleFail EvalStatus = "possible_fail"
	StatusUnknown      EvalStatus = "unknown"
	StatusInvalidRule  EvalStatus = "invalid_rule"
	StatusSystemError  EvalStatus = "system_error"
)

func (s EvalStatus) Priority() int {
	switch s {
	case StatusPass:
		return 5
	case StatusDefiniteFail:
		return 4
	case StatusPossibleFail:
		return 3
	case StatusUnknown:
		return 2
	case StatusInvalidRule:
		return 1
	case StatusSystemError:
		return 0
	}
	return -1
}

func WorstStatus(a, b EvalStatus) EvalStatus {
	if b.Priority() < a.Priority() {
		return b
	}
	return a
}
