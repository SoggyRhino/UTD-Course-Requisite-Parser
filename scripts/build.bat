@echo off
java -jar antlr.jar -Dlanguage=Go -visitor -package parser -o ../parser ../Requirements.g4