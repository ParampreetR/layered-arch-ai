# Project Guidelines - CRA Error Rectification Service

This document outlines the coding standards, conventions, and best practices for the CRA Error Rectification Service project. All developers and AI agents working on this codebase **MUST** follow these guidelines to maintain consistency and quality.

---

## 📁 Project Structure

### Root Structure
```
.
├── cmd/                    # Application entry points
│   ├── server/            # HTTP server main
│   └── worker/            # Background worker main
├── internal/              # Private application code
│   ├── common/           # Shared internal packages
│   └── erm-transfer-bank/ # Core service module
├── pkg/                   # Public reusable packages
├── scripts/              # Build and automation scripts
├── devdocs/              # Development documentation
├── docs/                 # User-facing documentation
└── .claude/              # AI agent configurations
```

### Module Structure Pattern
Each service module (e.g., `erm-transfer-bank`) follows this structure:
```
module/
├── api/
│   ├── dtos/         # Data Transfer Objects
│   ├── handlers/     # HTTP handlers
│   ├── middlewares/  # HTTP middlewares
│   └── routes/       # Route definitions
├── di/               # Dependency injection
├── models/           # Domain models
├── repositories/     # Data access layer
├── services/         # Business logic layer
└── CONTEXT.md        # Module-specific documentation
```

---

## 🏗️ Architecture Principles

### 1. **Layered Architecture**
- **API Layer**: HTTP handling, request/response formatting
- **Service Layer**: Business logic, domain rules
- **Repository Layer**: Data access, database operations
- **Model Layer**: Domain entities, data structures

### 2. **Separation of Concerns**
- No business logic in handlers
- No HTTP concerns in services
- No direct DB access outside repositories
- Keep layers loosely coupled through interfaces

### 3. **Dependency Injection**
- Use DI containers for managing dependencies
- Initialize dependencies in `di/init.go` files
- Store dependencies in structured containers

---

## 📝 Naming Conventions

### 1. **Package Naming**
- Use **snake_case** for package names
- Include module/layer in package name for clarity
- Examples:
  ```go
  package common_di
  package common_config
  package master_models
  package erm_transfer_bank_di
  ```

### 2. **File Naming**
- Use **snake_case** for Go files: `prao_mst.go`, `env.go`
- Use **_impl.go** suffix for implementation files
- Use **_test.go** suffix for test files
- Use **kebab-case** for markdown files: `go-cocreator.md`
- Use **UPPERCASE** for special documentation: `CONTEXT.md`, `RULES_FOR_AI.md`

### 3. **Interface Naming**
- Start interface names with **'I'** prefix
- Use PascalCase after the prefix
- Examples:
  ```go
  type IUserService interface { }
  type IRepository interface { }
  type IAuthMiddleware interface { }
  ```

### 4. **Struct Naming**
- Use **PascalCase** for structs
- Be descriptive and domain-specific
- Examples:
  ```go
  type PraoMst struct { }
  type ProjectDependencies struct { }
  type Config struct { }
  ```

### 5. **Function/Method Naming**
- Use **PascalCase** for exported functions/methods
- Use **camelCase** for unexported functions/methods
- Be descriptive about what the function does
- Examples:
  ```go
  func InitDB(cfg *Config) (*gorm.DB, error)
  func getEnv(key, defaultValue string) string
  ```

---

## 📏 Code Organization Rules

### 1. **File Size Limit**
- **CRITICAL**: No file should exceed 250 lines
- Split large files into logical components
- Use separate files for interfaces and implementations

### 2. **File Structure Pattern**
- Keep interfaces in dedicated interface files
- Place implementations in `*_impl.go` files
- Group related functionality together
- Example:
  ```
  services/
  ├── user_service.go      # Interface definition
  ├── user_service_impl.go  # Implementation
  └── user_service_test.go  # Tests
  ```

### 3. **Import Organization**
- Group imports in this order:
  1. Standard library imports
  2. Third-party imports
  3. Internal imports
- Use blank lines between groups
- Use import aliases for clarity when needed

---

## 🔧 Coding Standards

### 1. **Error Handling**
- Always check and propagate errors
- Wrap errors with context using `fmt.Errorf`
- Never use `panic` for error handling
- Example:
  ```go
  if err != nil {
      return fmt.Errorf("failed to connect to database: %w", err)
  }
  ```

### 2. **Configuration Management**
- Use environment variables for configuration
- Provide sensible defaults
- Never hardcode values
- Use the `Config` struct pattern

### 3. **Database Conventions**
- Use GORM tags for column mapping
- Match database column names exactly
- Use pointers for nullable fields
- Include proper indexes and constraints
- Example:
  ```go
  type Model struct {
      ID        int       `gorm:"column:id;primaryKey;autoIncrement"`
      Name      string    `gorm:"column:name;type:varchar(100);not null"`
      DeletedAt *time.Time `gorm:"column:deleted_at;type:timestamp"`
  }
  ```

### 4. **Logging Standards**
- Use structured logging
- Include context in log messages
- Use appropriate log levels
- Avoid logging sensitive information

### 5. **Testing Requirements**
- Write unit tests for all services and repositories
- Place test files next to implementation files
- Use table-driven tests where appropriate
- Mock external dependencies

---

## 📂 Documentation Standards

### 1. **Documentation Structure**
```
devdocs/
├── AI/         # AI-generated documentation (cache/memory)
├── Devs/       # Human-verified developer docs
├── Generated/  # AI-assisted but human-reviewed docs
└── QA/         # Testing documentation
```

### 2. **Documentation Rules**
- **AI agents** can only write to `/devdocs/AI/`
- **Developers** write to `/devdocs/Devs/` after manual verification
- Each module must have a `CONTEXT.md` file
- Update documentation when making significant changes

### 3. **Code Comments**
- Add comments for complex logic
- Document public APIs with godoc comments
- Keep comments concise and relevant
- Example:
  ```go
  // InitDB initializes a database connection with the provided configuration.
  // It returns a configured gorm.DB instance or an error if connection fails.
  func InitDB(cfg *common_config.Config) (*gorm.DB, error) {
  ```

---

## 🚫 Prohibited Practices

### 1. **Never Do**
- Modify `CONTEXT.md` files in modules
- Create circular dependencies
- Use `any` type without strong justification
- Hardcode configuration values
- Mix concerns between layers
- Exceed 250 lines per file
- Use panic for error handling
- Store sensitive data in code

### 2. **AI Agent Restrictions**
- Cannot modify `/devdocs/Devs/` folder
- Cannot edit `RULES_FOR_AI.md` or `CONTEXT.md`
- Must follow all conventions in this document
- Must update `/devdocs/AI/` after significant work

---

## 🔄 Development Workflow

### 1. **Before Starting**
- Read the module's `CONTEXT.md`
- Understand the existing patterns
- Check for reusable components in `/pkg/`

### 2. **While Coding**
- Follow the layered architecture
- Keep files under 250 lines
- Write tests alongside implementation
- Use dependency injection

### 3. **After Completing**
- Update relevant documentation
- Ensure all tests pass
- Verify no circular dependencies
- Check file sizes

---

## 🎯 Module-Specific Guidelines

### Common Module (`internal/common/`)
- Shared utilities and configurations
- Database initialization and migrations
- Common models used across modules
- Base dependency injection setup

### ERM Transfer Bank Module (`internal/erm-transfer-bank/`)
- Implements Error Rectification SOP
- Follows Maker → Checker → Authorizer flow
- Handles fund transfer rectifications
- Must read `CONTEXT.md` before modifications

---

## 🔍 Code Review Checklist

- [ ] All files under 250 lines
- [ ] Interfaces start with 'I' prefix
- [ ] Package names use snake_case
- [ ] No business logic in handlers
- [ ] Errors properly wrapped with context
- [ ] Tests written for new code
- [ ] Documentation updated
- [ ] No circular dependencies
- [ ] Follows layered architecture
- [ ] Uses dependency injection properly

---

*This document is maintained collaboratively. AI agents must follow these guidelines strictly. Developers should update this document when establishing new patterns.*