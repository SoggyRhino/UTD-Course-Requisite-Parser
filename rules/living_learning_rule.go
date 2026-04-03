package rules

type LivingLearningRule struct {
	Prefixes []string
	Degrees  []string
}

func NewLivingLearningRuleFromPrefixes(prefixes []string) *LivingLearningRule {
	return &LivingLearningRule{Prefixes: prefixes}
}

func NewLivingLearningRuleFromDegrees(degrees []string) *LivingLearningRule {
	return &LivingLearningRule{Degrees: degrees}
}

func (r *LivingLearningRule) isRule() bool {
	return true
}
