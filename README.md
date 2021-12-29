"# golang_rest_redis_rabbitmq" 

# Golang Rest Api tutorial with Redis and RabbitMQ using 


# Install With Docker
- docker-compose build
- docker-compose up -d
- docker-compose ps
- docker-compose logs -f app

# Before checking api , check table orders exists in database
- if not manually import sql file from ./database/init.sql

# Run without docker
- Edit .env in file following lines:
- DB_HOST=127.0.0.1
- REDIS_HOST=localhost:6379
- RABBITMQ_HOST=localhost:5672

