@echo off

IF "%~1"=="" (
	ECHO Falta un parametro ^(ej: %0 centinela^)
	EXIT /B 1
)

SET current_path=%cd%
SET zip_file=build_%1.zip
SET project_folder=C:\Users\Haintech\1373PY-T4G-V4\portal_t4g
SET command=.\node_modules\.bin\ng build -c=%1 --outputHashing=all

GOTO COMPILE

:COMPILE
ECHO Ejecutando script para build de %1;
CD %project_folder%
ECHO Compilando proyecto... esto puede tardar hasta 5 minutos
START /B/W %command%
ECHO Compilación terminada
EXIT /B && GOTO ZIPFILE

:ZIPFILE
ECHO Generando compresión en %zip_file%
CD %project_folder%\dist
ZIP -r %zip_file% *
ECHO Archivo comprimido, moviendo %zip_file% al directorio actual...
MOVE %zip_file% %current_path%
TIMEOUT /t 3 /nobreak > nul
CD %current_path%
ECHO Proceso terminado ya puedes subir el build %zip_file% al proyecto de %1