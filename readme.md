# Go Weather Microservice

A lightweight **Go-based weather microservice** with both **CLI** and **REST API** modes — fully containerized and deployable on **Kubernetes**, with **Terraform** and **GitHub Actions** support.

## Features

- Developed in **Go** — dual mode:
  - CLI for quick forecasts
  - REST API server (`/weather?city=Prague`)
- **Dockerized** for consistent builds
- **Kubernetes manifests** for local orchestration (via `kind`)
- **Terraform IaC** (Kubernetes provider)
- **GitHub Actions CI/CD** (build, test, lint, deploy)
- Optional **monitoring** (Prometheus & Grafana)
- 100% **cloud-native**, works fully **locally** and offline

## Structure

cmd/         – CLI entry point  
internal/    – Core logic (API client, models, server)  
deploy/      – Kubernetes manifests (Deployment, Service, Secret)  
Dockerfile   – Multi-stage build  


## Run Locally

```bash
# CLI mode
go run . --city=Prague

# Server mode
go run . --mode=server
curl http://localhost:8080/weather?city=Prague

Docker

docker build -t go-weather .
docker run -p 8080:8080 go-weather --mode=server

Kubernetes (kind)

kubectl apply -f deploy/secret.yaml
kubectl apply -f deploy/deployment.yaml
kubectl apply -f deploy/service.yaml
kubectl port-forward service/go-weather-service 8080:8080
kubectl get pods

Access at: http://localhost:<NodePort>/weather?city=Prague