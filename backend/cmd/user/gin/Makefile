SLS_CMD_PATH = ../../node_modules/.bin/serverless

build:
	env GOARCH=amd64 env GOOS=linux go build -ldflags="-s -w" -o bin/main .

dev: build
	go run .