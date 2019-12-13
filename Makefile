default: test clean
	go build -v -o bin/main .

run:
	go run .

test:
	go test

.PHONY: clean
clean:
	rm -f bin/*
