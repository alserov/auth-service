pb:
	protoc -I grpc grpc/proto/auth.proto \
			--go_opt=paths=source_relative --go_out=./grpc/gen/auth \
			--go-grpc_opt=paths=source_relative --go-grpc_out=./grpc/gen/auth grpc/proto/auth.proto
