@echo off

if NOT EXIST %userprofile%\.foxpi (
    mkdir %userprofile%\.foxpi
)

go build -o %userprofile%\.foxpi\foxpi.exe

cp ..\foxpi.json %userprofile%\.foxpi

:: Check if the session is elevated
fsutil dirty query %systemdrive% >nul 2>&1

if %ERRORLEVEL% equ 0 (
    reg add "HKLM\SOFTWARE\Mozilla\NativeMessagingHosts\foxpi" /ve /t REG_SZ /d "%USERPROFILE%\.foxpi\foxpi.json" /f

    if %ERRORLEVEL% equ 0 (
        echo Registry key was successfully set.
    ) else (
        echo Failed to set the registry key.
    )
) else (
    echo The script is not running as administrator. Please add the registry key manually:
    echo reg add "HKLM\SOFTWARE\Mozilla\NativeMessagingHosts\foxpi" /ve /t REG_SZ /d "%USERPROFILE%\.foxpi\foxpi.json" /f
)

