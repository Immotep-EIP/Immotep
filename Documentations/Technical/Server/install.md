# Setup and Installation

This section provides a step-by-step guide to setting up the development environment and running the application locally.

## Prerequisites

Before setting up the project, ensure that you have the following installed on your machine:

1. **Go**: The application is built using the Go programming language. Install the latest version from [golang.org](https://golang.org/).
2. **PostgreSQL**: The project uses PostgreSQL as the database. You can download it from [postgresql.org](https://www.postgresql.org/).
3. **golangci-lint**: The project uses a linter to check code readability. You can install it from [golangci-lint.run](https://golangci-lint.run/).
4. **Swag**: The API documentation is build using Swagger with a Go package named Swaggo. You can install it from [github.com/swaggo/swag](https://github.com/swaggo/swag).
5. **Make**: The project uses Makefile for build automation. Ensure you have `make` installed on your system. You can install it using your package manager (e.g., `apt`, `brew`, etc.).

## Local Setup

Follow these steps to set up the project locally:

### 1. Clone the Repository

Clone the project repository from your version control system (e.g., GitHub or GitLab):

```bash
git clone https://github.com/Immotep-EIP/Immotep.git
cd Immotep/Code/Server
```

### 2. Configure Environment Variables

Create a `.env` file in the root directory from the example file and configure the required environment variables. Below is an example of the environment variables needed:

```bash
cp .env.example .env
```

```txt
PORT='3001'
PUBLIC_URL='http://localhost:3001'
WEB_PUBLIC_URL='http://localhost:4242'
SHOWCASE_PUBLIC_URL='http://localhost:4343'
DATABASE_URL='postgresql://user:password@localhost:5432/immotep'
SECRET_KEY='MySecretKey'
OPENAI_API_KEY='MyOpenAIKey'
BREVO_API_KEY='MyBrevoKey'
```

- Replace `user` and `password` with your actual PostgreSQL credentials.
- Configure SecretKey settings with the appropriate values for your authentication setup.
- Configure OpenAI API Key with the appropriate value for your OpenAI account.
- Configure Brevo API Key with the appropriate value for your Brevo account.

### 3. Install Dependencies

Ensure Go modules are enabled and then download the project dependencies:

```bash
go mod tidy
go mod download
```

### 4. Set Up the Database

1. **Start PostgreSQL**: If PostgreSQL is not already running, start it using the native installation from your OS or using the Docker image.
2. **Run Prisma Migrations**: Apply database migrations and set up the database schema:
   ```bash
   cd ..
   ./update.sh
   cd Server
   ```

### 5. Seed the Database (Optional)

```bash
make db_reset
```

### 6. Run the Application

Build and start the backend server:

```bash
make run
```

The application should now be accessible at `http://localhost:3001`.

If you face any issues with the linter, you can run the following command to run the server without linting:

```bash
make build && ./backend
```

## Deployment Instructions

TODO

## Running Tests

Run the project's tests using:

```bash
make test
```

Include unit tests, integration tests, and any mock tests you have written for the API.
