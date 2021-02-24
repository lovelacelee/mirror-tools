@echo off
set cnt=0
:loop
if "%1"=="" ( echo off ) else (set /a cnt+=1&shift /1&goto :loop)


if not %cnt% == 3 (
    echo "Usage: %0 [project name] [spider name] [url]"
    pause
    exit
)

REM python scrapy required
REM @param1: project name & directory
REM @param2: spider name
REM @param3: start url

echo Create scrapy project:  [%1] [%2] [%3]

scrapy startproject %1
cd %1
scrapy genspider %2 %3
cd ../