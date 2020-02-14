proto:
	@protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/s2.proto
build:
	@docker build -t s2 .
run:
	@docker run -e MICRO_SERVER_ADDRESS=:50052 --rm --name s2 s2

.PHONY: proto build run
