@echo off
set cnt=0
:loop
if "%1"=="" ( echo off ) else (set /a cnt+=1&shift /1&goto :loop)


echo %cnt%