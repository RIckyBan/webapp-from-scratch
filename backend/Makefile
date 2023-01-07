migrate-up:
	cd ./db && sql-migrate up

migrate-down:
	cd ./db && sql-migrate down

generate-models:
	cd ./db && sqlboiler mysql --output ./models

server:
	go run main.go
	
