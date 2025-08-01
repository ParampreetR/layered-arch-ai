# 📦 Project Structure & Feature Template

This document outlines key features and structural highlights of the project to guide developers, AI agents, and contributors. Use this template to evaluate, maintain, or extend the service while preserving core architecture principles.

---

## 🔍 Key Features

### ✅ 1. **AI Agent Integration via Claude**

* **Path:** `.claude/`

  * `agents/java-to-go-transpiler.md`: Prompts or instructions for Java-to-Go conversion.
  * `settings.local.json`: Configuration for local Claude agent behavior.
* **Purpose:** Automate code transformation and documentation tasks.

---

### ✅ 2. **Microservice Entry Points**

* **Path:** `cmd/`

  * `server/`: Placeholder for HTTP/API server startup code.
  * `worker/`: Placeholder for background jobs or async tasks.
* **Purpose:** Separation of concerns for scalable deployment.

---

### ✅ 3. **Dedicated Internal Module: `erm-transfer-bank`**

* **Path:** `internal/erm-transfer-bank/`

  * `api/`: HTTP layer including `dtos`, `handlers`, `middlewares`, and `routes`.
  * `models/`: Domain-level structs and shared types.
  * `repositories/`: DB access interfaces and implementations.
  * `services/`: Core business logic implementations.
* **Purpose:** Implements the Error Rectification SOP as a standalone microservice.

---

### ✅ 4. **Documentation Structure**

* **Path:** `devdocs/`

  * `AI/`: For AI-specific documentation or workflows.
  * `Devs/`: Developer-facing documentation.
  * `QA/`: QA guidelines, test strategies, or manual test cases.
* **Path:** `docs/`: Reserved for structured user or product documentation.
* **Purpose:** Knowledge transfer, documentation versioning, and team collaboration.

---

### ✅ 5. **Containerization & Automation Scripts**

* **File:** `docker-compose.yml`: Defines service orchestration.
* **Path:** `scripts/`

  * `add_servers_yaml.sh`: YAML configuration appending automation.
  * `compile_swag.sh`: Swagger doc compilation automation.
  * `generate_swagger.sh`: Swagger generation command.
* **Purpose:** DevOps-friendly setup and repeatable build workflows.

---

### ✅ 6. **Governance Rules for AI Contributions**

* **File:** `RULES_FOR_AI.md`
* **Purpose:** Define what AI agents can and cannot do in the repository to avoid architectural or behavioral inconsistencies.

---

### ✅ 7. **Standard Go Project Setup**

* **File:** `go.mod`: Go module definition for dependency management.
* **Folders:** `.git/`, `pkg/`, `README.md`
* **Purpose:** Ensures project is idiomatic, version-controlled, and portable.

---

> ✨ **Note for AI Agents:** Follow the defined folder usage, file responsibilities, and development conventions when generating or modifying code.

---

*Maintained collaboratively by humans and guided by AI.*

