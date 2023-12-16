db-down-v:
	docker-compose down -v

db-down:
	docker volume rm auth_db_data
	docker-compose down

db-up:
	docker-compose up

create-auth-db:
	docker exec -it auth_db createdb --username=auth_db_user --owner=auth_db_user authdb


# ==========================
# running postgreSQL on k8s
pg-k8s-up:
	kind create cluster --name pgcluster --image 

# ==============
KIND_CLUSTER := rhythmify

CLUSTER_NAMESPACE := rhythmify-namespace

cluster-up:
	kind create cluster --image kindest/node:v1.21.1@sha256:69860bda5563ac81e3c0057d654b5253219618a22ec3a346306239bba8cfa1a6 --name $(KIND_CLUSTER) 
# --config ./infra/kind-config.yaml
	kubectl create namespace $(CLUSTER_NAMESPACE)
	kubectl config set-context --current --namespace=$(CLUSTER_NAMESPACE)

cluster-down:
	kind delete cluster --name $(KIND_CLUSTER)

create-depl:
# the -- separtes between the kubectl commands [before it] and shell command running within the container [after it]
	kubectl create deployment pingpong --image=alpine -- ping localhost

scale:
	kubectl scale deployment pingpong --replicas 3

watch-all-pods:
	watch kubectl get pods -o wide

load-img-to-cluster:
	kind load docker-image api:1.0 --name $(KIND_CLUSTER)

apply-to-cluster:
	kubectl apply -f ./infra/k8s/api-depl.yaml

build:
	docker build -t api:1.0 .