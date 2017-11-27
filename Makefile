clean:
	go fmt ./...

test: clean
	go test

bench: clean
	go test -bench .

commit: test
	git commit
