tests:
	go test -v ./... -coverpkg=./...

build:
	go build ./cmd/app/main.go

chords:
	go run ./cmd/app/main.go