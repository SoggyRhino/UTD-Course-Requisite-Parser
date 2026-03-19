@echo off
setlocal enabledelayedexpansion

set GENERATED=..\parser\requirements_visitor.go
set VISITORS_DIR=..\visitors
set SHOW_MISSING=0

for %%A in (%*) do (
    if "%%A"=="--missing" set SHOW_MISSING=1
)

powershell -Command "Select-String -Path '%GENERATED%' -Pattern 'Visit[A-Z][a-zA-Z]+' -AllMatches -CaseSensitive | ForEach-Object { $_.Matches } | ForEach-Object { $_.Value } | Sort-Object -Unique" > temp_total.txt

powershell -Command "Select-String -Path '%VISITORS_DIR%\*.go' -Pattern 'func \(v \*RequisiteVisitor\) (Visit[A-Z][a-zA-Z]+)' -AllMatches -CaseSensitive | ForEach-Object { $_.Matches } | ForEach-Object { $_.Groups[1].Value } | Sort-Object -Unique" > temp_implemented.txt

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

echo $total = Get-Content 'temp_total.txt'                                                              > temp_coverage.ps1
echo $implemented = Get-Content 'temp_implemented.txt'                                                 >> temp_coverage.ps1
echo $missing = Compare-Object $total $implemented ^| Where-Object { $_.SideIndicator -eq '^<=' } ^| ForEach-Object { $_.InputObject } >> temp_coverage.ps1
echo Write-Output '# Visitor Coverage'                                                                 >> temp_coverage.ps1
echo Write-Output ''                                                                                    >> temp_coverage.ps1
echo foreach ($item in $implemented) { Write-Output ('- [x] ' + $item) }                              >> temp_coverage.ps1
echo foreach ($item in $missing)     { Write-Output ('- [ ] ' + $item) }                              >> temp_coverage.ps1

if "%SHOW_MISSING%"=="1" (
    echo.
    powershell -ExecutionPolicy Bypass -File temp_coverage.ps1
    echo.
)

:: Cleanup
del temp_total.txt temp_implemented.txt temp_coverage.ps1 2>nul