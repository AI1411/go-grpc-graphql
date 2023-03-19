.PHONY: up gen-proto
up:
	docker-compose --env-file .env.local -f docker-compose.yml up -d --build
upd:
	docker-compose --env-file .env.local -f docker-compose.yml up -d
gen-proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --doc_out=./docs --doc_opt=html,index.html grpc/*.proto
gqlgen:
	go run github.com/99designs/gqlgen
logs:
	docker logs -f star
test:
	docker exec -t --env-file .env.test star gotestsum -- -p 1 -count=1 ./...