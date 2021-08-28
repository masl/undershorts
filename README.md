[![build](https://github.com/masl/undershorts/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/masl/undershorts/actions/workflows/build.yml)
[![docker](https://github.com/masl/undershorts/actions/workflows/docker.yml/badge.svg)](https://github.com/masl/undershorts/actions/workflows/docker.yml)

# undershorts
Brevity is the soul of wit - Undershorts is a simple URL shortener

## Deployment
My recommended usage of undershorts with docker-compose may look like this 
```yaml
version: "3"

services:
  undershorts:
    image: "ghcr.io/masl/undershorts:latest"
    container_name: "undershorts"
    ports:
      - "8000:8000"
    volumes:
      - "./data:/app"
```
