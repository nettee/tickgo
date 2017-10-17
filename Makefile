.PHONY: proto
proto:
	protoc -I tick/ tick/tick.proto --go_out=plugins=grpc:tick

.PHONY: start-server
start-server:
	go run server/server.go

.PHONY: start-client
start-client:
	go run client/client.go
