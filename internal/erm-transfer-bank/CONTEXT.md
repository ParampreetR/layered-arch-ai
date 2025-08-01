# Project Folder Structure Guide for AI Agents

This `CONTEXT.md` is specifically written to help **AI agents and collaborators** understand the **structure, purpose, and logic** behind each folder in this project. Follow this guide to write code in the correct locations and preserve project architecture integrity.

> ✅ **IMPORTANT FOR AI AGENTS:** Always respect the structure and purpose of each directory. If you're unsure where a file should go, refer to this guide before creating new directories or placing code.

---

## Overview

```
.
├── api
│   ├── dtos
│   ├── handlers
│   ├── middlewares
│   └── routes
├── models
├── repositories
└── services
```

Each of these directories plays a role in a **layered architecture** model — separating concerns between data definitions, business logic, API interfaces, and persistence mechanisms.

---

## 📁 `api/`

This folder handles all **HTTP interface logic**, acting as the gateway to your application. It contains:

### `api/dtos/`

* **Purpose:** Define **Data Transfer Objects (DTOs)** to map incoming or outgoing payloads.
* **Used by:** Handlers and external clients.
* **AI Agent Note:** Do not define domain logic or models here. Use this for shaping request/response formats only.

### `api/handlers/`

* **Purpose:** Contains **HTTP handler functions** that process requests.
* **Responsibilities:**

  * Parse input from DTOs
  * Call relevant services
  * Return formatted responses
* **AI Agent Note:** Avoid direct DB calls or core logic here — delegate to `services/`.

### `api/middlewares/`

* **Purpose:** Houses HTTP **middleware** like logging, CORS, authentication, request parsing, etc.
* **Example:** JWT validation, rate limiting.

### `api/routes/`

* **Purpose:** Wire up all routes to the appropriate handlers.
* **Pattern:** Usually contains route registration functions (e.g. `RegisterUserRoutes`).
* **AI Agent Note:** Do not put logic here — only route definitions.

---

## 📁 `models/`

* **Purpose:** Define core **domain structs/entities** shared across the project.
* **Examples:** `User`, `Order`, `Product`, etc.
* **AI Agent Note:**

  * These models represent the database or domain structures.
  * Should be reused across services, repositories, and sometimes DTOs with conversion logic.
  * Avoid putting database queries here.

---

## 📁 `repositories/`

* **Purpose:** Encapsulate **data access logic** (e.g., database operations).
* **Pattern:** Use interfaces for repositories and implement them per DB (e.g., PostgreSQL, MongoDB).
* **Example Structure:**

  * `UserRepository interface`
  * `PostgresUserRepository struct`
* **AI Agent Note:**

  * Do not implement business logic here.
  * Keep repository methods focused on **CRUD and query logic**.

---

## 📁 `services/`

* **Purpose:** Implement the **business logic** of the application.
* **Responsibilities:**

  * Coordinate between repositories and handlers.
  * Enforce domain rules and workflow.
  * Reuse across interfaces (e.g., CLI, HTTP).
* **AI Agent Note:**

  * Keep this free from HTTP-specific logic.
  * Services should be testable in isolation.

---

## 🧠 AI Agent Development Guidelines

* 🛑 **NEVER** add new folders unless you are explicitly instructed to.
* ✅ Always reuse models, services, and repositories.
* 🧩 If a file does not clearly fit in any category:

  * Re-evaluate if it belongs in `pkg/` (if exists)
  * Consult human developer for placement
* 🧹 Maintain separation of concerns:

  * API layer (input/output)
  * Service layer (logic)
  * Repository layer (data access)
  * Model layer (data structure)

---

## 🔧 Future Enhancements

If the project grows, the following folders may be introduced:

* `config/` – Central app configurations
* `pkg/` – Reusable packages across different services
* `internal/` – Private modules scoped to the app
* `scripts/` – Utility scripts or automation helpers

> ⚠️ AI Agents: Do not create these directories unless specified by maintainers.

---

## 💬 Summary

| Folder          | Responsibility           | Do Not Include                    |
| --------------- | ------------------------ | --------------------------------- |
| `api/`          | Web layer, HTTP handling | Business logic, direct DB access  |
| `models/`       | Domain structs/entities  | API-specific formatting           |
| `repositories/` | DB access, CRUD logic    | Business decisions, HTTP handling |
| `services/`     | Core application logic   | Route handling, DB query logic    |

Let this markdown be the **source of truth** for all directory conventions. Respect it while generating or editing any file.

IMPORTANT: IN ANY CASE YOU CAN NOT EDIT OR DELETE THIS FILE!!!

*Maintained by: Human engineers & enforced by AI guardians 🤖.*

