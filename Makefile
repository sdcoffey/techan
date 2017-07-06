clean:
	go fmt ./...

test: clean
	go test -v

bench: clean
	go test -bench .

commit: test
	git commit