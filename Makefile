.PHONY: deps clean build

deps:
	dep ensure

clean:
	rm -fr ./bin

build:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/slackbot ./slack
