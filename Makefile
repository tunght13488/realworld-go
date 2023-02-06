default: build run

build:
	go build github.com/tunght13488/realworld-go

run:
	./realworld-go

.PHONY: default build run