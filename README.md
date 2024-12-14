API Rate Limiter
Project Description
This project implements a Redis-based API rate limiter. It restricts the number of API requests a user can make within a specific time window (e.g., 1 minute). If the limit is exceeded, the server responds with a 429 Too Many Requests error.

The project is built using Go (Golang) with the Gin framework and uses Redis as the backend for tracking request counts. It is fully containerized with Docker Compose for seamless deployment.