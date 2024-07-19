gen:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		rpc/schema/schema.proto

run:
	godotenv go run .

test:
	godotenv go test -v -count=1 ./tests/...
