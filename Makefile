.PHONY: protos

protos: 
	protoc -I europa/internal/server/auth/protos europa/internal/server/auth/protos/messages.proto --go_out=plugins=grpc:europa/internal/server/auth/protos