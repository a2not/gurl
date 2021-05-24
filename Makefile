
get: 
	go run main.go get http://google.com

run:
	go run main.go

test:
	go test ./cmd

.PHONY: get run test
