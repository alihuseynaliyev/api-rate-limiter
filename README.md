# API Rate Limiter

## Project Description
This project implements a Redis-based API rate limiter. It restricts the number of API requests a user can make within a specific time window (e.g., 1 minute). If the limit is exceeded, the server responds with a `429 Too Many Requests` error.

The project is built using **Go** (Golang) with the **Gin framework** and uses **Redis** as the backend for tracking request counts. It is fully containerized with **Docker Compose** for seamless deployment.

---

## Features
- Redis-based rate limiting
- Limits requests based on client IP
- Scalable and easy-to-extend architecture
- Full Docker Compose support
- Healthcheck integration for Redis

---

## Technologies Used
- **Programming Language:** Go (Golang)
- **Framework:** Gin (Web Framework)
- **Database:** Redis
- **Containerization:** Docker Compose
- **Dependencies:**
    - `gin-gonic/gin`: Web framework
    - `go-redis/redis/v8`: Redis client

---

## Getting Started

### Prerequisites
Before starting, ensure you have the following installed:
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Git](https://git-scm.com/)

### Clone the Repository
Clone this repository to your local machine:
```bash
git clone https://github.com/alihuseynaliyev/api-rate-limiter.git
