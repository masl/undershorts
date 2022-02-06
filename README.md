[![build](https://github.com/masl/undershorts/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/masl/undershorts/actions/workflows/build.yml)
[![docker](https://github.com/masl/undershorts/actions/workflows/docker.yml/badge.svg)](https://github.com/masl/undershorts/actions/workflows/docker.yml)

# undershorts
Brevity is the soul of wit - Undershorts is a simple URL shortener

## Deployment 
A containerized solution for undershorts with redis may look like this
```yaml
version: "3"
services:
  undershorts_redis:
    image: "redis"
    container_name: "undershorts_redis"
    command: sh -c "redis-server --requirepass PASSWORD --appendonly yes"
    volumes:
      - "./redis/db/:/data/"
    networks:
      - "undernet"

  undershorts:
    image: "ghcr.io/masl/undershorts:latest"
    container_name: "undershorts"
    environment:
      UNDERSHORTS_REDIS_URL: "redis://:PASSWORD@undershorts_redis:6379"
      UNDERSHORTS_WEB_ADDRESS: "0.0.0.0:8080"
    networks:
      - "undernet"
    depends_on:
      - "undershorts_redis"

networks:
  undernet:
```
## Environment Variables
| Environment Variable      | Default Value                              |
|---------------------------|--------------------------------------------|
| `UNDERSHORTS_REDIS_URL`   | `redis://:PASSWORD@undershorts_redis:6379` |
| `UNDERSHORTS_WEB_ADDRESS` | `0.0.0.0:8000`                             |
