.PHONY:

build:
	go build -o ./.bin/file cmd/main.go

run: build
	./.bin/file