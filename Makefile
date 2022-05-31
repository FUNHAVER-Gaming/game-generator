proto:
	protoc -I ./pkg/proto --go_out=plugins=grpc:./pkg/proto league.proto
