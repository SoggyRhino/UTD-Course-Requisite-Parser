package visitors

import (
	"parser/conditions"
	"parser/parser"
	"parser/rules"
	"parser/utils"

	"github.com/antlr4-go/antlr/v4"
)

var _ parser.RequirementsVisitor = (*RequisiteVisitor)(nil)

type Requirements struct {
	PreReqs     conditions.Condition
	CoReqs      conditions.Condition
	PreOrCoReqs conditions.Condition
	Rules       []rules.Rule
	Notices     []utils.Notice
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

func (v *RequisiteVisitor) appendNotice(notice utils.Notice) {
	v.Requirements.Notices = append(v.Requirements.Notices, notice)
}

func (v *RequisiteVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *RequisiteVisitor) ReportError(ctx antlr.BaseParserRuleContext, err error) {
	//todo use ctx for better error handling
	v.Errors = append(v.Errors, err)
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
