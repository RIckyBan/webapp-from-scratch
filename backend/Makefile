migrate-up:
	cd ./db && sql-migrate up

migrate-down:
	cd ./db && sql-migrate down

generate-models:
	sqlboiler mysql --output ./db/models

server:
	go run main.go
	