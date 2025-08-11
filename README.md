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
- `GET /fetch` - Fetches data from an external API with API key authentication

**Environment Variables:**
- `PORT` - Server port (default: 8080)
- `ENVIRONMENT` - Environment name (default: development)
- `VERSION` - Application version (default: 1.0.0)
- `API_KEY` - API key for external API authentication (default: dummy-secret-key-12345)

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

4. **Test the fetch endpoint** with API key:
   ```bash
   # Set a custom API key
   export API_KEY="your-secret-api-key"
   
   # Test the fetch endpoint
   curl http://localhost:8080/fetch
   
   # Or test with the default API key
   curl http://localhost:8080/fetch
   ```

## 🔐 API Key Configuration

The application uses an `API_KEY` environment variable to authenticate with external APIs. The fetch endpoint (`/fetch`) automatically includes this API key in the `X-API-Key` header when making requests to external services.

- **Default value**: `dummy-secret-key-12345` (for development/testing)
- **Production use**: Set `API_KEY` to your actual API key
- **Security**: Never commit real API keys to version control
