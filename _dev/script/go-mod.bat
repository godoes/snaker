@echo off
color 07
title ���²����� GO ģ������
:: file-encoding=GBK
rem by iTanken

cd /d %~dp0/../../
echo 1. ������������...
cd
::& go get -d -u & echo.

go get github.com/onsi/ginkgo@latest
go get github.com/onsi/gomega@latest

echo 2. ����ģ������...
go mod tidy & echo.

:: echo 3. ����ģ�������� vendor Ŀ¼...
:: go mod vendor & echo.

git add .

call "%~dp0/done-time-pause.bat"