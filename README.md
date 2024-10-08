Go API Kit
==========

Go API Kit is a boilerplate project for building RESTful APIs using Go and PostgreSQL. It follows clean architecture principles, provides user management, authentication, and middleware for handling authentication and profile-specific routes. The project includes migrations for database schema management and Docker configuration for easy setup and deployment.

Features
--------

*   **User Management**: Basic user creation, retrieval, update, and deletion functionalities.
*   **Authentication**: JWT-based authentication for secure access to protected routes.
*   **Middleware**: Includes middleware for authentication and profile-specific route protection.
*   **Database Migrations**: Manage database schema changes using `go-migrate`.
*   **Docker Setup**: Docker Compose configuration for setting up PostgreSQL and Adminer.
*   **Swagger Documentation**: API is documented and accessible via Swagger UI.

Getting Started
---------------

### Prerequisites

*   Go 1.22+
*   Docker
*   Docker Compose

### Installation

1.  Clone the repository:

    `git clone https://github.com/aslam-ep/go-api-kit.git`
    
    `cd go-api-kit`
    
2.  Copy the `.env.example` file to `.env` and update the environment variables as needed.
    
    `cp .env.example .env`
    

### Running the Application

Use the provided `Makefile` to manage the application lifecycle.

1.  **Start Docker Containers**:
    
    To start the PostgreSQL and Adminer containers, use the `make up` command:
    
    `make up`
    
2.  **Run Migrations**:
    
    To apply the database migrations, use the `migrate` command:
    
    `make migrate_up`
    
3.  **Stop Docker Containers**:
    
    To stop the containers, use the `make down` command:
    
    `make down`
    

### API Documentation

The API documentation is provided using Swagger. After starting the application, you can access the Swagger UI at:

`http://localhost:8080/api/v1/swagger/`

Folder Structure
----------------

*   **cmd**: Contains the main application entry point.
*   **config**: Configuration management for environment variables.
*   **database**: Database migration files.
*   **docs**: API documentation files postman collection and swagger.
*   **internal**: Business logic and domain-specific code.
    *   **user**: User-related functionality (handlers, services, repositories, domain models).
    *   **auth**: Authentication-related functionality.
    *   **middleware**: Middlewares for request handling.
*   **router**: Contains router files.
*   *   **middleware**: Middlewares for the restricting routes.
*   **utils**: Utility functions and helpers.

Contributing
------------

Contributions are welcome! Please fork the repository and create a pull request with your changes.

License
-------
This project is licensed under the MIT License. See the [LICENSE](https://github.com/aslam-ep/go-api-kit/blob/main/LICENSE) file for details.