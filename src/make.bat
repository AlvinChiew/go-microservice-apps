@REM @ECHO off    
@REM WHERE /q swagger
@REM IF %ERRORLEVEL% NEQ 0 go get -u github.com/go-swagger/go-swagger/cmd/swagger 

@REM swagger generate spec -o ./swagger.yaml --scan-models