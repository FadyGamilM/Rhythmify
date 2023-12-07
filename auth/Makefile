db-down-v:
	docker-compose down -v

db-down:
	docker volume rm auth_db_data
	docker-compose down

db-up:
	docker-compose up

create-auth-db:
	docker exec -it auth_db createdb --username=auth_db_user --owner=auth_db_user authdb

