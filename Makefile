.PHONY: build

BUILD_DIR=build

build:
	GOOS=windows GOARCH=386 go build -o ${BUILD_DIR}/netshare86.exe main.go
	GOOS=windows GOARCH=amd64 go build -o ${BUILD_DIR}/netshare64.exe main.go
	GOOS=linux GOARCH=amd64 go build -o ${BUILD_DIR}/netshare64.bin main.go
	GOOS=linux GOARCH=arm go build -o ${BUILD_DIR}/netshareArm.bin main.go
	GOOS=freebsd GOARCH=amd64 go build -o ${BUILD_DIR}/netshareFreebsd64.bin main.go

tests_integration:
	go test -v -tags=integration ./...

tests_unit:
	go test -v -tags=unit ./...

run:
	go run main.go