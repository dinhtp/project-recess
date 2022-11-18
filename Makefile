.PHONY: install test build serve clean pack deploy ship
TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

export TAG

install:
	go get .

build: install
	go build -ldflags "-X main.version=$(TAG)" -o ./bin/project-recess .

serve: build
	./bin/go-company-service serve grpc

clean:
	rm -f ./bin/go-company-service

dev:
	GOOS=linux make build
	docker build -t localhost:5000/project-recess .
	docker push localhost:5000/project-recess