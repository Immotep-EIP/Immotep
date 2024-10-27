# Setup and Installation

This section provides a step-by-step guide to setting up the development environment, running the application locally, and deploying it to various environments.

## Prerequisites

Before setting up the project, ensure that you have the following installed on your machine:

1. **Go**: The application is built using the Go programming language. Install the latest version from [golang.org](https://golang.org/).
2. **PostgreSQL**: The project uses PostgreSQL as the database. You can download it from [postgresql.org](https://www.postgresql.org/).
3. **golangci-lint**: The project uses a linter to check code readability. You can install it from [golangci-lint.run](https://golangci-lint.run/).
4. **Swag**: The API documentation is build using Swagger with a Go package named Swaggo. You can install it from [github.com/swaggo/swag](https://github.com/swaggo/swag).

## Local Setup

Follow these steps to set up the project locally:

### 1. Clone the Repository

Clone the project repository from your version control system (e.g., GitHub or GitLab):

```bash
git clone https://github.com/Immotep-EIP/Immotep.git
cd Immotep
```

### 2. Configure Environment Variables

Create a `.env` file in the root directory from the example file and configure the required environment variables. Below is an example of the environment variables needed:

```bash
cp .env.example .env
```

```plaintext
PORT='3001'
DATABASE_URL='postgresql://user:password@localhost:5432/immotep'
SECRET_KEY='MySecretKey'
```

- Replace `user` and `password` with your actual PostgreSQL credentials.
- Configure SecretKey settings with the appropriate values for your authentication setup.

### 3. Install Dependencies

Ensure Go modules are enabled and then download the project dependencies:

```bash
go mod tidy
```

### 4. Set Up the Database

1. **Start PostgreSQL**: If PostgreSQL is not already running, start it using the native installation from your OS or using the Docker image.
2. **Run Prisma Migrations**: Apply database migrations and set up the database schema:
   ```bash
   ./update.sh
   ```

### 5. Seed the Database (Optional)

TODO

### 6. Run the Application

Start the backend server:

```bash
./run.sh
```

The application should now be accessible at `http://localhost:3001`.

## Deployment Instructions

TODO

## Running Tests

Run the project's tests using:

```bash
./run_tests.sh
```

Include unit tests, integration tests, and any mock tests you have written for the API.
