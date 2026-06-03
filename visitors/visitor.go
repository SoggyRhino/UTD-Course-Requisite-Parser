package visitors

import (
	"encoding/json"
	"parser/conditions"
	"parser/constants"
	"parser/parser"
	"parser/rules"

	"github.com/antlr4-go/antlr/v4"
)

var _ parser.RequirementsVisitor = (*RequisiteVisitor)(nil)

type Requirements struct {
	PreReqs     conditions.Condition `json:"pre_reqs,omitempty"`
	CoReqs      conditions.Condition `json:"co_reqs,omitempty"`
	PreOrCoReqs conditions.Condition `json:"pre_or_co_reqs,omitempty"`
	Rules       []rules.Rule         `json:"rules,omitempty"`
	Notices     []constants.Notice   `json:"notices,omitempty"`
}

func (r *Requirements) UnmarshalJSON(b []byte) error {
	type Alias Requirements
	raw := struct {
		PreReqs     json.RawMessage   `json:"pre_reqs,omitempty"`
		CoReqs      json.RawMessage   `json:"co_reqs,omitempty"`
		PreOrCoReqs json.RawMessage   `json:"pre_or_co_reqs,omitempty"`
		Rules       []json.RawMessage `json:"rules,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	if len(raw.PreReqs) > 0 {
		cond, err := conditions.UnmarshalCondition(raw.PreReqs)
		if err != nil {
			return err
		}
		r.PreReqs = cond
	}
	if len(raw.CoReqs) > 0 {
		cond, err := conditions.UnmarshalCondition(raw.CoReqs)
		if err != nil {
			return err
		}
		r.CoReqs = cond
	}
	if len(raw.PreOrCoReqs) > 0 {
		cond, err := conditions.UnmarshalCondition(raw.PreOrCoReqs)
		if err != nil {
			return err
		}
		r.PreOrCoReqs = cond
	}

	if len(raw.Rules) > 0 {
		r.Rules = make([]rules.Rule, len(raw.Rules))
		for i, rawRule := range raw.Rules {
			rule, err := rules.UnmarshalRule(rawRule)
			if err != nil {
				return err
			}
			r.Rules[i] = rule
		}
	}

	return nil
}

type RequisiteVisitor struct {
	parser.BaseRequirementsVisitor
	Tokens       *antlr.CommonTokenStream
	Errors       []error
	Requirements Requirements
}

func NewRequisiteVisitor(tokens *antlr.CommonTokenStream) *RequisiteVisitor {
	return &RequisiteVisitor{
		Tokens: tokens,
		Errors: make([]error, 0),
	}
}

func (v *RequisiteVisitor) appendPreReq(condition conditions.Condition) {
	if v.Requirements.PreReqs == nil {
		v.Requirements.PreReqs = condition
	} else {
		v.Requirements.PreReqs = conditions.NewAndConditionFromExpr(v.Requirements.PreReqs, condition)
	}
}

func (v *RequisiteVisitor) appendCoReq(condition conditions.Condition) {
	if v.Requirements.CoReqs == nil {
		v.Requirements.CoReqs = condition
	} else {
		v.Requirements.CoReqs = conditions.NewAndConditionFromExpr(v.Requirements.CoReqs, condition)
	}
}

func (v *RequisiteVisitor) appendPreOrCoReq(condition conditions.Condition) {
	if v.Requirements.PreOrCoReqs == nil {
		v.Requirements.PreOrCoReqs = condition
	} else {
		v.Requirements.PreOrCoReqs = conditions.NewAndConditionFromExpr(v.Requirements.PreOrCoReqs, condition)
	}
}

func (v *RequisiteVisitor) appendRule(rule rules.Rule) {
	v.Requirements.Rules = append(v.Requirements.Rules, rule)
}

func (v *RequisiteVisitor) appendNotice(notice constants.Notice) {
	v.Requirements.Notices = append(v.Requirements.Notices, notice)
}

func (v *RequisiteVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *RequisiteVisitor) getText(ctx antlr.BaseParserRuleContext) string {
	start := ctx.GetStart().GetStart()
	stop := ctx.GetStop().GetStop()

	return ctx.GetStart().GetInputStream().GetTextFromInterval(antlr.Interval{Start: start, Stop: stop})
}

func (v *RequisiteVisitor) getTextOrDefault(node antlr.TerminalNode, str string) string {
	if node != nil {
		return node.GetText()
	}
	return str
}

func (v *RequisiteVisitor) firstOrNil(nodes []antlr.TerminalNode) antlr.TerminalNode {
	if len(nodes) == 0 {
		return nil
	}
	return nodes[0]
}
