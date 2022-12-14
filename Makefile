TAG := ""

ECR = ""
TARGET_REPOSITORY := ""


docker-build:
	docker build -f Dockerfile --tag ${TAG} .
push:
	docker push


images-push: docker-build push

