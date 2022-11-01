build: deps
	make recompile

recompile:
	GO111MODULE=on
	make clean
	sam build
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