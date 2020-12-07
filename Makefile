.PHONY: build

build:
	GOOS=windows GOARCH=386 go build -o netshare86.exe main.go
	GOOS=windows GOARCH=amd64 go build -o netshare64.exe main.go
	GOOS=linux GOARCH=amd64 go build -o netshare64.bin main.go
	GOOS=linux GOARCH=arm go build -o netshareArm.bin main.go
	GOOS=freebsd GOARCH=amd64 go build -o netshareFreebsd64.bin main.go