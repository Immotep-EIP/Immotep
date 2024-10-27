# System Architecture

## Component Breakdown

The system architecture is designed to ensure modularity, scalability, and efficient communication between components. The main components are:

1. **Backend (API)**:
   - The backend is implemented using the Go programming language and the Gin web framework, which provides a lightweight and high-performance HTTP router.
   - The application is organized into a modular folder structure:
     - **`controllers`**: Responsible for handling HTTP requests and sending appropriate responses. Each controller corresponds to a specific feature or module of the application.
     - **`models`**: Contains data structure definitions that represent database entities and other key types used across the application.
     - **`prisma`**: Manages the Prisma configuration and schema, including model definitions, database migrations, and seeding. Prisma is used as the ORM (Object-Relational Mapping) tool to interact with the PostgreSQL database.
     - **`router`**: Sets up the application's routing, defining the paths for each endpoint and integrating middleware for tasks such as authentication, logging, and error handling.
     - **`services`**: Encapsulates the business logic, providing reusable functions and methods to interact with database, external services, and other components.
   - The Gin framework allows for a clean separation of concerns, promoting organized code that is easy to maintain and extend.

2. **Database**:
   - The application uses a PostgreSQL database to store and manage all persistent data.
   - The database is accessed via Prisma, a modern ORM that provides type-safe query building, migrations, and schema management.
   - Prisma's migration tool is used to manage changes to the database schema, ensuring that updates are applied in a controlled manner across different environments.

3. **Authentication**:
   - User authentication is implemented using OAuth, with the `github.com/maxzerbini/oauth` package.
   - OAuth is used to manage user login and secure access to the API endpoints. It ensures that protected resources can only be accessed by authenticated users who have valid access tokens.
   - Roles and permissions can be implemented to control access to specific API routes, enabling fine-grained access control.

4. **API Documentation**:
   - The API documentation is hosted using Swagger at `BASE_URL/docs/index.html`.
   - Swagger provides a user-friendly interface for exploring the API endpoints, viewing request and response formats, and testing API interactions directly from the documentation.
   - The documentation is automatically generated from code annotations, ensuring that it remains up-to-date as the codebase evolves.

5. **External Services and Integrations**:
   - If the application integrates with any external services (e.g., third-party APIs, payment gateways), these interactions should be documented here.
   - For example, describe how the services are accessed, any authentication requirements, and the format of requests and responses.

## Data Flow

1. **Request Handling**:
   - A client sends an HTTP request to the server (e.g., to create a new user).
   - The request is routed by the `router` to the appropriate `controller` based on the URL path and HTTP method.

2. **Business Logic Execution**:
   - The `controller` calls the relevant function in the `services` layer to perform the necessary operations.
   - Business logic is encapsulated within the `services` layer to keep controllers lightweight.

3. **Database Interaction**:
   - If data needs to be persisted or retrieved, the service interacts with Prisma to communicate with the PostgreSQL database.
   - The `models` layer defines the structure of the data being exchanged.

4. **Response Handling**:
   - The `controller` prepares the response and sends it back to the client, including any data or status messages.
   - If an error occurs, it is handled by middleware before returning an appropriate response to the client.

This modular architecture allows the system to be easily maintained, tested, and extended.
