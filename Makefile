# Variables
ENV ?= staging
VERSION ?= 0.1.0
SPEAKBUDDYBEAPI_NAMESPACE = speakbuddybe-ns
MYSQL_NAMESPACE = db-ns
MANIFESTS_DIR = manifests

DOCKER_IMAGE_NAME = speakbuddybeapik8s
DOCKER_IMAGE_TAG = $(VERSION)-$(ENV)
DOCKERHUB_USERNAME = tandrysyawaludin

# DEPLOY MYSQL

## MySQL setup
mysql-setup:
	kubectl create -f $(MANIFESTS_DIR)/mysql/$(ENV)/ns-mysql.yaml
	kubectl create -f $(MANIFESTS_DIR)/mysql/$(ENV)/pvc-mysql.yaml
	kubectl create -f $(MANIFESTS_DIR)/mysql/$(ENV)/secret-mysql.yaml
	kubectl create -f $(MANIFESTS_DIR)/mysql/$(ENV)/service-mysql.yaml
	kubectl create -f $(MANIFESTS_DIR)/mysql/$(ENV)/deployment-mysql.yaml

## Open database
mysql-open:
	kubectl exec -it -n $(MYSQL_NAMESPACE) $$(kubectl get pods -n $(MYSQL_NAMESPACE) -l app=mysql -o jsonpath="{.items[0].metadata.name}") -- bash

# DEPLOY APP
deploy-app: docker-setup speakbuddybeapi-setup

## Docker build
docker-setup: docker-build docker-tag docker-login docker-push
docker-build:
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) .
docker-tag:
	docker tag $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) $(DOCKERHUB_USERNAME)/$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)
	docker images | grep $(DOCKER_IMAGE_NAME)
docker-login:
	docker login
docker-push:
	docker push $(DOCKERHUB_USERNAME)/$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)

## SpeakbuddybeAPI setup
speakbuddybeapi-setup:
	kubectl create -f $(MANIFESTS_DIR)/speakbuddybeapi/$(ENV)/ns-speakbuddybeapi.yaml
	kubectl create -f $(MANIFESTS_DIR)/speakbuddybeapi/$(ENV)/secret-speakbuddybeapi.yaml
	kubectl create -f $(MANIFESTS_DIR)/speakbuddybeapi/$(ENV)/service-speakbuddybeapi.yaml
	kubectl create -f $(MANIFESTS_DIR)/speakbuddybeapi/$(ENV)/deployment-speakbuddybeapi.yaml
	kubectl create -f $(MANIFESTS_DIR)/speakbuddybeapi/$(ENV)/configmap-speakbuddybeapi.yaml 


## SERVE APPLICATION
serve:
	kubectl port-forward -n $(SPEAKBUDDYBEAPI_NAMESPACE) svc/speakbuddybeapi 8081

# CLEAN UP RESOURCES
clean-all: clean-db clean-app
clean-db:
	kubectl delete ns $(MYSQL_NAMESPACE) || true
clean-app:
	kubectl delete ns $(SPEAKBUDDYBEAPI_NAMESPACE) || true
