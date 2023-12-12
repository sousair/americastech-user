# Americas Tech - User API

This API has been developed for the Golang Developer position at Americas Technology.

You can find the Exchange API built for the same test [here](https://github.com/sousair/americastech-exchange).

## Description

This repository consists of two servers:

### HTTP
This server hosts most of the functionalities, providing basic CRUD operations for users.
<br>
You can refer to the Postman documentation [here](https://).

### gRPC
Used exclusively for user token validation between APIs ([Exchange API]()).


## How to run

### Prerequisites

- Go (version 1.21.3 or later)
- Docker & Docker Compose

### Installing

1. Clone the repository:
  ``` bash
    git clone https://github.com/sousair/americastech-user.git
    cd americastech-user
  ```


2. Create a `.env` file and fill it with the credentials.

3. Build and run the application:

``` bash
  docker compose -f build/docker-compose.yml up --build
```

## Running the tests
--
