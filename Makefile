all: clean build


# go
.PHONY: clean
.PHONY: build
.PHONY: run

clean:
	rm -f out/uuid-srv 
	docker rm -f tjcelaya/uuid-srv && echo ok
	docker rmi -f tjcelaya/uuid-srv && echo ok

build:
	go fmt .
	go mod tidy
	CGO_ENABLED=0 GOOS=linux go build -o out/uuid-srv main.go

run:
	go run main.go

# docker
.PHONY: docker
.PHONY: docker-run

docker: build
	docker build --no-cache -t tjcelaya/uuid-srv .

docker-run: docker
	docker run -it --rm -p 9999:9999 tjcelaya/uuid-srv
