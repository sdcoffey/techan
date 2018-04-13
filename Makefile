files := $(shell find . -name "*.go" | grep -v vendor)

clean:
	rm -rf bin
	goimports -w $(files)
	mkdir bin

test: clean
	go test

lint: clean
	golint -set_exit_status
	golint -set_exit_status example

bench: clean
	go test -bench .

commit: test
	git commit

release: test lint
	./scripts/release.sh

coverage: clean
	go test -race -cover -covermode=atomic -coverprofile=bin/coverage.txt
	go tool cover -html bin/coverage.txt
