IMAGE_NAME := short_link_service

build:
	docker build . -t $(IMAGE_NAME)

run:
	cd deployments && docker-compose up -d

stop:
	cd deployments && docker-compose down -v --remove-orphans
