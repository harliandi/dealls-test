# SOFTWARE ENGINEER TEST

## Tech Stack

- Go version 1.22.x
- MySQL version 8.0
- Docker version 19.x
- Redis version 6.2.2

## Golang Modules

- [chi](https://github.com/go-chi/chi) for HTTP router
- [sqlx](https://github.com/jmoiron/sqlx) for SQL client
- [go-redis](https://github.com/redis/go-redis) for Redis Client

## How to run

1. Clone the repository
2. Update DSN configuration for MySQL
3. Import the MySQL file to your database
4. Install all dependencies using command `go mod download`
5. Using terminal run command `go run main.go` to start the API

## How to run MySQL and Redis using Docker Compose

1. Install Docker from this [link](https://www.docker.com/get-started/) to your local machine
2. After install run this command to start mysql and redis docker `docker compose up -d`

## How to Deploy using Docker

1. Build the docker image using this command `docker build  -t your-repo/api-service:latest .`
2. Change "your-repo" to name of your repository
3. Tag the image to your repository eg: docker hub using command `docker tag your-repo/api-service:latest`
4. Push the image to image registry using command `docker push your-repo/api-service:latest`
5. Run the image on your server using command `docker run -it --rm -p 3001:3001 your-repo/api-service:latest`

## File Structure

- main.go
- model
  - m_user.go
  - m_user_test.go
  - m_auth.go
  - m_auth_test.go
- request/
  - auth.go
- utils/
  - http.go
  - validation.go
- go.mod
- go.sum
- dealls-db.sql
- Dealls.postman_collection.json
- Sequence.txt

## Tools

- VSCode
- Golangci-lint for Linter
