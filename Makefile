.PHONY: build clean deploy

build:
	env GO111MODULE=on GOOS=linux go build -ldflags="-s -w" -o bin/getFolders src/adapters/primary/getFolders/*

test: 
	env GO111MODULE=on go test ./... -cover
	
clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
