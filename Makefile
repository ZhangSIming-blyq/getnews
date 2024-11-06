# Makefile for building Docker images and starting services

# Docker image names
FRONTEND_IMAGE=frontend-app
BACKEND_IMAGE=backend-app

# Build frontend and backend images
build:
	cd frontend && docker build -f deploy/Dockerfile -t $(FRONTEND_IMAGE) .
	cd backend && docker build -f deploy/Dockerfile -t $(BACKEND_IMAGE) .

# Run docker-compose to start services
run: build
	docker-compose up

# Clean up Docker images and containers
clean:
	docker-compose down
	docker rmi -f $(FRONTEND_IMAGE) $(BACKEND_IMAGE)

.PHONY: build run clean
