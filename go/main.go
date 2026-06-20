package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"parser/objects"
	"parser/parser"
	"parser/visitors"
	"slices"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type Course struct {
	SubjectPrefix  string `json:"subject_prefix"`
	CourseNumber   string `json:"course_number"`
	EnrollmentReqs string `json:"enrollment_reqs"`
}

func main() {
	inputPath := "static/courses.json"
	outputPath := "static/requirements.json"

	if err := os.MkdirAll("static", 0755); err != nil {
		log.Fatalf("Failed to create static directory: %v", err)
	}

	file, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	var courses []Course
	if err := json.Unmarshal(file, &courses); err != nil {
		log.Fatalf("Failed to unmarshal courses: %v", err)
	}

	allRequirements := make(map[string]objects.Requirements)
	failed := make([]string, 0, 20)

	for _, course := range courses {
		trimmedReq := strings.TrimSpace(course.EnrollmentReqs)
		if trimmedReq == "" {
			continue
		}

		func() {
			key := fmt.Sprintf("%s %s", course.SubjectPrefix, course.CourseNumber)

			defer func() {
				if r := recover(); r != nil {
					failed = append(failed, key)
				}
			}()

			reqs, hasErr := ParseString(trimmedReq)
			if hasErr {
				failed = append(failed, key)
			} else {
				allRequirements[key] = reqs
			}
		}()
	}

	outputData, err := json.MarshalIndent(allRequirements, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal requirements: %v", err)
	}

	if err := os.WriteFile(outputPath, outputData, 0644); err != nil {
		log.Fatalf("Failed to write output file: %v", err)
	}

	skipped := len(courses) - len(allRequirements) - len(failed)
	fmt.Printf("Successfully parsed %d of %d (%.2f) courses that had requirements (%d courses skipped because they had no requirements)\n", len(allRequirements), len(courses)-skipped, float64(len(allRequirements))/float64(len(courses)-skipped), skipped)
	if len(failed) > 0 {
		slices.Sort(failed)
		fmt.Printf("Failed to parse the following courses: \n%s", strings.Join(failed, "\n"))
	}
}

type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	HasError bool
}

func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.HasError = true
}

func ParseString(input string) (objects.Requirements, bool) {
	trimmed := strings.TrimSpace(input)

	stream := antlr.NewInputStream(trimmed)
	lexer := parser.NewRequirementsLexer(stream)

	lexerErrors := &CustomErrorListener{}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrors)

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewRequirementsParser(tokens)

	parserErrors := &CustomErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(parserErrors)

	tree := p.Prog()
	if lexerErrors.HasError || parserErrors.HasError {
		return objects.Requirements{}, true
	}

	visitor := visitors.NewRequisiteVisitor(tokens)
	visitor.Visit(tree)

	return visitor.Requirements, false
}
