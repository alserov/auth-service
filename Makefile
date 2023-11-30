pb:
	protoc -I grpc grpc/proto/auth.proto \
			--go_opt=paths=source_relative --go_out=./grpc/gen \
			--go-grpc_opt=paths=source_relative --go-grpc_out=./grpc/gen grpc/proto/auth.proto
