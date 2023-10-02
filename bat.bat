@echo off
echo Author: Samsonova Alena Ivanovna
echo Course/Group: 3/185

rem Task 2
if exist "%~1" (
echo Text file %~1 exists.
) else (
echo Text file %~1 is missing.
exit /b 1
)

rem Task 3
if exist "%~2" (
echo Executable file %~2 exists.
go run exec.go text.txt output.txt
) else (
echo Executable file %~2 is missing.
exit /b 1
)

rem Task 4
set "directory=hw3"
if not exist "%directory%" (
mkdir "%directory%"
echo Directory %directory% created.
) else (
echo Directory %directory% already exists.
)

rem Task 6
"%~2" "%~1"
set "app_exit_code=%errorlevel%"

rem Task 7
if %app_exit_code% equ 0 (
echo Program completed successfully.
) else (
echo Error code %app_exit_code%.

if %app_exit_code% equ 1 (
echo Error 1: Error reading the file.
) else if %app_exit_code% equ 2 (
echo Error 2: Error writing to the file.
) else (
echo Unknown error.
)
)

rem Task 8
exit /b 0 