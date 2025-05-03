docker:
	@echo "Docker Build..."
# docker build . -t k8s-combat:v1
	docker build . -t haruhi07734/k8s-combat:configmap && docker image push haruhi07734/k8s-combat:configmap