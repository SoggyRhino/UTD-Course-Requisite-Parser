# UTD Course Requisite Parser 

ANTLR4 based approach for parsing course requirements. 


# Scripts
The `scripts/` folder contains several utilities for development:

- **`build.bat`**: Generates Go code from the ANTLR4 grammar.
- **`completed.bat`**: Reports visitor implementation coverage.
  - `--missing` flag prints a Markdown list of missing visitors.
- **`extract_inputs.bat`**: Extracts captured inputs for specific rules, useful for unit testing.
- **`test_grammar.bat`**: Validates changes against the grammar.
- **`compile_ts.bat`**: Transpiles the Go `objects` package to TypeScript using [goscript](https://github.com/aperturerobotics/goscript).
  - Requires `goscript` to be installed: `go install github.com/aperturerobotics/goscript/cmd/goscript@latest`
