[![build](https://github.com/masl/undershorts/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/masl/undershorts/actions/workflows/build.yml)
[![docker](https://github.com/masl/undershorts/actions/workflows/docker.yml/badge.svg)](https://github.com/masl/undershorts/actions/workflows/docker.yml)

# undershorts
Brevity is the soul of wit - Undershorts is a simple URL shortener

## Deployment 
A containerized solution for undershorts with redis may look like this
```yaml
version: "3.9"
services:
  undershorts_redis:
    image: "redis"
    container_name: "undershorts_redis"
    ports:
      - "6379:6379"
      - "8000:8000"
  
  undershorts:
    image: "ghcr.io/masl/undershorts:latest"
    container_name: "undershorts"
    environment:
      UNDERSHORTS_REDIS_ADDRESS: "127.0.0.1:6379"
      UNDERSHORTS_REDIS_PASSWORD: "YOUR_REDIS_PASSWORD"
      UNDERSHORTS_WEB_ADDRESS: "0.0.0.0.8000"
    network_mode: "container:undershorts_redis"
    depends_on:
      - "undershorts_redis"
```
## Environment Variables
| Environment Variable         | Default Value    |
|------------------------------|------------------|
| `UNDERSHORTS_WEB_ADDRESS`    | `0.0.0.0:8000`   |
| `UNDERSHORTS_REDIS_ADDRESS`  | `127.0.0.1:6379` |
| `UNDERSHORTS_REDIS_PASSWORD` | ` `              |
