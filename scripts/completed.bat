@echo off
setlocal enabledelayedexpansion

set GENERATED=..\parser\requirements_visitor.go
set VISITORS_DIR=..\visitors
set SHOW_MISSING=0

:: Parse flags
for %%A in (%*) do (
    if "%%A"=="--missing" set SHOW_MISSING=1
)

:: Extract all Visit* methods from generated base
powershell -Command "Select-String -Path '%GENERATED%' -Pattern 'Visit[A-Z][a-zA-Z]+' -AllMatches | ForEach-Object { $_.Matches } | ForEach-Object { $_.Value } | Sort-Object -Unique" > temp_total.txt

:: Extract implemented Visit* methods
powershell -Command "Select-String -Path '%VISITORS_DIR%\*.go' -Pattern 'func \(v \*RequisiteVisitor\) (Visit[A-Z][a-zA-Z]+)' -AllMatches | ForEach-Object { $_.Matches } | ForEach-Object { $_.Groups[1].Value } | Sort-Object -Unique" > temp_implemented.txt

:: Count
for /f %%A in ('find /c /v "" ^< temp_total.txt') do set total=%%A
for /f %%A in ('find /c /v "" ^< temp_implemented.txt') do set implemented=%%A
set /a missing=total-implemented
set /a pct=(implemented*100)/total

echo.
echo  ================================
echo   Visitor Coverage Report
echo  ================================
echo   Implemented : %implemented%
echo   Missing     : %missing%
echo   Total       : %total%
echo   Coverage    : %pct%%%
echo  ================================

if "%SHOW_MISSING%"=="1" (
    echo.
    echo  Missing Visitors:
    echo  -----------------
    powershell -Command "Compare-Object (Get-Content temp_total.txt) (Get-Content temp_implemented.txt) | Where-Object { $_.SideIndicator -eq '<=' } | ForEach-Object { '  - ' + $_.InputObject }"
    echo.
)

:: Cleanup
del temp_total.txt temp_implemented.txt 2>nul