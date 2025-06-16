# Deployment Architecture

## Overview

This document provides a detailed guide on building and deploying the Keyz project. The project uses Docker containers for consistent environments and GitHub Actions for continuous integration and deployment.

## Docker Containers

The Keyz project is structured to run in a multi-container Docker environment. Each component of the application is encapsulated in its own Docker container, allowing for modular development and deployment. A `compose.yaml` file is used to define and manage these containers, ensuring that they can communicate with each other and share resources as needed.

The following sections describe the Dockerfiles for each component and how they interact within the Docker Compose setup.

### Dockerfile for Web Application

The web application is built for production using Node.js and Vite. The Dockerfile for the web application is located in the [Code/Web/](/Code/Web/Dockerfile) directory.

### Dockerfile for Showcase website

The showcase website is built using Node.js and Vite, similar to the web application. The Dockerfile for the showcase website is located in the [Code/ShowcaseKeyz/](/Code/ShowcaseKeyz/Dockerfile) directory.

### Dockerfile for Database migration

The database migration is handled using Prisma, an ORM for PostgreSQL. The Dockerfile for the database migration is located in the [Code/Server/](/Code/Server/Dockerfile.migrate) directory. This Dockerfile sets up the environment to run Prisma migrations against the PostgreSQL database.

### Dockerfile for Server Application

The server application is built using Go and uses the Gin framework. The Dockerfile for the server application is located in the [Code/Server/](/Code/Server/Dockerfile) directory.

### Docker Compose

The [compose.yaml](/Code/compose.yaml) file orchestrates the multi-container application. It defines five services: `showcase`, `frontend`, `server`, `migrate` and `db`. Once the `db` service is up and running, the `migrate` service is run to apply database migrations. Once the migrations are complete and the container exits, the `server` is started. Then, the `frontend` and `showcase` services are started. The `db` service uses a PostgreSQL image and mounts a volume for persistent data storage. The `server` and `migrate` services use a `.env.docker` file for environment variables, and the `db` service uses a `password.txt` file for the database password.

## Continuous Integration and Deployment (CI/CD)

The Keyz project employs a robust CI/CD pipeline using GitHub Actions. This pipeline automates the build, test, and deployment processes for the web application, server application, and Android application. The CI/CD setup ensures that code changes are automatically tested and deployed to development and production environments, maintaining high code quality and reducing manual intervention.

### GitHub Actions

The project uses GitHub Actions for CI/CD, leveraging a modular structure with reusable workflows and dedicated deployment workflows. All workflow files are located in the `.github/workflows/` directory.

#### Reusable Workflows

- **build-web.yml**: Builds and tests the web application using Node.js and Vite. Supports multiple Node.js versions and runs on pushes to the `web-dev` branch or pull requests to `main`.
- **build-server.yml**: Builds and tests the server application using Go. Runs linter, unit tests, and enforces coverage requirements. Triggered on pushes to the `backend` branch or pull requests to `main`.
- **build-android.yml**: Runs instrumentation and unit tests for the Android application. Triggered on pushes to the `android-dev` branch or pull requests to `main`.
- **create-version-tag.yml**: Handles versioning and creates a new version tag based on the `version.rc` file, ensuring semantic versioning and preventing regressions. This workflow is triggered by the `deploy-prod.yml` workflow after a successful deployment to production.

#### Deployment Workflows

- **deploy-dev.yml**: Deploys the application to the development environment on pushes to the `deploy/dev` branch. This workflow first calls the reusable build workflows (`build-web.yml`, `build-server.yml`, `build-android.yml`) to ensure all components are built and tested. If successful, it deploys the application to the dev server using SSH commands to pull the latest code and restart the Docker containers.

- **deploy-prod.yml**: Deploys the application to the production environment on pushes to the `main` branch. Like the dev workflow, it first calls the reusable build workflows. After a successful deployment, it triggers the `create-version-tag.yml` workflow to generate a new version tag, and then creates a GitHub release using the new tag. This ensures that every production deployment is versioned and released with release notes, which are automatically generated from the commit history and pull requests since the last release.

This modular approach ensures that builds are consistent across environments, and deployments are automated, versioned, and traceable. All secrets and environment variables required for deployment are managed via GitHub environments that contain secrets and variables.

## Deployment environment

The application is deployed to a private VM running Ubuntu 24.04 and is accessible with a domain name. The deployment process is automated, ensuring that the latest changes are deployed to the development and production environments without manual intervention.

The VMs are configured with Docker and Docker Compose to run the multi-container application. The application is deployed behind an Nginx reverse proxy to handle incoming requests and route them to the appropriate services. The Nginx configuration includes SSL certificates for secure communication over HTTPS.
