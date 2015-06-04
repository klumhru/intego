all: build container

build:
	go build

container: build
	docker build -t intego/server .

run-container: container
	docker run -p 8080:8080 -it intego/server

run:
	fresh
