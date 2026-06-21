# UTD Course Requisite Parser 

ANTLR4 based approach for parsing UTD course requirements. The parsing and generation of course requirement is intended to be 
run once through the Go library (see main.go for sample usage); while the evaluation from student info supports client-side 
evaluation through TypeScript generation (credit to [goscript](https://github.com/aperturerobotics/goscript)).

The project supports over 90% of unique course requirements (as of Spring Semester 2026), the remaining courses contain
one-off requirements and/or typos that would be disproportionately onerous to parse. This highlights a small flaw in the main 
approach in this parser: course requirements are not a formal grammar, just pretty close. But at the end of the day, this was 
mostly an exuse to build a project using ANTLR4, and of those that can't be parsed, most are random, tiny courses, so the 
vast majority of students wouldn't even notice. 

There is a live demo of this project hosted in GitHub pages [here](). The tect stack is: React + shadcn ui. 





# Development



## Scripts
The `scripts/` folder contains several utilities for development:

- **`build.bat`**: Generates Go code from the ANTLR4 grammar.
- **`completed.bat`**: Reports visitor implementation coverage.
  - `--missing` flag prints a Markdown list of missing visitors.
- **`extract_inputs.bat`**: Extracts captured inputs for specific rules, useful for unit testing.
- **`test_grammar.bat`**: Validates changes against the grammar.
- **`compile_ts.bat`**: Transpiles the Go `objects` package to TypeScript using [goscript](https://github.com/aperturerobotics/goscript).
  - Requires `goscript` to be installed: `go install github.com/aperturerobotics/goscript/cmd/goscript@latest`
