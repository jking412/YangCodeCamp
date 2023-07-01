
IMAGE_NAME = "yangcodecamp-web"

all:
	docker build -t $(IMAGE_NAME) .

clean:
	docker rmi $(IMAGE_NAME)

