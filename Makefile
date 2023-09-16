build:
	docker-compose build library
up:
	docker-compose up library
stop:
	docker-compose stop
down:
	docker-compose down -v
remove-api-img:
	docker rmi bookstorage-library
ps:
	docker-compose ps
