files := $(shell find . -name "*.go" | grep -v vendor)

bootstrap:
	go get golang.org/x/lint/golint
	go get honnef.co/go/tools/...

clean:
	goimports -w $(files)

test: clean
	go test

lint:
	golint -set_exit_status
	golint -set_exit_status example
	staticcheck github.com/sdcoffey/techan
	staticcheck github.com/sdcoffey/techan/example

bench: clean
	go test -bench .

commit: test
	git commit

release: clean test lint
	./scripts/release.sh

test-with-coverage:
	go test -race -cover -covermode=atomic -coverprofile=coverage.txt github.com/sdcoffey/techan

view-coverage: test-with-coverage
	go tool cover -html coverage.txt
