build:
	go build -v

clean:
	go clean -x

coverage:
	go tool cover -html coverage.txt -o coverage.html

test:
	go test -race -v -coverprofile coverage.txt ./...

