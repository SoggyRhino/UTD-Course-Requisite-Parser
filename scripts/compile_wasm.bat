@echo off
setlocal

cd /d "%~dp0..\go"

echo Compiling Go project to WebAssembly using standard Go...
if not exist "..\out" mkdir "..\out"
set GOOS=js
set GOARCH=wasm
go build -o ..\out\main.wasm .\main.go

if %ERRORLEVEL% equ 0 (
    echo Successfully compiled to out\main.wasm
    FOR /F "tokens=*" %%g IN ('go env GOROOT') do (
        echo Copying wasm_exec.js from %%g...
        copy "%%g\misc\wasm\wasm_exec.js" "..\out\wasm_exec.js" >nul
    )
    
    echo Also copying to site\public for the web application...
    if not exist "..\site\public" mkdir "..\site\public"
    copy "..\out\main.wasm" "..\site\public\main.wasm" >nul
    copy "..\out\wasm_exec.js" "..\site\public\wasm_exec.js" >nul
) else (
    echo Failed to compile to WebAssembly.
)

endlocal
