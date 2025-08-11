# 🚀 DevOps Challenge: Go Application Deployment

Welcome to the DevOps challenge! This repository contains a simple Go web application that you need to containerize, test, and deploy using modern DevOps practices.

## 📋 Challenge Overview

Your task is to implement a complete DevOps pipeline for this Go application. You'll be evaluated on your ability to:

1. **Containerize** the application with Docker
2. **Implement CI/CD** using GitHub Actions
3. **Deploy** to Kubernetes

## 🎯 What You Need to Create

### 1. Dockerfile
Create a production-ready Dockerfile

### 2. GitHub Actions Workflow
Create a comprehensive CI/CD pipeline.

### 3. Kubernetes Manifests
Create the K8s resources

## 🏗️ Application Details

This is a simple Go web service with the following endpoints:

- `GET /` - Homepage with application info
- `GET /health` - Health check endpoint (returns 204 No Content)
- `GET /info` - Application information
- `GET /nytime` - Current New York time

**Environment Variables:**
- `PORT` - Server port (default: 8080)
- `ENVIRONMENT` - Environment name (default: development)
- `VERSION` - Application version (default: 1.0.0)

## 🚀 Getting Started

1. **Clone and explore** the repository
2. **Run locally** to understand the application:
   ```bash
   go run main.go
   curl -v http://localhost:8080/health
   curl http://localhost:8080/nytime
   ```

3. **Run tests** to ensure everything works:
   ```bash
   go test -v
   go test -v -cover
   ```
