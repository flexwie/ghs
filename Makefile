test:
	go test ./... -cover

build: test
	rm -rf dist
	go build -o dist/ghs main.go

run:
	go run main.go