<p align="center">
  <picture>
  <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/dkrasnovdev/siberiana-public-assets/main/assets/siberiana-logo-dark-background.svg">
  <img src="https://raw.githubusercontent.com/dkrasnovdev/siberiana-public-assets/main/assets/siberiana-logo-dark-background.svg" width="240" height="90" alt="Logo for Siberiana">
</picture>
</p>

<p align="center">
Siberiana | Aggregator of the Historical and Cultural Heritage of the Yenisei Siberia.
</p>

# Siberiana GraphQL API

Siberiana GraphQL API serves as an essential gateway to access and interact with the vast historical and cultural heritage of the Yenisei Siberia region. By providing a powerful GraphQL interface, this API offers a comprehensive way to query and retrieve information about the rich historical artifacts, cultural landmarks, and heritage sites of the Yenisei Siberia area.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/) (at least version 1.19)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

You will also need to set up the Task CLI tool and Air. Here's how:

1. **Install Task**: Task is a task runner that simplifies command execution. If you haven't already, install Task by following the instructions at [https://taskfile.dev/installation/](https://taskfile.dev/installation/).

2. **Install Air**: Air is used for automatically rebuilding and restarting the application when code changes occur. There are multiple methods to install Air:

   a. **Install via Binary** (preferred):

   ```bash
   # Install Air binary into $(go env GOPATH)/bin/
   curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

   # Alternatively, install it into ./bin/
   # curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
   ```

   b. **Install via Go Install**:

   ```bash
   go install github.com/cosmtrek/air@latest
   ```

   c. **Using Docker**:

   ```bash
   docker run -it --rm \
       -w "<PROJECT>" \
       -e "air_wd=<PROJECT>" \
       -v $(pwd):<PROJECT> \
       -p <PORT>:<APP SERVER PORT> \
       cosmtrek/air
       -c <CONF>
   ```

   For more information on installation and usage, visit [https://github.com/cosmtrek/air](https://github.com/cosmtrek/air).

Ensure you have running Keycloak and MinIO instances. If you don't have them set up, use the following repositories to set them up:

- **Keycloak Instance**: Follow the instructions at [https://github.com/dkrasnovdev/siberiana-keycloak](https://github.com/dkrasnovdev/siberiana-keycloak) to set up a running Keycloak instance. Keycloak provides authentication and authorization services for the Siberiana GraphQL API.

- **MinIO Instance**: Set up a running MinIO instance by following the instructions at [https://github.com/dkrasnovdev/siberiana-minio](https://github.com/dkrasnovdev/siberiana-minio). MinIO is used for object storage and is an essential part of the Siberiana GraphQL API.

Follow the steps outlined in the next section to start containers, run tasks, and access the application.

## Getting Started

1. Clone this repository to your local machine:

```bash
git clone https://github.com/dkrasnovdev/siberiana-api.git
cd siberiana-api
```

2. Run the following command to download the required packages specified in the Taskfile:

```bash
task install
```

3. Copy the `.env.example` file and rename it to `.env`. Fill in the required environment variables with appropriate values.

```bash
cp .env.example .env
# Edit .env file with your configuration
```

## For Local Development

To work on the project locally, follow these steps:

1. Start the required containers for local development:

```bash
docker-compose up -d siberiana_postgres siberiana_redis
```

2. Run the development task:

```bash
task dev
```

This will start the development server using Air, allowing automatic rebuilding and restarting on code changes.

## Build a Binary

To build a binary of the application, use the following task:

```bash
task build
```

After building, you can start the application using:

```bash
task start
```

## Running Application in Docker

To run the application in a Docker container, follow these steps:

1. Build the Docker image:

```bash
docker-compose build siberiana_graphql
```

2. Start the application in a Docker container:

```bash
docker-compose up -d siberiana_graphql
```

The application will be accessible at http://localhost:4000.

## Deployment Using Nginx

For deploying the project using Nginx as a reverse proxy, refer to the instructions in the [Siberiana Nginx Setup](https://github.com/dkrasnovdev/siberiana-nginx) repository.

## Troubleshooting

If you encounter any issues or have questions, please feel free to create an issue on this repository.

## Related Repositories

- [Siberiana GraphQL API](https://github.com/dkrasnovdev/siberiana-api)
- [Siberiana Nginx Setup](https://github.com/dkrasnovdev/siberiana-nginx)
- [Siberiana Keycloak Setup](https://github.com/dkrasnovdev/siberiana-keycloak)
- [Siberiana MinIO Setup](https://github.com/dkrasnovdev/siberiana-minio)
- [Siberiana Public Assets](https://github.com/dkrasnovdev/siberiana-public-assets)
