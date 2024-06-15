# Simple Go Redis Server

This project is a simple Redis-like server implemented in Go. It supports basic Redis commands like `PING`, `SET`, `GET`, `HSET`, and `HGET`. The server also supports Append-Only File (AOF) persistence to ensure data durability.

## Features

- **PING**: Check the server's status.
- **SET**: Set a string key to a string value.
- **GET**: Get the value of a string key.
- **HSET**: Set a field in a hash.
- **HGET**: Get the value of a field in a hash.
- **AOF Persistence**: Commands that modify data (`SET` and `HSET`) are logged to an AOF file to ensure data durability.

## Getting Started

### Prerequisites

- Docker
- Docker Compose

### Running the Server

1. **Build and run the server using Docker Compose**:

    ```sh
    docker-compose up --build
    ```

2. **Interact with the server using a Redis client**:

    You can use any Redis client to connect to the server running on `localhost:6379`. For example, using `redis-cli`:

    ```sh
    redis-cli -p 6379
    ```
    or utilize the existing docker compose setup like:
    ```sh
    docker compose exec redis-client sh
    redis-cli -h redis-server -p 6379
    ```
    after which you can start sending commands to the server.

    Example commands:

    ```sh
    SET key1 value1
    GET key1
    HSET hash1 field1 value1
    HGET hash1 field1
    ```

### Volumes

The AOF file is stored in a Docker volume to persist data across container restarts. The file is mapped to the `./app` directory in the container.
