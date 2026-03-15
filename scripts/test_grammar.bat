@echo off
setlocal enabledelayedexpansion

java -jar .\antlr.jar -o .\scripts ..\Requirements.g4
if %errorlevel% neq 0 exit /b %errorlevel%

javac -cp ".;.\antlr.jar" *.java
if %errorlevel% neq 0 exit /b %errorlevel%

java -cp ".;.\antlr.jar" FastTester

echo.
echo Cleaning up ANTLR-generated files...
del /q *.tokens 2>nul
del /q *.interp 2>nul
del /q *.class  2>nul
del /q requirements*.java 2>nul
echo Done.