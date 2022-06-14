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

setx /M "JAVA_HOME_17" "D:\Program Files\Eclipse Adoptium\jdk-17.0.1.12-hotspot"
setx /M "JAVA_HOME_8" "D:\ConnardApp\JavaEnv\JDK\jdk8u312-b07"
setx /M "JAVA_HOME_16" "D:\ConnardApp\JavaEnv\JDK\jdk-16.0.2+7"

:again
cls
echo.
echo.
echo ��ѡ��Ҫ���õ�JAVA_HOME
echo [ 8: JAVA_HOME_8 ]
echo [16: JAVA_HOME_16]
echo [17: JAVA_HOME_17]
echo.
echo.
echo ��ѡ��Ҫִ�еĲ���
set /p num=

cls
echo.
echo.
setx /M "JAVA_HOME" "%%JAVA_HOME_%num%%%"
setx /M "JRE_HOME" "%%JAVA_HOME_%num%%%\jre"
echo JAVA_HOME �������

pause



