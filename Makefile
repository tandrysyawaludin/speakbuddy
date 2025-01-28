# VARIABLES
ENV ?= staging
VERSION ?= 0.1.0
MANIFESTS_DIR = manifests

DOCKER_IMAGE_NAME = speakbuddybeapik8s
DOCKERHUB_USERNAME = tandrysyawaludin
DOCKER_IMAGE_TAG = $(VERSION)-$(ENV)
DOCKER_IMAGE = $(DOCKERHUB_USERNAME)/$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)

SPEAKBUDDYBE_NAMESPACE = speakbuddybe-ns
SPEAKBUDDYBE_DEPLOYMENT_NAME ?= speakbuddybeapi
SPEAKBUDDYBEAPI_CONFIGMAP_NAME ?= speakbuddybe-cm
SPEAKBUDDYBE_DB_NAMESPACE = speakbuddybe-db-ns
SPEAKBUDDYBE_DB_SECRET_NAME ?= speakbuddybe-db-password
SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH = $(MANIFESTS_DIR)/speakbuddybeapi/$(ENV)/deployment-speakbuddybeapi.yaml

# DEPLOY MYSQL

## mysql setup
mysql-setup:
	kubectl create -f $(MANIFESTS_DIR)/mysql/$(ENV)/ns-mysql.yaml
	kubectl create -f $(MANIFESTS_DIR)/mysql/$(ENV)/pvc-mysql.yaml
	kubectl create -f $(MANIFESTS_DIR)/mysql/$(ENV)/secret-mysql.yaml
	kubectl create -f $(MANIFESTS_DIR)/mysql/$(ENV)/service-mysql.yaml
	kubectl create -f $(MANIFESTS_DIR)/mysql/$(ENV)/deployment-mysql.yaml

## open database
mysql-open:
	kubectl exec -it -n $(SPEAKBUDDYBE_DB_NAMESPACE) $$(kubectl get pods -n $(SPEAKBUDDYBE_DB_NAMESPACE) -l app=mysql -o jsonpath="{.items[0].metadata.name}") -- bash

# DEPLOY APP
deploy-app: docker-setup speakbuddybeapi-setup

## docker build
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

## speakbuddybeAPI setup
speakbuddybeapi-setup: generate-speakbuddybeapi-deployment-file speakbuddybeapi-create-ns
speakbuddybeapi-create-ns:	
	kubectl create -f $(MANIFESTS_DIR)/speakbuddybeapi/$(ENV)/ns-speakbuddybeapi.yaml
	kubectl create -f $(MANIFESTS_DIR)/speakbuddybeapi/$(ENV)/secret-speakbuddybeapi.yaml
	kubectl create -f $(MANIFESTS_DIR)/speakbuddybeapi/$(ENV)/service-speakbuddybeapi.yaml
	kubectl create -f $(MANIFESTS_DIR)/speakbuddybeapi/$(ENV)/deployment-speakbuddybeapi.yaml
	kubectl create -f $(MANIFESTS_DIR)/speakbuddybeapi/$(ENV)/configmap-speakbuddybeapi.yaml

## SERVE APPLICATION
serve:
	kubectl port-forward -n $(SPEAKBUDDYBE_NAMESPACE) svc/speakbuddybeapi 8081

# CLEAN UP RESOURCES
clean-all: clean-db clean-app clean-docker
clean-db:
	kubectl delete ns $(SPEAKBUDDYBE_DB_NAMESPACE) || true
clean-app:
	kubectl delete ns $(SPEAKBUDDYBE_NAMESPACE) || true
clean-docker:
	docker system prune -a

# GENERATE FILE
## generate deployement file
generate-speakbuddybeapi-deployment-file:
	@echo "Generating $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)..."
	@echo "apiVersion: apps/v1" > $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "kind: Deployment" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "metadata:" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "  creationTimestamp: null" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "  labels:" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "    app: $(SPEAKBUDDYBE_DEPLOYMENT_NAME)" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "  name: $(SPEAKBUDDYBE_DEPLOYMENT_NAME)" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "  namespace: $(SPEAKBUDDYBE_NAMESPACE)" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "spec:" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "  replicas: 1" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "  selector:" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "    matchLabels:" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "      app: $(SPEAKBUDDYBE_DEPLOYMENT_NAME)" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "  strategy: {}" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "  template:" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "    metadata:" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "      creationTimestamp: null" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "      labels:" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "        app: $(SPEAKBUDDYBE_DEPLOYMENT_NAME)" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "    spec:" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "      containers:" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "      - image: $(DOCKER_IMAGE)" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "        name: $(DOCKER_IMAGE_NAME)" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "        resources: {}" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "        env:" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "          - name: CONFIG_DBPASS" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "            valueFrom:" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "              secretKeyRef:" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "                key: rootpassword" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "                name: $(SPEAKBUDDYBE_DB_SECRET_NAME)" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "          - name: CONFIG_DBNAME" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "            valueFrom:" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "              configMapKeyRef:" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "                key: dbname" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "                name: $(SPEAKBUDDYBEAPI_CONFIGMAP_NAME)" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "          - name: CONFIG_DBUSER" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "            valueFrom:" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "              configMapKeyRef:" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "                key: dbuser" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "                name: $(SPEAKBUDDYBEAPI_CONFIGMAP_NAME)" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "          - name: CONFIG_SERVER_PORT" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "            valueFrom:" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "              configMapKeyRef:" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "                key: serverport" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "                name: $(SPEAKBUDDYBEAPI_CONFIGMAP_NAME)" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "          - name: CONFIG_DBHOST" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "            valueFrom:" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "              configMapKeyRef:" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "                key: dbhost" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "                name: $(SPEAKBUDDYBEAPI_CONFIGMAP_NAME)" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "status: {}" >> $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH)
	@echo "Generated $(SPEAKBUDDYBE_DEPLOYMENT_FILE_PATH) successfully!"
