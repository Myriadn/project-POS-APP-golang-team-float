run:
	go run main.go

build:
	go build -o posapp.exe main.go

tidy:
	go mod tidy

test:
	go test -v ./...

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
