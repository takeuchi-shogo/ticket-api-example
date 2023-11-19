
buildup:
	@echo "Building docker image."
	docker-compose --env-file ./server/.env up --build
	@echo "Build docker image success."

up:
	@echo "Running docker image."
	docker-compose --env-file ./server/.env up
	@echo "Running docker image success."

rm:
	@echo "Removing docker image."
	docker-compose --env-file ./server/.env rm
	@echo "Removing docker image success."
