run:
	go run cmd/api/main.go

start-all:
	docker compose up postgres redis

stop-all:
	docker compose down postgres redis

generate-sql:
	sqlc generate

test:
	go test ./... -v

start-alloy:
	export $(grep -v '^#' .env | xargs) && alloy run monitoring/config.alloy --server.http.listen-addr=0.0.0.0:12345

fetch-metrics:
	curl -u "alloy:alloy" http://localhost:8100/metrics

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
