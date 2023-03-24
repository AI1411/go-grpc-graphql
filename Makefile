.PHONY: up gen-proto
up:
	docker-compose --env-file .env.local -f docker-compose.yml up -d --build
upd:
	docker-compose --env-file .env.local -f docker-compose.yml up -d
gen-proto:
	rm -f grpc/*.go && protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --doc_out=./docs --doc_opt=html,index.html grpc/*.proto
gqlgen:
	go run github.com/99designs/gqlgen
logs:
	docker logs -f star
test:
	docker exec -t --env-file .env.test star gotestsum -- -p 1 -count=1 ./...
mockgen-chat:
	mockgen -source ./internal/infra/repository/chat/chat_repository.go -destination=./internal/infra/repository/chat/mock/mock_chat_repository.go
mockgen-user:
	mockgen -source ./internal/infra/repository/user/user_repository.go -destination=./internal/infra/repository/user/mock/mock_user_repository.go
mockgen-room:
	mockgen -source ./internal/infra/repository/room/room_repository.go -destination=./internal/infra/repository/room/mock/mock_room_repository.go
mockgen-tweet:
	mockgen -source ./internal/infra/repository/tweet/tweet_repository.go -destination=./internal/infra/repository/tweet/mock/mock_tweet_repository.go