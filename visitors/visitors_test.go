package visitors

import (
	"parser/parser"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/google/go-cmp/cmp"
)

type treeCreator func(*parser.RequirementsParser) antlr.ParseTree

func rule[T antlr.ParseTree](fn func(*parser.RequirementsParser) T) treeCreator {
	return func(p *parser.RequirementsParser) antlr.ParseTree {
		return fn(p)
	}
}

func testTree[T any](t *testing.T, input string, creator treeCreator, result T) {
	t.Helper()

	stream := antlr.NewInputStream(input)
	lexer := parser.NewRequirementsLexer(stream)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewRequirementsParser(tokens)
	tree := creator(p)

	visitor := NewRequisiteVisitor(tokens)

	output := visitor.Visit(tree).(T)
	if diff := cmp.Diff(result, output); diff != "" {
		t.Errorf("Unexpected output (-want +got):\n%s", diff)
	}
}


