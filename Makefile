include .env

migrate-up:
	migrate -path db/migrations -database ${DATABASE_URL} up

migrate-down-one:
	migrate -path db/migrations -database ${DATABASE_URL} down 1

migrate-down:
	migrate -path db/migrations -database ${DATABASE_URL} down 

migrate-force1:
	migrate -path db/migrations -database ${DATABASE_URL} force 1