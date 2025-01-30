# Internal Directory Structure

This directory follows the principles of Clean Architecture, organizing the code into separate layers to ensure maintainability, testability, and scalability.

## Structure Overview

```
internal/
├── application/
├── domain/
│   ├── entities/
│   └── interfaces/
├── infrastructure/
└── lib/
```

### application/
This layer contains the use cases of the application. It orchestrates business logic and serves as an intermediary between the domain and the infrastructure layers.

### domain/
The core business logic of the application is placed in this layer. It follows a domain-driven design approach and is divided into:
- **entities/**: Contains the core business entities, which define the application's essential data structures.
- **interfaces/**: Defines interfaces that help decouple business logic from external dependencies, such as repositories and services.

### infrastructure/
This layer contains the implementation details of external dependencies, such as database connections, API clients, and messaging systems.

### lib/
A dedicated package for utility functions and reusable components that are project-agnostic. Important constraints for this directory:
- Files inside `lib` **must not import** anything outside of `lib`.
- They can import other libraries inside `lib` or external libraries.

#### Example:
```
lib/
└── logger/
    └── logger.go  # Provides a logging utility
```

This structure ensures a clear separation of concerns, enabling a scalable and maintainable codebase.
