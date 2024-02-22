
# 根据 .proto 生成 .go   .mock
generate:
	mkdir -p ./api/proto/gen
	# user
	protoc --go_out=./api/proto/gen --go-grpc_out=./api/proto/gen ./api/proto/user/v1/user.proto

mock:
	# user
	mockgen -source=./api/proto/gen/user/v1/user_grpc.pb.go -package=artmocks -destination=./api/proto/gen/user/v1/mocks/user_grpc.mock.go

clean:
	rm -rf ./api/proto/gen/*
# Default target (run 'make' without arguments)
all: clean generate mock