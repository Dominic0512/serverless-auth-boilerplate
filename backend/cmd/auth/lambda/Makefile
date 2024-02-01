SLS_CMD_PATH = ../../../node_modules/.bin/serverless

build:
	env GOARCH=amd64 env GOOS=linux go build -ldflags="-s -w" -o bin/main .

deploy: build
	${SLS_CMD_PATH} --stage prod --aws-profile me_dev

list:
	${SLS_CMD_PATH} list --stage prod --aws-profile me_dev

remove:
	${SLS_CMD_PATH} --stage prod --aws-profile me_dev

offline:
	${SLS_CMD_PATH} offline --stage dev --useDocker