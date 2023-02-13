default: dev

dev:
	air
	#go run github.com/tunght13488/realworld-go

build:
	go build github.com/tunght13488/realworld-go

run:
	./realworld-go

.PHONY: default build run