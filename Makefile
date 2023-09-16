build:
	docker-compose build library
up:
	docker-compose up library
down:
	docker-compose down -v
remove-api-img:
	docker rmi bookstorage-library

