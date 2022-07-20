gen:
	protoc --go_out=. --go-grpc_out=. proto/*.proto
.PHONY: gen

srv:
	go run ./server
.PHONY: server-run

clt:
	go run ./client
.PHONY: server-run