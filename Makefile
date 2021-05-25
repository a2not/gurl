
get: 
	go run main.go get http://google.com

run:
	go run main.go

test:
	go test -v ./cmd -count=1

.PHONY: get run test
