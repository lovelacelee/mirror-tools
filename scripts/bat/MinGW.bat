@echo off

goto check_Permissions

:check_Permissions

    net session >nul 2>&1
    if %errorLevel% == 0 (
        echo Success: Administrative permissions confirmed.
    ) else (
        echo Failure: ��ʹ�ù���ԱȨ�޴�Ӧ��.
        pause
        exit
    )

setx /M "mingw810_64" "D:\Qt\Qt6.2\Tools\mingw810_64"
setx /M "mingw810_32" "D:\Qt\Qt6.2\Tools\mingw810_32"

:again
cls
echo.
echo.
echo ��ѡ��Ҫ���õ�MINGW_HOME
echo [ 32: mingw810_32 ]
echo [ 64: mingw810_64 ]
echo.
echo.
echo ��ѡ��Ҫִ�еĲ���
set /p num=

cls
echo.
echo.
setx /M "MINGW_HOME" "%%mingw810_%num%%%"
echo PYTHON_HOME �������

pause



