@echo off
setlocal

:: Check if goscript is installed
where goscript >nul 2>err
if %errorlevel% neq 0 (
    echo [ERROR] goscript not found. Please install it using:
    echo go install github.com/s4wave/goscript/cmd/goscript@latest
    exit /b 1
)

:: Create output directory if it doesn't exist
if not exist "..\out\ts" mkdir "..\out\ts"

echo Compiling Go objects to TypeScript...
goscript compile --dir ..\go --package parser/objects/... --output ..\out\ts --all-dependencies

if %errorlevel% equ 0 (
    echo Done. TypeScript files are in the 'out\ts' directory.

    :: Ensure a tsconfig.json exists with the correct library settings
    if not exist "..\out\ts\tsconfig.json" (
        echo { "compilerOptions": { "target": "es2024", "lib": ["es2024", "dom"], "moduleResolution": "bundler", "allowImportingTsExtensions": true, "noEmit": true, "strict": true, "baseUrl": ".", "paths": { "@goscript/*": ["./@goscript/*"] } } } > "..\out\ts\tsconfig.json"
    )
) else (
    echo [ERROR] Compilation failed.
)