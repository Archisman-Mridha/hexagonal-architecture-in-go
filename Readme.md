# Hexagonal Architecture in GoLang

Demonstrating `Hexagonal Architecture` in GoLang.

## The application layer

Here are the *usecases of the application layer* in *large scale distributed systems* -

- **Enforcing security and authorization**: The application layer can implement authentication and authorization checks, and ensure that only authorized requests are forwarded to the business logic layer. This helps to protect the application against unauthorized access and data breaches.

- **Handling cross-cutting concerns**: The application layer can implement common features, such as logging, caching, and error handling, that are used by multiple use cases in the application. By encapsulating these concerns in the application layer, you can ensure that they are applied consistently across the application.

- **Adapting requests and responses**: The application layer can translate between the format used by the inbound adapter (e.g., gRPC messages) and the format used by the business logic layer (e.g., domain models). This can simplify the implementation of the business logic layer, as it does not need to be aware of the specific formats used by each adapter.