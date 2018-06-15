build:
	@go build -v ./pkg/*.go

example: example-build example-run

example-build-all: example-build-codec-service example-build-codec-yaml example-build-codec-xml example-build-codec-json

example-run-all: example-run-codec-service example-run-codec-yaml example-run-codec-xml example-run-codec-json

example-build-codec-service:
	@go build -o ./bin/example-service-codec example/service-codec/main.go
	@./bin/example-service-codec

example-run-codec-service:
	@go run -race example/service-codec/main.go

example-build-codec-yaml:
	# @go build -o ./bin/example-simple-yaml-codec ./example/simple-yaml-codec/main.go
	# @./bin/example-simple-yaml-codec

example-run-codec-yaml:
	# @go run -race example/simple-yaml-codec/main.go

example-build-codec-xml:
	@go build -o ./bin/example-simple-xml-codec ./example/simple-xml-codec/main.go
	@./bin/example-simple-xml-codec

example-run-codec-xml:
	@go run -race example/simple-xml-codec/main.go

example-build-codec-json:
	@go build -o ./bin/example-simple-json-codec ./example/simple-json-codec/main.go
	@./bin/example-simple-json-codec

example-run-codec-json:
	@go run -race example/simple-json-codec/main.go

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
