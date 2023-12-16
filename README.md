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



okay here is my yaml files after updates : 
- the depl.yaml for my golang api : 

# The namesapce 
apiVersion: v1
kind: Namespace
metadata:
  name: rhythmify-namespace
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: auth
  namespace: rhythmify-namespace
  labels:
    app: auth
spec:
  replicas: 3
  selector:
    matchLabels:
      app: auth
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 3
  template:
    metadata:
      labels:
        app: auth
    spec:
      containers:
      - name: auth
        image: fadygamil/auth:latest
        envFrom:
        - configMapRef:
            name: auth-configmap
        - secretRef:
            name: auth-secret
        ports:
        - containerPort: 3005

- the service for my golang api : 
apiVersion: v1
kind: Service
metadata:
  name: auth
  namespace: rhythmify-namespace
spec:
  selector:
    app: auth
  type: ClusterIP
  ports:
  - port: 5003
    targetPort: 5003
    protocol: TCP


- the secret.yaml and config.yaml 
# apiVersion: v1
# kind: Secret
# metadata:
#   name: auth-secret
#   namespace: rhythmify-namespace
# type: Opaque
# data:
#   POSTGRES_PASSWORD: "auth123"
#   JWT_SECRET_KEY: "just_test_auth_service_in_microservice_app"

apiVersion: v1
kind: Secret
metadata:
  name: auth-secret
  namespace: rhythmify-namespace
type: Opaque
data:
  POSTGRES_PASSWORD: "YWRtaW4xMjM=" # "auth123" base64 encoded
  JWT_SECRET_KEY: "amVtZW50dXJlcjEtdGVzdF9hdXRob3JpemVkLWNhbGxvY2F0aW9uLWF0dGVudGlzLmFjdGl2ZS5hcHBz" # "just_test_auth_service_in_microservice_app" base64 encoded


apiVersion: v1
kind: ConfigMap
metadata:
  name: auth-configmap
  namespace: rhythmify-namespace
data:
  SERVER_PORT: "3005"
  POSTGRES_HOST: "postgres-service:5432"
  POSTGRES_USER: "auth_user"
  POSTGRES_SSLMODE: "disable"
  POSTGRES_DB: "auth"



- the postgres.yaml and postgres-srv.yaml:
apiVersion: apps/v1
kind: Deployment
metadata:
  name: postgres-deployment
spec:
  replicas: 1
  selector:
    matchLabels:
      app: postgres
  template:
    metadata:
      labels:
        app: postgres
    spec:
      containers:
      - name: postgres
        image: postgres:14
        env:
        - name: POSTGRES_USER
          value: auth_user
        - name: POSTGRES_PASSWORD
          value: auth123
        - name: POSTGRES_DB
          value: auth
        - name: POSTGRES_SSLMODE
          value: disable
        ports:
        - containerPort: 5432

apiVersion: v1
kind: Service
metadata:
  name: postgres-service
spec:
  selector:
    app: postgres
  ports:
  - protocol: TCP
    port: 5432
    targetPort: 5432
  type: ClusterIP

and this is how i create a new cluster 
KIND_CLUSTER := rhythmify

CLUSTER_NAMESPACE := rhythmify-namespace

cluster-up:
	kind create cluster --image kindest/node:v1.21.1@sha256:69860bda5563ac81e3c0057d654b5253219618a22ec3a346306239bba8cfa1a6 --name $(KIND_CLUSTER) 
# --config ./infra/kind-config.yaml
	kubectl create namespace $(CLUSTER_NAMESPACE)
	kubectl config set-context --current --namespace=$(CLUSTER_NAMESPACE)

is everything is okay ? 
should i run make cluster-up and then apply the yaml files using kubectl ? or i  will have errors ?
ANSWER
