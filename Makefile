tests:
	go test -v ./... -coverprofile=./c.out -coverpkg=./...
	go tool cover -func ./c.out | grep total

# Run all tests and show coverage
tests-cover:
	go test -v ./... -coverprofile=./c.out -coverpkg=./...
	go tool cover -func ./c.out

build:
	go build ./cmd/infernote/main.go

chords:
	go run ./cmd/infernote/main.go