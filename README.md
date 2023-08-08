<p align="center">
  <picture>
  <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/dkrasnovdev/siberiana-public-assets/main/assets/siberiana-logo-dark-background.svg">
  <img src="https://raw.githubusercontent.com/dkrasnovdev/siberiana-public-assets/main/assets/siberiana-logo-dark-background.svg" width="240" height="90" alt="Logo for Siberiana">
</picture>
</p>

<p align="center">
Siberiana | Aggregator of the Historical and Cultural Heritage of the Yenisei Siberia.
</p>

# Siberiana GraphQL API Setup

## Prerequisites

Before you begin, ensure you have the following installed:

- Go (at least version 1.19)
- Docker
- Docker Compose

## Getting Started

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/dkrasnovdev/siberiana-api.git
   cd siberiana-api
   ```

2. Copy the `.env.example` file and rename it to `.env`. Fill in the required environment variables with appropriate values.

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
