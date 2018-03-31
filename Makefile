SERVER_OUT := "bin/server"
CLIENT_OUT := "bin/client"
API_OUT := "api/api.pb.go"
API_REST_OUT := "api/api.pb.gw.go"
SERVER_PKG_BUILD := "server/main.go"
CLIENT_PKG_BUILD := "client/main.go"
GOPATH=$(shell go env GOPATH)

.PHONY: all api server client

all: clean api server client

api/api.pb.go: 
	protoc -I/usr/local/include -I api/ \
    -I${GOPATH}/src \
    -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --go_out=google/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:api \
    api/api.proto

api/api.pb.gw.go:
	protoc -I/usr/local/include -I api/ \
    -I${GOPATH}/src \
    -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --grpc-gateway_out=logtostderr=true:api \
    api/api.proto

api: api/api.pb.go api/api.pb.gw.go

server:
	GOOS=darwin go build -i -v -o $(SERVER_OUT) $(SERVER_PKG_BUILD)

client:
	GOOS=darwin go build -i -v -o $(CLIENT_OUT) $(CLIENT_PKG_BUILD)

clean:
	rm $(SERVER_OUT) $(CLIENT_OUT) $(API_OUT) $(API_REST_OUT)
