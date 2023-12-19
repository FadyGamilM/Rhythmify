# Rhythmify
Rhythmify is a microservice application built in Go to allow users to upload mp4 videos and download a mp3 sound version of the video

# Auth Microservice
1. Create k8s cluster using `Kind`
```bash
cluster-up:
	kind create cluster --image kindest/node:v1.21.1@sha256:69860bda5563ac81e3c0057d654b5253219618a22ec3a346306239bba8cfa1a6 --name $(KIND_CLUSTER) --config ./infra/kind-config.yaml
	kubectl create namespace $(CLUSTER_NAMESPACE)
	kubectl config set-context --current --namespace=$(CLUSTER_NAMESPACE)
```


