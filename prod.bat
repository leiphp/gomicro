@echo off

start "prod1" go run main.go --server_address :8081&
start "prod2" go run main.go --server_address :8082&
start "prod3" go run main.go --server_address :8083
pause