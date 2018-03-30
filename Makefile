files := $(shell find . -name "*.go" | grep -v vendor)

clean:
	goimports -w $(files)

test: clean
	go test

bench: clean
	go test -bench .

commit: test
	git commit
