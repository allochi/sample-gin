default: test clean
	go build -v -o bin/main .

run:
	go run .

test:
	go test

.PHONY: clean
clean:
	rm -f bin/*

.PHONY: build-docker
build-docker:
	docker build -t sample-gin .

.PHONY: run-docker
run-docker:
	docker run -it -p 3300:3300 sample-gin