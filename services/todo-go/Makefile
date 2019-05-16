.PHONY: build clean deploy

build:

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
