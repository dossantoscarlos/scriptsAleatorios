@echo off

REM Cria a pasta build se não existir
if not exist "build" (
    mkdir "build"
)

REM Compila o programa para o diretório build com o nome app.exe
go build -o app.exe . 

REM Verifica se a compilação foi bem-sucedida
if %errorlevel% neq 0 (
    echo Erro ao compilar o programa!
    exit /b %errorlevel%
)
echo Build realizado com sucesso!

move app.exe build/app.exe
if %errorlevel% neq 0 (
    echo Error ao mover arquivo!
    exit /b %errorlevel%
)

echo Movido com sucesso!!!
