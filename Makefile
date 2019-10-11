all: build

build:
	CGO_ENABLED=0 go build -o ./bin/media-to-github *.go
