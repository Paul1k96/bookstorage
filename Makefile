run-app:
	docker network create skynet
	docker-compose build library
	docker-compose up library
run-app-detach:
	docker network create skynet
	docker-compose build library
	docker-compose up library -d
remove-app:
	docker-compose down -v
	docker rmi bookstorage-library
	docker network rm skynet
build:
	docker-compose build library
up:
	docker-compose up library
up-detach:
	docker-compose up -d
stop:
	docker-compose stop
down:
	docker-compose down -v
remove-api-img:
	docker rmi bookstorage-library
ps:
	docker-compose ps
create-net:
	docker network create skynet
delete-net:
	docker network rm skynet
