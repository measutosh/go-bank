#.PHONY: run
#un: main
#	./$<
#
#main: *.go go.mod
#	go build -o $@ .
#
#.PHONY: all
#all: main


build:
	@go build -o bin/go-bank
run: build
	@./bin/go-bank
test:
	@go test -v ./..