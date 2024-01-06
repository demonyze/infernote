tests:
	go test -v ./... -coverpkg=./...

build:
	go build ./cmd/main.go

chords:
	go run ./cmd/main.go