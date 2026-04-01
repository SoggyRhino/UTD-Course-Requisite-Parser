# UTD Course Requisite Parser 

ANTLR4 based approach for parsing course requirements. 

Currently, the grammar successfully parses* ~97% of the unique, 25f course requirements. At this point adding the remaining 20 ish 
 inputs is not worth focusing on since they are all mostly unique edge cases and/or typos.

> \*Antlr parses them into an AST, not parsed yet into a useful JSON object.
# Build 

Currently, this is a work in progress. I am building up the visitors from the leaf nodes, so everything will pretty
much not be functional until all visitors are completed. 

Antlr4 is uses grammar to generate go code for a parser, lexer and visitors. This project
ignores the first 2 and only uses visitors. Visitors are functions that are called when a node in the ast is visited.
They essentially act as a visitor for the ast work as reducers where we can take a node, say `course: PREFIX NUMBER`, and 
transform it into a useful object. 

Since antlr4 generates code, you must run the following command
```bash
cd .. scripts 
./build.bat
```

# Scripts
There are also some scripts in the scripts folder to help development. 
> These are AI slop, but since it's not actually part of the project, I couldn't be bothered to make something better.

completed.bat
 - Prints out how many visitors are created and how many are missing
 - `--missing` flag prints out a Markdown list for README.md

extract_inputs.bat 
 - Prints out all the inputs that are captured for a specific rule 
 - Useful for creating the unit tests

test_grammar.bat 
 - Runs the grammar and makes sure that changes don't make the grammar worse 


# Visitor Progress
 - Implemented : 93
 - Missing     : 57
 - Total       : 150
 - Coverage    : 62%