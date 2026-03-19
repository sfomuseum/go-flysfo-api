GOMOD=vendor

vuln:
	govulncheck -show verbose ./...

cli:
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/api cmd/api/main.go
