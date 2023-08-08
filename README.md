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

- Go (at least version 1.19)
- Docker
- Docker Compose

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

Follow the steps outlined in the next section to start containers, run tasks, and access the application.

## Getting Started

1.  Clone this repository to your local machine:

```bash
git clone https://github.com/dkrasnovdev/siberiana-api.git
cd siberiana-api
```

2.  Run the following command to download the required packages specified in the Taskfile:

```bash
task install
```

3.  Copy the `.env.example` file and rename it to `.env`. Fill in the required environment variables with appropriate values.

```bash
cp .env.example .env
# Edit .env file with your configuration
```

## Troubleshooting

If you encounter any issues or have questions, please feel free to create an issue on this repository.

## Related Repositories

- [Siberiana GraphQL API](https://github.com/dkrasnovdev/siberiana-api)
- [Siberiana Nginx Setup](https://github.com/dkrasnovdev/siberiana-nginx)
- [Siberiana Keycloak Setup](https://github.com/dkrasnovdev/siberiana-keycloak)
- [Siberiana MinIO Setup](https://github.com/dkrasnovdev/siberiana-minio)
- [Siberiana Public Assets](https://github.com/dkrasnovdev/siberiana-public-assets)
