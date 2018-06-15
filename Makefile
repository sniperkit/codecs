build:
	@go build -v ./pkg/*.go

example: example-build example-run

example-build:
	@go build -o bin/example-mimedb example/main.go
	@./bin/example-mimedb

example-run:
	@go run -race example/main.go

install:
	@go install ./pkg/*.go

install-deps:
	@glide install --strip-vendor

install-deps-dev: install-deps
	@go get github.com/golang/lint/golint

update-deps:
	@glide update

update-deps-dev: update-deps
	@go get -u github.com/golang/lint/golint

test:
	@go test -v --race $$(glide novendor)

test-with-coverage:
	@go test --race -cover $$(glide novendor)

test-with-coverage-formatted:
	@go test --race -cover $$(glide novendor) | column -t | sort -r

cover:
	@rm -rf *.coverprofile
	@go test -coverprofile=compressible-go.coverprofile ./pkg/...
	@gover
	@go tool cover -html=compressible-go.coverprofile ./pkg/...

lint: install-deps-dev
	@errors=$$(gofmt -l .); if [ "$${errors}" != "" ]; then echo "$${errors}"; exit 1; fi
	@errors=$$(glide novendor | xargs -n 1 golint -min_confidence=0.3); if [ "$${errors}" != "" ]; then echo "$${errors}"; exit 1; fi

vet:
	@go vet $$(glide novendor)
