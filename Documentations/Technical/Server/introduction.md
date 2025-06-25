# Introduction

## Project Overview

This application is the backend part of the Keyz application. It provides an API responsible for managing the application's core functionalities, including user management, data processing, and integration with external services. The API communicates with the frontend applications (web and mobile) and is directly connected to a database to store and manage data.

## Technologies used

The API is built using the Go programming language with the Gin framework, ensuring efficient handling of requests and responses. The system relies on a PostgreSQL database, leveraging Prisma as the ORM (Object-Relational Mapping) tool to facilitate seamless interactions with the database. User authentication is implemented using OAuth, using the `github.com/maxzerbini/oauth` package to provide secure access control to protected resources. The backend also integrates with OpenAI's ChatGPT for AI-powered features, Brevo for email notifications, and supports PDF generation for inventory reports.

## API documentation

API documentation is available via Swagger, allowing developers to explore and test API endpoints directly. The documentation is hosted at `BASE_URL/docs/index.html` and is automatically updated based on code comments to reflect the current state of the API.

## System Architecture

The system's architecture is available [here](./architecture.md).

## Deployment Architecture

The deployment architecture for the whole project is available [here](../deploy.md).
