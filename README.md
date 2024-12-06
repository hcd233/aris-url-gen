# Aris-url-gen

[ English | [简体中文](README_zh.md) ]

## Introduction

A high-performance URL shortening service developed in Go. The project name comes from the character Aris in Blue Archive, as shown in the figure below.

---
<p align="center">
  <img src="https://raw.githubusercontent.com/hcd233/Aris-AI/master/assets/110531412.jpg" width="50%">
  <br>Aris: A character from Blue Archive
</p>
---

## Features

- Generate short URLs: Converts long URLs into short links
- Customizable expiration time: Allows setting an expiration date for links
- Bidirectional caching: High-performance caching using Redis
- RESTful API: Provides a standard HTTP interface
- Data persistence: URL mappings are stored in MySQL

## Tech Stack

- **Web Framework**: [Fiber](https://github.com/gofiber/fiber)
- **ORM**: [GORM](https://gorm.io)
- **Caching**: Redis
- **Database**: MySQL
- **Logging**: [Zap](https://github.com/uber-go/zap)

## API Endpoints

### 1. Generate Short URL

```http
POST /v1/shortURL
Content-Type: application/json
{
    "originalURL": "https://example.com/very/long/url",
    "expireDays": 7  // Optional, expiration in days
}
```

### 2. Access Short URL

```http
GET /v1/s/{shortURL}
```

## Project Structure

```
.
├── cmd/                # Command line entry
├── internal/          
│   ├── api/           # API-related code
│   │   ├── dao/       # Data access layer
│   │   ├── dto/       # Data transfer objects
│   │   ├── handler/   # Request handlers
│   │   └── service/   # Business logic layer
│   ├── config/        # Configuration files
│   ├── logger/        # Logging
│   ├── resource/      # Resources
│   └── util/          # Utility functions
└── main.go            # Main entry point
```

## Installation and Deployment

### Prerequisites

- Go 1.20+
- MySQL 8.0+
- Redis 6.0+

### Local Development

1. Clone the repository

```bash
git clone https://github.com/hcd233/Aris-url-gen.git
cd Aris-url-gen
```

2. Install dependencies

```bash
go mod download
```

3. Configure environment variables

Refer to `api.env.template` to set up the required environment variables

4. Run the service

```bash
go run main.go
```

## Deployment

### Docker Deployment

1. Create required volumes:

```bash
docker volume create mysql-data
docker volume create redis-data
```

2. Deploy with docker-compose:

```bash
# Development environment
docker compose -f docker/debug-docker-compose.yml up -d

# Production environment
docker compose -f docker/docker-compose.yml up -d
```

### Kubernetes Deployment

1. Create namespace and configs:

```bash
kubectl apply -f kubernetes/namespace.yml
kubectl apply -f kubernetes/configmaps.yml
```

2. Create secrets (configure secrets.yml first):

```bash
cp kubernetes/secrets.yml.template kubernetes/secrets.yml
# Edit secrets.yml with actual secret values
kubectl apply -f kubernetes/secrets.yml
```

3. Create storage and deploy services:

```bash
kubectl apply -f kubernetes/pvc.yml
kubectl apply -f kubernetes/deployments.yml
kubectl apply -f kubernetes/services.yml
```

### Helm Deployment

1. Configure values.yaml:

```bash
cp helm/aris-url-gen/values.yaml.template helm/aris-url-gen/values.yaml
# Edit values.yaml with your configuration
```

2. Install with Helm:

```bash
helm install aris-url-gen helm/aris-url-gen
```

3. Upgrade or uninstall:

```bash
# Upgrade
helm upgrade aris-url-gen helm/aris-url-gen

# Uninstall
helm uninstall aris-url-gen
```

## License

This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for more details.