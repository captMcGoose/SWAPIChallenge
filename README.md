# mercedesChallenge

## Overview

This project implements a microservice in Go that interacts with the Star Wars API (SWAPI). The service fetches data from the `people` endpoint of the API, sorts it by name in ascending order, and returns the sorted data through a custom API endpoint. The service is containerized using Docker, orchestrated with Docker Compose, and deployed to a local Kubernetes cluster using Minikube.

This repo contains the solution to the technical challenge provided. 
It's a microservice in Go that fetches the "people" endpoint from the Star Wars API, sorts the data, and returns it via a custom endpoint.

As requested, the service is containerized using Docker, orchestrated with docker-compose, and deployed to a local Minikube k8s cluster.

## project Structure

- `main.go`: Main microservice script in go
- `Dockerfile`: Docker configuration file to containerize the microservice.
- `docker-compose.yml`: docker-compose configuration to run the service locally.
- `deployment.yaml`: Kubernetes Deployment configuration.
- `service.yaml`: Kubernetes Service configuration.
- `autoscaler.yaml`: Kubernetes Horizontal Pod Autoscaler configuration.
- `load_test.sh`: Script to load test the service.

## Requirements

- Go 1.20 or later
- Docker
- Docker Compose
- Minikube (or another Kubernetes local cluster solution)

## Setup Instructions

### Running Locally

1. **Clone the repository**:
   ```bash
   git clone https://github.com/captMcGoose/mercedesChallenge.git
   cd mercedesChallenge
   ```

2. **Run the application**:
   ```bash
    go run main.go
   ```

3. **Access the API**:

http://localhost:8080/people

### Running with Docker

1. **Build the Docker image**:
   ```bash
    docker build -t mercedeschallenge .
   ```

2. **Run the Docker container**:
   ```bash
    docker run -p 8080:8080 mercedeschallenge
   ```

3. **Access the API**:

http://localhost:8080/people

### Running with Docker Compose

1. **Start the service**:
   ```bash
    docker-compose up --build
   ```
2. **Access the API**:

http://localhost:8080/people


### Deploying to Kubernetes

1. **Start Minikube**:

   ```bash
    minikube start
   ```    
2. **Deploy the application**:
   ```bash
    kubectl apply -f deployment.yaml
    kubectl apply -f service.yaml
   ```
3. **Access the service**:
   ```bash
    minikube service mercedeschallenge-service
   ```

### Horizontal Pod Autoscaler

Deploy the autoscaler:
   ```bash
    kubectl apply -f autoscaler.yaml
   ```

## Load Testing

The script simulates concurrent traffic to the endpoint. It runs multiple HTTP requests in parallel and helps evaluate how well the service handles a heavy load. 
It can help identify bottlenecks, monitor response times under load, and ensure that the service scales as expected.

1. **Start the cluster**:

```bash
    minikube start
```

2. **Deploy the microservice**:

```bash
   kubectl apply -f deployment.yaml
   kubectl apply -f service.yaml
```
you can aditionally deploy the autoscaler:

```bash
    kubectl apply -f autoscaler.yaml
```
2. **Run the load test script**:

   ```bash
    ./load_test.sh
   ```
3. **Monitor the performance**:
Use the following commands:
   ```bash
   kubectl get pods
   kubectl logs <pod-name>
   kubectl top pods
   kubectl get hpa
   ```
