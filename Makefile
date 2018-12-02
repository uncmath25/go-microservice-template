.PHONY: clean build deploy remove

clean:
	@echo "*** Cleaning build artifacts ***"
	rm -rf ./bin ./vendor Gopkg.lock *.log

build: clean
	@echo "*** Building go binaries with specified dependencies ***"
	if [ ! -f Gopkg.toml ]; then dep init; fi
	dep ensure -v
	go build -ldflags="-s -w" -o bin/run_http_server cmd/httpserver/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/run_lambda_server cmd/lambdaserver/main.go

deploy: clean build
	@echo "*** Deploying project to serverless ***"
	sls deploy -v --aws-profile colton

remove:
	@echo "*** Removing serverless deployment ***"
	sls remove -v --aws-profile colton
