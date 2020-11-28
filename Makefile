compile:
	GOOS=freebsd GOARCH=386   go build -o bin/decritic-freebsd-386 decritic.go
	GOOS=darwin  GOARCH=386   go build -o bin/decritic-darwin-386 decritic.go
	GOOS=linux   GOARCH=386   go build -o bin/decritic-linux-386 decritic.go
	GOOS=windows GOARCH=386   go build -o bin/decritic-windows-386 decritic.go
	GOOS=freebsd GOARCH=amd64 go build -o bin/decritic-freebsd-amd64 decritic.go
	GOOS=darwin  GOARCH=amd64 go build -o bin/decritic-darwin-amd64 decritic.go
	GOOS=linux   GOARCH=amd64 go build -o bin/decritic-linux-amd64 decritic.go
	GOOS=windows GOARCH=amd64 go build -o bin/decritic-windows-amd64 decritic.go

build:
	go build -o bin/decritic cmd/decritic/decritic.go

run:
	go run decritic.go
