# UTD Course Requisite Parser 

[![Go Report Card](https://goreportcard.com/badge/github.com/SoggyRhino/UTD-Course-Requisite-Parser/go)](https://goreportcard.com/report/github.com/SoggyRhino/UTD-Course-Requisite-Parser/go)
[![Build Status](https://github.com/SoggyRhino/UTD-Course-Requisite-Parser/actions/workflows/go-tests.yml/badge.svg)](https://github.com/SoggyRhino/UTD-Course-Requisite-Parser/actions)
[![codecov](https://codecov.io/gh/SoggyRhino/UTD-Course-Requisite-Parser/branch/main/graph/badge.svg)](https://codecov.io/gh/SoggyRhino/UTD-Course-Requisite-Parser)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


ANTLR4 based approach for parsing UTD course requirements (the requirements needed for a student to be able to take a course, 
essentially its prerequisites). The parsing and generation of course requirement is intended to be run once through the Go 
library (see main.go for sample usage); while the evaluation from student info supports client-side evaluation through 
TypeScript generation (credit to [goscript](https://github.com/s4wave/goscript)).

The project supports over 95% of unique course requirements (as of Spring Semester 2026). This highlights a small flaw in 
this parser: course requirements are not a formal grammar, just pretty close, but at the end of the day, this was mostly an
excuse to build a project using ANTLR4. Known limitations include: typos, paraphrasing, and just one-off conditions. 

There is a live demo of this project hosted in GitHub pages [here](https://soggyrhino.github.io/UTD-Course-Requisite-Parser/). The tech stack is: React + shadcn ui. 

## Examples 

### JSON 

Here is an example output for the requirements of CS 4390:

Input: 

`Prerequisite: CE 3345 or CS 3345 or SE 3345 or TE 3345 or equivalent. Credit cannot be received for both courses, (CE 4390 or CS 4390 or TE 4390) and EE 4390.`

This json represents the `Requirements` object used to evaluate whether a student is eligible for CS 4390:
```json
{
  "pre_reqs": {
    "type": "alternative",
    "condition": {
      "type": "or",
      "conditions": [
        {
          "type": "course",
          "course": {
            "prefix": "CE",
            "number": "3345"
          }
        },
        {
          "type": "course",
          "course": {
            "prefix": "CS",
            "number": "3345"
          }
        },
        {
          "type": "course",
          "course": {
            "prefix": "SE",
            "number": "3345"
          }
        },
        {
          "type": "course",
          "course": {
            "prefix": "TE",
            "number": "3345"
          }
        }
      ]
    }
  },
  "rules": [
    {
      "type": "credit_for",
      "courses": {
        "type": "and",
        "courses": [
          {
            "type": "or",
            "courses": [
              {
                "type": "simple",
                "course": [
                  {
                    "prefix": "CE",
                    "number": "4390"
                  }
                ]
              },
              {
                "type": "simple",
                "course": [
                  {
                    "prefix": "CS",
                    "number": "4390"
                  }
                ]
              },
              {
                "type": "simple",
                "course": [
                  {
                    "prefix": "TE",
                    "number": "4390"
                  }
                ]
              }
            ]
          },
          {
            "type": "simple",
            "course": [
              {
                "prefix": "EE",
                "number": "4390"
              }
            ]
          }
        ]
      }
    }
  ]
}


```

### Demo 
![CS 4390.png](images/CS%204390.png)


## Architecture 

 - The grammar is defined using ANTLR4, and through this the parser code is generated. 
 - The visitor pattern is used to recursively build an AST of conditions 
    - example: `course OR course`
      - Visit the first and create a `CourseCondition` 
      - Visit the second and create another `CourseCondition`
      - Then visit the OR and return an `OrCondition` that contains both `CourseCondition`s
   - There are many edge cases here, and conditions need to be hoisted up to the root, but the 
     bottom line is that we traverse the AST and emit go objects at each node
 - Then once all nodes are visited emit a `Requirements` Object
 - Next there is the logic engine that recursively evaluates conditions and returns a status (pass, fail, possible pass,  
    etc.) given what is known about the student
 - Finally, this evaluation part of the code was written in such that it supports transpilation to TypeScript and 
   can be used client-side to evaluate course requirements, as seen in the demo.
 - The demo site is built with React and shadcn ui.


## Development

- This project has extensive test coverage, with over 500 unit tests averaging about 95% coverage for every file. 
   - This was a necessity in protecting against regression when I was iteratively expanding 
     the visitors as the grammar is so large 
- The project grammar is around 490 lines long 
   - Approximately 150 grammar rules

### Scripts
The `scripts/` folder contains several utilities for development:

- **`build.bat`**: Generates Go code from the ANTLR4 grammar.
- **`completed.bat`**: Reports visitor implementation coverage.
  - `--missing` flag prints a Markdown list of missing visitors.
- **`extract_inputs.bat`**: Extracts captured inputs for specific rules, useful for unit testing.
- **`test_grammar.bat`**: Validates changes against the grammar.
- **`compile_ts.bat`**: Transpiles the Go `objects` package to TypeScript using [goscript](https://github.com/s4wave/goscript).
  - Requires `goscript` to be installed: `go install github.com/s4wave/goscript/cmd/goscript@latest`

### Usage 

#### Go 
> Requires Go 1.26+

Sample code (generates `static/requirements.json`): 
  - `cd go`
  - `go run main.go`

#### Website 
> Requires Node.js

 Build typescript source 
 - `cd scripts`
 - `compile_ts.bat`

 Run the site locally:
 - `cd site`
 - `npm install`
 - `npm run dev`
 