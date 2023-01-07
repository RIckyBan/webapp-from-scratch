NAME = foo

migrate-up:
	cd ./db && sql-migrate up

migrate-down:
	cd ./db && sql-migrate down

migrate-new:
	cd ./db && sql-migrate new ${NAME}

import-seed:
	docker compose -f docker-compose.yml exec db bash -c "mysql -u user -ppassword webapp < /webapp/db/seeds/development.sql"

generate-models:
	cd ./db && sqlboiler mysql --output ./models

server:
	go run main.go
	
