gen:
	protoc --proto_path=proto proto/*.proto  --go_out=:pb --go-grpc_out=:pb

dev-up:
	docker-compose -f docker/docker-compose.dev.yml up -d

dev-down:
	docker-compose -f docker/docker-compose.dev.yml down

server:
	PORT=8080 DB_HOST=0.0.0.0 DB_USER=user DB_PASS=pass DB_NAME=user_feature go run cmd/server/main.go -port 8080

client:
	go run cmd/client/main.go -address 0.0.0.0:8080

build:
	docker build -t grpcserver -f docker/Dockerfile .

test-up:
	docker-compose -f docker/docker-compose.test.yml up -d

test-down:
	docker-compose -f docker/docker-compose.test.yml down

.PHONY: gen dev-up dev-down docker-run server client build test-up test-down