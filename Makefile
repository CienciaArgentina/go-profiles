BINARY=go-profiles
test: 
	go test -v -cover -covermode=atomic ./...

engine:
	go build -o ${BINARY} .


unittest:
	go test -short  ./...

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker-compose build

run:
	docker-compose up --build -d

status:
	docker-compose ps

stop:
	docker-compose down

lint-prepare:
	@echo "Installing golangci-lint" 
	@which golangci-lint > /dev/null 2>&1 || go get github.com/golangci/golangci-lint

lint: lint-prepare
	golangci-lint run ./...

.PHONY: clean install unittest build docker run stop vendor lint-prepare lint
