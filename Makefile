gen:
	protoc -I=api --go_out=. --go-grpc_out=. song.proto