# System Architecture

## Component Breakdown

The system architecture is designed to ensure modularity, scalability, and efficient communication between components. The main components are:

1. **Backend (API)**:
   - The backend is implemented using the Go programming language and the Gin web framework, which provides a lightweight and high-performance HTTP router.
   - The application is organized into a modular folder structure:
      - **`controllers`**: Responsible for handling HTTP requests and sending appropriate responses. Each controller corresponds to a specific feature or module of the application.
      - **`docs`**: Stores the Swagger API documentation, which is automatically generated and updated based on code annotations when building the application.
      - **`models`**: Contains data structure definitions that represent database entities, request and response payloads used in the controllers and other key types used across the application.
      - **`prisma`**: Manages the Prisma configuration and schema, including model definitions, database migrations, and seeding. Prisma is used as the ORM (Object-Relational Mapping) tool to interact with the PostgreSQL database.
      - **`router`**: Sets up the application's routing, defining the paths for each endpoint and integrating middlewares for tasks such as authentication, logging, and error handling. There are two main routers: one the the owner and one for the tenant. The `oauth.go` file contains the OAuth configuration and logic for user authentication.
         - **`middlewares`**: Contains custom middleware functions that can be applied to specific routes or globally to handle common tasks such as authentication, authorization, logging, and error handling. Most of the middlewares store values in the Gin context, which can be accessed later in the controller. This allows for a clean separation of concerns, where the middleware handles cross-cutting concerns while the controllers focus on business logic.
         - **`validators`**: Contains validation logic for incoming requests fields, ensuring that the data meets the required format and constraints before being processed by the controllers. This helps prevent invalid data from being processed and ensures that the application behaves as expected. Validators mostly check enum values.
      - **`services`**: Encapsulates the business logic, providing reusable functions and methods to interact with database, external services, and other components.
         - **`brevo`**: Handles email notifications via Brevo.
         - **`chatgpt`**: Contains logic for interacting with OpenAI's ChatGPT API for AI-powered features such as summarizing and comparing inventory report pictures.
         - **`database`**: Contains functions for database operations, such as querying and updating data. It uses the Prisma ORM to interact with the PostgreSQL database.
         - **`pdf`**: Manages PDF generation, such as creating inventory reports.
      - **`utils`**: Contains utility functions and helper methods that are used across the application for tasks such as error handling, data validation, and formatting. It also includes error codes and constants used throughout the application.
   - The Gin framework allows for a clean separation of concerns, promoting organized code that is easy to maintain and extend.

2. **Database**:
   - The application uses a PostgreSQL database to store and manage all persistent data.
   - The database is accessed via Prisma, a modern ORM that provides type-safe query building, migrations, and schema management.
   - Prisma's migration tool is used to manage changes to the database schema, ensuring that updates are applied in a controlled manner across different environments.

3. **Authentication**:
   - User authentication is implemented using OAuth, with the `github.com/maxzerbini/oauth` package.
   - OAuth is used to manage user login and secure access to the API endpoints. It ensures that protected resources can only be accessed by authenticated users who have valid access tokens.
   - Roles and permissions are managed within the application, allowing fine-grained control over user access to different features and resources.

4. **API Documentation**:
   - The API documentation is hosted using Swagger at `BASE_URL/docs/index.html`.
   - Swagger provides a user-friendly interface for exploring the API endpoints, viewing request and response formats, and testing API interactions directly from the documentation.
   - The documentation is automatically generated from code annotations, ensuring that it remains up-to-date as the codebase evolves.

5. **External Services and Integrations**:
   - The application integrates with external services and APIs via HTTP requests, with the `services` layer handling communication and data processing.
   - One example of an external service is the database management system, which stores and retrieves data from the PostgreSQL database using Prisma.
   - Another example of an external service is the email notification system, which sends notifications to users based on specific events or triggers within the application. Emails are sent using the Brevo service, which is integrated into the `services/brevo` package.
   - The application also uses OpenAI's ChatGPT API to summarize and compare inventory report pictures.

## Data Flow

1. **Request Handling**:
   - A client sends an HTTP request to the server (e.g., to create a new user).
   - The request is received by the `router` and executes any linked middleware functions (e.g., authentication, authorization, global error handling).
   - The request is routed by the `router` to the appropriate `controller` based on the URL path and HTTP method.

2. **Business Logic Execution**:
   - The `controller` calls the relevant function in the `services` layer to perform the necessary operations.
   - Business logic is encapsulated within the `services` layer to keep controllers lightweight.

3. **Database Interaction**:
   - If data needs to be persisted or retrieved, the service interacts with Prisma to communicate with the PostgreSQL database.
   - The `models` layer defines the structure of the data being exchanged.

4. **Response Handling**:
   - The `controller` prepares the response and sends it back to the client, including any data or status messages.

This modular architecture allows the system to be easily maintained, tested, and extended.

## Error Handling

Error handling in the system is designed to provide clear and consistent feedback to clients while maintaining the integrity and security of the application. The main components involved in error handling are:

1. **Controllers**:
   - Controllers are responsible for catching errors that occur during the processing of HTTP requests.
   - When an error is detected, the controller uses utility functions from the `utils` package to send a standardized error response to the client.
   - The `utils.SendError` function is commonly used to send error responses with appropriate HTTP status codes and error messages.

2. **Utils Package**:
   - The `utils` package contains utility functions and constants for error handling.
   - It defines a set of error codes to represent different types of errors.
      - The application defines a set of error codes in the `utils` package to represent different types of errors.
      - These error codes are used in error responses to provide clients with specific information about the nature of the error.
      - This way, clients can easily identify the nature of the error based on the error code and translate the error message into the current language.
      - Examples of error codes include `utils.MissingFields`, `utils.PropertyNotFound` and `utils.PropertyNotYours`.
   - The `utils.SendError` function formats error responses and includes details such as the error code, message, and any additional context.

3. **Middlewares**:
   - Middleware functions are used to handle errors that occur often across multiple routes, such as authentication errors or validation errors.
   - Middleware functions can intercept requests, perform checks, and return an error response if necessary.
   - For example, an authentication middleware can verify the user's access token and return an error if the token is invalid or expired. Also, a validation middleware can check if a resource exists in the database and belongs to the current user. It returns an error if it does not.

4. **Logging**:
   - Errors are logged using the standard Go `log` package to provide visibility into issues that occur during request processing.
   - Every time an error occurs, the `utils.SendError` function calls the gin context's `Error` method to log the error message.
   - Logging helps with debugging and monitoring the application in production environments.

By following these practices, the system ensures that errors are handled gracefully and that clients receive meaningful feedback when issues arise.

## Security

Security is a critical aspect of the system architecture, ensuring that data and operations are protected from unauthorized access and potential threats. The main security measures implemented in the system are:

1. **Authentication and Authorization**:
   - The system uses OAuth for user authentication, ensuring that only authenticated users can access protected resources.
   - Access tokens are used to verify user identity and permissions. Tokens are validated on each request to ensure they are still valid and have not expired.
   - Role-based access control (RBAC) is implemented to manage user permissions. Different roles (e.g., owner, tenant) have different levels of access to resources and operations.

2. **Data Encryption**:
   - Sensitive data, such as passwords, are hashed using bcrypt before being stored in the database. This ensures that even if the database is compromised, the actual passwords are not exposed.
   - Communication between the client and server is secured using HTTPS, encrypting data in transit to prevent interception and tampering.

3. **Input Validation and Sanitization**:
   - Input data is validated and sanitized to prevent common security vulnerabilities such as SQL injection, cross-site scripting (XSS), and cross-site request forgery (CSRF).
   - The Gin framework's built-in validation features are used to ensure that input data meets the required format and constraints.
   - Request data validation is defined in the `models` layer using Go struct annotation, which are automatically validated by Gin.
   - Prisma's query builder is used to prevent SQL injection attacks by automatically escaping user input.

4. **Error Handling**:
   - Detailed error messages are not exposed to clients to prevent leaking sensitive information. Instead, standardized error codes and messages are used.
   - Errors are logged for internal monitoring and debugging purposes, but the logs do not contain sensitive information.

5. **Session Management**:
   - User sessions are managed securely, with options for session expiration and invalidation to prevent unauthorized access.

6. **API Security**:
   - API endpoints are protected using OAuth tokens, ensuring that only authorized users can access them.
   - Rate limiting and throttling mechanisms are implemented to prevent abuse and denial-of-service (DoS) attacks.

7. **Database Security**:
   - The PostgreSQL database is configured with secure access controls, ensuring that only authorized applications and users can connect to it.
   - Regular backups and monitoring are in place to ensure data integrity and availability.

By implementing these security measures, the system ensures that user data and operations are protected from unauthorized access and potential threats, maintaining the integrity and confidentiality of the application.
