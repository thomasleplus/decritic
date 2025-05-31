default: compile

.PHONY: build compile dep vet staticcheck run test clean

build:
	go build -o bin/decritic decritic.go

compile:
	for dist in $(shell go tool dist list | grep -v -e '^android/'  -e '^ios/' ); do \
		if [ "$${dist%/*}" = "windows" ]; then \
			ext='.exe' ; \
		fi ; \
		echo GOOS="$${dist%/*}" GOARCH="$${dist#*/}" ext="$${ext}" ; \
		GOOS="$${dist%/*}" ; \
		GOARCH="$${dist#*/}" ; \
		go build -o "bin/$${GOOS}/$${GOARCH}/decritic$${ext}" decritic.go ; \
	done

dep:
	go mod download

vet:
	go vet ./...

staticcheck:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	`go env GOPATH`/bin/staticcheck ./...

run:
	go run decritic.go

test:
	go test
	go install
	cat test.csv | ./test.sh

clean:
	go clean
	rm -Rf bin
