build:
	go build -o bin/decritic cmd/decritic/decritic.go

compile:
	GOOS=freebsd GOARCH=386   go build -o bin/decritic-freebsd-386 cmd/decritic/decritic.go
	GOOS=darwin  GOARCH=386   go build -o bin/decritic-darwin-386 cmd/decritic/decritic.go
	GOOS=linux   GOARCH=386   go build -o bin/decritic-linux-386 cmd/decritic/decritic.go
	GOOS=windows GOARCH=386   go build -o bin/decritic-windows-386 cmd/decritic/decritic.go
	GOOS=freebsd GOARCH=amd64 go build -o bin/decritic-freebsd-amd64 cmd/decritic/decritic.go
	GOOS=darwin  GOARCH=amd64 go build -o bin/decritic-darwin-amd64 cmd/decritic/decritic.go
	GOOS=linux   GOARCH=amd64 go build -o bin/decritic-linux-amd64 cmd/decritic/decritic.go
	GOOS=windows GOARCH=amd64 go build -o bin/decritic-windows-amd64 cmd/decritic/decritic.go

dep:
	go get -v -t -d ./...

staticcheck:
	go get -u honnef.co/go/tools/cmd/staticcheck
	`go env GOPATH`/bin/staticcheck ./...

run:
	go run cmd/decritic/decritic.go
