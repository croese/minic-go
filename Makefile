build:
	go build -o cmd/minic/monkey cmd/minic/main.go

runInterp:
	go run cmd/minic/main.go

test:
	go test ./...