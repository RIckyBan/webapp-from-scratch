NAME = foo

db-migrate-up:
	cd ./db && sql-migrate up

db-migrate-down:
	cd ./db && sql-migrate down

db-migrate-new:
	cd ./db && sql-migrate new ${NAME}

db-import-seed:
	docker compose -f docker-compose.yml exec db bash -c "mysql -u user -ppassword webapp < /webapp/db/seeds/development.sql"

generate-models:
	cd ./db && sqlboiler mysql --output ./models

server:
	go run main.go
	
