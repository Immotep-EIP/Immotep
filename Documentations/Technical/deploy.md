# Deployment Architecture

## Overview

This document provides a detailed guide on building and deploying the Immotep project. The project uses Docker containers for consistent environments and GitHub Actions for continuous integration and deployment.

## Docker Containers

### Dockerfile for Web Application

The web application is built for production using Node.js and Vite. The Dockerfile for the web application is located in the [Code/Web/](/Code/Web/Dockerfile) directory.

### Dockerfile for Server Application

The server application is built using Go and uses the Gin framework. Before starting the server, it runs the Prisma migration to deploy the database schema to the PostgreSQL database. The Dockerfile for the server application is located in the [Code/Server/](/Code/Server/Dockerfile) directory.

### Docker Compose

The [compose.yaml](/Code/compose.yaml) file orchestrates the multi-container application. It defines three services: `frontend`, `server`, and `db`. The `frontend` service depends on the `server` service, and the `server` service depends on the `db` service. The `db` service uses a PostgreSQL image and mounts a volume for persistent data storage. The `server` service uses a `.env.docker` file for environment variables, and the `db` service uses a `password.txt` file for the database password.

## Continuous Integration and Deployment (CI/CD)

### GitHub Actions

The project uses GitHub Actions for CI/CD. There are two workflows:

- [dev.yml](/.github/workflows/dev.yml): This workflow builds, tests, and deploys the application to the development environment on every push to the `dev` branch or pull request to the `main` branch. Here are the different jobs:
  - `build-web`: Builds the web application using Node.js.
  - `build-server`: Builds the server application using Go, runs linter, runs unit tests, and checks coverage requirements.
  - `build-android`: Runs instrumentation tests for the Android application.
  - `deploy-dev`: Deploys the application to the development environment if everything went well using SSH commands (pulls the latest code, builds the Docker images, and starts the containers).
- [prod.yml](/.github/workflows/prod.yml): This workflow deploys the application to the production environment on every push to the `main` branch using SSH commands (pulls the latest code, builds the Docker images, and starts the containers).

## Deployment environment

The application is deployed to a private VM running Ubuntu 24.04 and is accessible with a domain name. The deployment process is automated, ensuring that the latest changes are deployed to the development and production environments without manual intervention.

The VMs are configured with Docker and Docker Compose to run the multi-container application. The application is deployed behind an Nginx reverse proxy to handle incoming requests and route them to the appropriate services. The Nginx configuration includes SSL certificates for secure communication over HTTPS.
