@echo off
TITLE Go Cross-Compilation Script

ECHO Starting cross-compilation for bot...

REM --- Configuration ---
SET LDFLAGS="-s -w"
SET SOURCE_FILE="main.go"
SET OUTPUT_BASE="skygo"
SET OUTPUT_PATH="dist"
REM -------------------

REM Ensure the Go toolchain is in your PATH
REM You might need to adjust this depending on your Go installation
REM For typical installs, Go is already in the PATH

REM --- Build for Windows (amd64) ---
ECHO Building for Windows (amd64)...
SET GOOS=windows
SET GOARCH=amd64
go build -ldflags=%LDFLAGS% -o %OUTPUT_PATH%/%OUTPUT_BASE%_%GOOS%_%GOARCH%.exe %SOURCE_FILE%
IF %ERRORLEVEL% NEQ 0 (
    ECHO "Error building for Windows (%GOARCH%)"
    GOTO end
)
ECHO Successfully built %OUTPUT_BASE%_%GOOS%_%GOARCH%.exe

REM --- Build for macOS (amd64) ---
REM Requires a Go toolchain installed on the build machine that supports the darwin GOOS
ECHO Building for macOS (amd64)...
SET GOOS=darwin
SET GOARCH=amd64
go build -ldflags=%LDFLAGS% -o %OUTPUT_PATH%/%OUTPUT_BASE%_%GOOS%_%GOARCH% %SOURCE_FILE%
IF %ERRORLEVEL% NEQ 0 (
    ECHO "Error building for macOS (%GOARCH%). Make sure your Go installation supports darwin/%GOARCH%."
    GOTO end
)
ECHO Successfully built %OUTPUT_BASE%_%GOOS%_%GOARCH%

REM --- Build for macOS (arm64) ---
REM Requires a Go toolchain installed on the build machine that supports the darwin GOOS
ECHO Building for macOS (arm64)...
SET GOOS=darwin
SET GOARCH=arm64
go build -ldflags=%LDFLAGS% -o %OUTPUT_PATH%/%OUTPUT_BASE%_%GOOS%_%GOARCH% %SOURCE_FILE%
IF %ERRORLEVEL% NEQ 0 (
    ECHO "Error building for macOS (%GOARCH%). Make sure your Go installation supports darwin/%GOARCH%."
    GOTO end
)
ECHO Successfully built %OUTPUT_BASE%_%GOOS%_%GOARCH%

REM --- Build for Linux (amd64) ---
REM Requires a Go toolchain installed on the build machine that supports the linux GOOS
ECHO Building for Linux (amd64)...
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags=%LDFLAGS% -o %OUTPUT_PATH%/%OUTPUT_BASE%_%GOOS%_%GOARCH% %SOURCE_FILE%
IF %ERRORLEVEL% NEQ 0 (
    ECHO "Error building for Linux (%GOARCH%). Make sure your Go installation supports linux/%GOARCH%."
    GOTO end
)
ECHO Successfully built %OUTPUT_BASE%_%GOOS%_%GOARCH%

REM --- Build for Linux (arm64) ---
REM Requires a Go toolchain installed on the build machine that supports the linux GOOS
ECHO Building for Linux (arm64)...
SET GOOS=linux
SET GOARCH=arm64
go build -ldflags=%LDFLAGS% -o %OUTPUT_PATH%/%OUTPUT_BASE%_%GOOS%_%GOARCH% %SOURCE_FILE%
IF %ERRORLEVEL% NEQ 0 (
    ECHO "Error building for Linux (%GOARCH%). Make sure your Go installation supports linux/%GOARCH%."
    GOTO end
)
ECHO Successfully built %OUTPUT_BASE%_%GOOS%_%GOARCH%

ECHO "All requested packages built successfully!"

:end
ECHO Build process finished.
PAUSE