.PHONY: compile

protoaja:
	protoc -I ./proto \
		--go_out ./proto --go_opt paths=source_relative \
		--go-grpc_out ./proto --go-grpc_opt paths=source_relative \
		./proto/helloworld/hello_world.proto

protogen:
	protoc -I ./proto \
		--go_out ./proto --go_opt paths=source_relative \
		--go-grpc_out ./proto --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \
		./proto/helloworld/hello_world.proto

protokelpin:
	protoc -I ./proto/helloworld \
		--go_out ./proto/helloworld --go_opt paths=source_relative \
		--go-grpc_out ./proto/helloworld --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./proto/helloworld --grpc-gateway_opt paths=source_relative \
		./proto/helloworld/*.proto