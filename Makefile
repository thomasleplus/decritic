build:
	go build -o bin/decritic decritic.go

compile:
	GOOS=linux   GOARCH=amd64 go build -o bin/decritic-linux-amd64 decritic.go
	GOOS=linux   GOARCH=386   go build -o bin/decritic-linux-386 decritic.go
	GOOS=linux   GOARCH=arm   go build -o bin/decritic-linux-arm decritic.go
	GOOS=linux   GOARCH=arm64 go build -o bin/decritic-linux-arm64 decritic.go
	GOOS=darwin  GOARCH=amd64 go build -o bin/decritic-darwin-amd64 decritic.go
	GOOS=windows GOARCH=amd64 go build -o bin/decritic-windows-amd64 decritic.go
	GOOS=windows GOARCH=386   go build -o bin/decritic-windows-386 decritic.go
	GOOS=freebsd GOARCH=amd64 go build -o bin/decritic-freebsd-amd64 decritic.go
	GOOS=freebsd GOARCH=386   go build -o bin/decritic-freebsd-386 decritic.go

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
