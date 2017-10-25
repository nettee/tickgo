.PHONY: proto
proto:
	protoc -I tick/ tick/tick.proto --go_out=plugins=grpc:tick

.PHONY: start-server
start-server:
	go run server/server.go :9040

.PHONY: start-client1 start-client2 start-client3 start-client4 start-client5 start-client6
start-client1:
	go run client/client.go localhost:9040 user1 pass1
start-client2:
	go run client/client.go localhost:9040 user2 pass2
start-client3:
	go run client/client.go localhost:9040 user3 pass3
start-client4:
	go run client/client.go localhost:9040 user4 pass4
start-client5:
	go run client/client.go localhost:9040 user5 pass5
start-client6:
	go run client/client.go localhost:9040 user6 pass6
