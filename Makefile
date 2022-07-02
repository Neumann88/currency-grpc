gen:
	protoc --go_out=. --go-grpc_out=. proto/currency.proto
.PHONY: gen

server-run:
	go run ./server
.PHONY: server-run

client-run:
	go run ./client
.PHONY: server-run