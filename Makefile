.PHONY: up upd down star gen-proto gqlgen logs test mockgen-chat mockgen-user mockgen-room mockgen-tweet cp-schema
up:
	docker-compose --env-file .env.local -f docker-compose.yml up -d --build
upd:
	docker-compose --env-file .env.local -f docker-compose.yml up -d
down:
	docker-compose --env-file .env.local -f docker-compose.yml down
star:
	docker-compose --env-file .env.local -f docker-compose.yml exec star bash
gen-proto:
	rm -f grpc/*.go && protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --doc_out=./docs --doc_opt=html,index.html grpc/*.proto
gqlgen:
	go run github.com/99designs/gqlgen
logs:
	docker logs -f star
test:
	docker exec -t --env-file .env.test star gotestsum -- -p 1 -count=1 ./...
test-with-coverage:
	docker exec -t --env-file .env.test star richgo test -cover -p 1 -count=1 ./...
mockgen-chat:
	mockgen -source ./internal/infra/repository/chat/chat_repository.go -destination=./internal/infra/repository/chat/mock/mock_chat_repository.go
mockgen-user:
	mockgen -source ./internal/infra/repository/user/user_repository.go -destination=./internal/infra/repository/user/mock/mock_user_repository.go
mockgen-room:
	mockgen -source ./internal/infra/repository/room/room_repository.go -destination=./internal/infra/repository/room/mock/mock_room_repository.go
mockgen-tweet:
	mockgen -source ./internal/infra/repository/tweet/tweet_repository.go -destination=./internal/infra/repository/tweet/mock/mock_tweet_repository.go
mockgen-category:
	mockgen -source ./internal/infra/repository/category/category_repository.go -destination=./internal/infra/repository/category/mock/category_tweet_repository.go
cp-schema:
	cat ./DDL/*.sql > ./DDL/scripts/schema.sql
cover:
	docker exec -t --env-file .env.test star go test -cover -- -p 1 -count=1 ./... -coverprofile=cover.out.tmp
    # 自動生成コードをカバレッジ対象から外し、カバレッジファイルを作成
	cat cover.out.tmp | grep -v "**_mock.go" > cover.out
	rm cover.out.tmp
	docker exec -t --env-file .env.test star go tool cover -html=cover.out -o cover.html
	open cover.html