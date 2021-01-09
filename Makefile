DOCKER_REGISTRY ?= "nhomble93"

.PHONY: build
windows-build:
	go build -o notify.exe ./main.go

.PHONY: docker-build
docker-build:
	mkdir -p fs
	GOOS=linux GOARCH=amd64 go build -o fs/notify ./main.go
	docker build -t $(DOCKER_REGISTRY)/groupme-notify:latest .

.PHONY: docker-push
docker-push:
	docker push $(DOCKER_REGISTRY)/groupme-notify