@echo off
java -jar antlr.jar -Dlanguage=Go -visitor -package parser -o ../go/parser ../Requirements.g4