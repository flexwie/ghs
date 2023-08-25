test:
	go test ./... -cover

build: test
	go build -o dist/ghs main.go

run:
	go run main.go