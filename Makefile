build: deps
	GO111MODULE=on
	make clean
	sam build

recompile:
	make build
	make start-local

deps:
	go mod tidy

start-dev:
	supervisor --no-restart-on exit -e go -i bin --exec make -- recompile

start-local:
	sam local start-api

# watch:
# 	make supervise & make start-local

clean:
	rm -rf .aws-sam/*

test:
	go test -coverprofile=coverage.out ./...

deploy-local:
	samlocal deploy --stack-name shortener-url-serverless-golang --capabilities CAPABILITY_IAM