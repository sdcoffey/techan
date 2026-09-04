GO ?= go
GOIMPORTS_VERSION := v0.49.0
STATICCHECK_VERSION := v0.8.1
GOIMPORTS = $(GO) run golang.org/x/tools/cmd/goimports@$(GOIMPORTS_VERSION)
STATICCHECK = $(GO) run honnef.co/go/tools/cmd/staticcheck@$(STATICCHECK_VERSION)

.PHONY: bootstrap format clean test lint bench commit release test-with-coverage view-coverage

bootstrap:
	$(GO) install golang.org/x/tools/cmd/goimports@$(GOIMPORTS_VERSION)
	$(GO) install honnef.co/go/tools/cmd/staticcheck@$(STATICCHECK_VERSION)

format:
	$(GOIMPORTS) -w .

clean: format

test:
	$(GO) test ./...

lint:
	$(GO) vet ./...
	$(STATICCHECK) ./...

bench:
	$(GO) test -run '^$$' -bench . -benchmem ./...

commit: test
	git commit

release: format test lint
	./scripts/release.sh

test-with-coverage:
	$(GO) test -race -covermode=atomic -coverprofile=coverage.txt ./...

view-coverage: test-with-coverage
	$(GO) tool cover -html coverage.txt
