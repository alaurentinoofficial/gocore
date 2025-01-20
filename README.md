# GO CORE

Go Core is a reusable library of common services designed for my personal Go projects. It follows the Dependency Inversion Principle (DIP) to promote flexibility, scalability, and reduced coupling in application design.

## Key Features

* *Plug-and-Play Components:* Easily integrate service components into your project.
* *Flexible Dependency Injection:* Compatible with both manual dependency injection and the powerful Uber Fx framework for lifecycle management and dependency resolution.

## Benefits of Using Dependency Inversion

* *Reduced Coupling:* Decouples high-level modules from low-level modules, making your codebase more modular and easier to manage.
* *Improved Testability:* Allows swapping out real implementations with mocks or stubs during testing, leading to more reliable and isolated unit tests.
* *Enhanced Maintainability:* Changes in low-level details or implementations do not require modifications to high-level modules.
* *Reusability:* Promotes the use of abstract interfaces, making your code reusable across multiple projects or contexts.
* *Scalability:* Simplifies adding new features or components without affecting existing implementations.
* *Clear Structure:* Encourages clean architecture principles, improving the overall structure and readability of your codebase.

## How to Use
Manual Dependency Injection: Import the required components and inject them directly into your application logic.
Using Uber Fx: Plug the components into your fx.App configuration for automatic lifecycle management.

---

Authored by *Anderson Laurentino*
