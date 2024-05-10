Project structure below, to be finished


- **api**/: Contains a file to config api documentation, like swagger options
- **cmd/**: Contains the main `main.go` file responsible for starting the API server.
- **internal/**: Contains the internal code of the application, including the controller, DTOs, entities, mappings, repositories, services, and dependency injection.
  - **controller/**: Contains the API controller, which handles HTTP requests.
  - **docs/**: Contains files related to Swagger documentation.
  - **dto/**: Contains the Data Transfer Objects (DTOs) used for communication between layers.
  - **entity/**: Contains domain entities, such as the `User` entity.
  - **mapper/**: Contains mappings between entities and DTOs.
  - **repository/**: Contains repositories responsible for data persistence.
  - **service/**: Contains services that implement business logic.
  - **di/**: Contains files related to dependency injection.
- **pkg/**: Can contain shared packages between different parts of the application.
- **config/**: Contains application configuration files, if any.
- **db/**: Can contain files related to the database, such as migration scripts.
- **migrations/**: Contains database migration scripts.
