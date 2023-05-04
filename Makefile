build:
	go build -v

clean:
	go clean -x

test:
	go test -race -v -coverprofile coverage.txt ./...
	go tool cover -html coverage.txt -o coverage.html
