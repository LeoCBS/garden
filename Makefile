BASE_BUILD_IMG = garden
GO_DIR=/go/src/github.com/LeoCBS/garden
RUN_GO=docker run -v `pwd`:$(GO_DIR) -w $(GO_DIR) $(BASE_BUILD_IMG) 

base-build:
	docker build -t $(BASE_BUILD_IMG) .

build: base-build
	$(RUN_GO) go build

build-arm: base-build
	$(RUN_GO) env GOOS=android GOARCH=arm go build

run: build
	./main

check: base-build
	$(RUN_GO) go test -v ./...	
