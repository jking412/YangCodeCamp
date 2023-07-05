
IMAGE_NAME=yangcodecamp-web

all:
	docker compose up -d

clean:
	docker compose down
	docker rmi $(IMAGE_NAME)

