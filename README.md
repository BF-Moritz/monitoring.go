# monitoring.go

## gRPC compilation:

`protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc_api/monitoring.proto`

## server

`go run server/server.go -c config/server.config.yml`

## client

`go run client/client.go -c config/client.config.yml`

## rest api

`go run api/api.go -c config/api.config.yml`
