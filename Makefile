docker:
	@echo "Docker Build..."
	docker build . -t k8s-combat:v1
# docker build . -t yzh/k8s-combat:v1
# && docker image push yzh/k8s-combat:v1