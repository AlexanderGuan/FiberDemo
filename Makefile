.PHONY: clean critic security lint test build run

APP_NAME = apiserver
BUILD_DIR = $(PWD)/build
MIGRATIONS_FOLDER = $(PWD)/platform/migrations
DATABASE_URL = postgres://postgres:password@cgapp-postgres/postgres?sslmode=disable

docker.run: docker.network docker.postgres docker.fiber docker.redis migrate.up

migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" up

migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down

docker.network:
	docker network inspect dev-network > /dev/null 2>&1 || \
	docker network create --driver bridge dev-network

docker.postgres:
	docker run --rm -d \
	--name cgapp-postgres \
	--network dev-network \
	-e POSTGRES_USER=postgres \
	-e POSTGRES_PASSWORD=postgres \
	-e POSTGRES_DB=postgres \
	-v ${HOME}/dev-postgres/data/:/var/lib/postgresql/data \
	-p 5432:5432 \
	postgres

docker.redis:
	docker run --rm -d \
		--name cgapp-redis \
		--network dev-network \
		-p 6379:6379 \
		redis

docker.fiber.build:
	docker build -t fiber .

docker.fiber: docker.fiber.build
	docker run --rm -d \
		--name cgapp-fiber \
		--network dev-network \
		-p 5000:5000 \
		fiber