@echo off
setlocal enabledelayedexpansion

java -jar .\antlr.jar -o .\scripts ..\Requirements.g4
if %errorlevel% neq 0 exit /b %errorlevel%

javac -cp ".;java;.\antlr.jar" java\*.java
if %errorlevel% neq 0 exit /b %errorlevel%

java -cp ".;java;.\antlr.jar" FastTester --baseline-file=out/passed_baseline.txt %*

del /q java\*.class *.interp *.tokens *.java *.class 2>nul
