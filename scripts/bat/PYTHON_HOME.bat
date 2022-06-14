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

setx /M "CLS_PYTHON_HOME_37" "D:\Python\Python37"
setx /M "CLS_PYTHON_HOME_38" "D:\Python\Python38"
setx /M "CLS_PYTHON_HOME_39" "D:\Python\Python39"
setx /M "CLS_PYTHON_HOME_310" "D:\Python\Python310"

:again
cls
echo.
echo.
echo ��ѡ��Ҫ���õ�PYTHON_HOME
echo [ 37: PYTHON_HOME_37 ]
echo [ 38: PYTHON_HOME_38 ]
echo [ 39: PYTHON_HOME_39 ]
echo [310: PYTHON_HOME_310]
echo.
echo.
echo ��ѡ��Ҫִ�еĲ���
set /p num=

cls
echo.
echo.
setx /M "CLS_PYTHON_HOME" "%%CLS_PYTHON_HOME_%num%%%"
echo PYTHON_HOME �������

pause



