.PHONY: build clean deploy

build:
	env GO111MODULE=on GOOS=linux go build -ldflags="-s -w" -o bin/getfolders src/adapters/primary/getfolders/*
	env GO111MODULE=on GOOS=linux go build -ldflags="-s -w" -o bin/getfiles src/adapters/primary/getfiles/*
test: 
	env GO111MODULE=on go test ./... -cover
	
clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose

local: build
	sam local start-api
