.PHONY: up gen-proto
up:
	docker-compose --env-file .env.local -f docker-compose.yml up -d --build
gen-proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --doc_out=./docs --doc_opt=html,index.html grpc/*.proto
gqlgen:
	go run github.com/99designs/gqlgen