build:
	go build -o bin/decritic decritic.go

compile:
	for dist in $(go tool dist list); do \
		if [ "${dist%/*}" = "windows" ]; then \
			ext=.exe \
		fi \
		GOOS="${dist%/*}" GOARCH="${dist#*/}" go build -o "bin/${GOOS}/${GOARCH}/decritic${ext}" decritic.go \
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

clean:
	go clean
	rm -Rf bin
