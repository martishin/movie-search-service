run:
	go run cmd/api/main.go

run-postgres:
	docker compose up postgres

generate-sql:
	sqlc generate

test:
	go test ./... -v

build-mac:
	docker build --build-arg TARGETOS=darwin --build-arg TARGETARCH=arm64 -t movie-search-server .

run-docker:
	docker compose up -d

stop-docker:
	docker compose down

remove-docker:
	docker compose down -v

logs-docker:
	docker compose logs -f

login-docker:
	aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 100381574725.dkr.ecr.us-east-1.amazonaws.com

build:
	docker build -t movie-search-server .

upload:
	docker tag movie-search-server:latest 100381574725.dkr.ecr.us-east-1.amazonaws.com/movie-search/server:latest && docker push 100381574725.dkr.ecr.us-east-1.amazonaws.com/movie-search/server:latest

deploy:
	aws ecs update-service --cluster movie-search-cluster --service movie-search-server --force-new-deployment
