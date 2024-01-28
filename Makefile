tests:
	go test -v ./... -coverpkg=./...

build:
	go build ./cmd/infernote/main.go

chords:
	go run ./cmd/infernote/main.go