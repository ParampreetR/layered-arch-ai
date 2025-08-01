> **Purpose**: This file serves as a central rulebook for AI agents interacting with this project. All AI tasks—whether generating code, updating documentation, or refactoring—**must follow these rules and best practices** to ensure project consistency, maintainability, and safety.

---

## 🧠 General Philosophy

AI agents must:

* **Behave deterministically** (repeatable output for the same input).
* **Preserve developer intent** where unclear instructions exist.
* **Be cautious and conservative** when editing files not explicitly mentioned.
* **Leave no unexplained side effects** in commits or code generation.

---

## 📁 Folder Structure Explanation

| Path                          | Purpose                                                                                    |
| ----------------------------- | ------------------------------------------------------------------------------------------ |
| `/devdocs/AI`                 | AI agent-specific documentation. You may store your logic here if needed.                  |
| `/devdocs/Devs`               | Developer-facing documentation. Do **not** overwrite unless instructed.                    |
| `/devdocs/QA`                 | QA/testing-related docs. Only modify if the change directly relates to tests you generate. |
| `/docs`                       | General project documentation. Add new files carefully.                                    |
| `/internal/erm-transfer-bank` | Core service logic. All AI-generated features related to this service must go here.        |
| `├── api/`                    | Request/response logic: dtos, handlers, middleware, routes.                                |
| `├── models/`                 | Business logic data models. Follow naming and validation conventions.                      |
| `├── repositories/`           | Data access layer. Use standard repo patterns.                                             |
| `├── services/`               | Core business logic lives here. Services must be stateless and testable.                   |
| `├── CONTEXT.md`              | Context file for explaining this module. Read before generating anything.                  |
| `/pkg/`                       | Shared libraries/utilities. Reuse instead of rewriting logic.                              |
| `/scripts/`                   | Shell scripts for CI/CD, build, or automation. Respect script usage and naming patterns.   |

---

## 🛡️ Behavioral Rules for AI

1. **Do not modify files** outside your scope unless explicitly asked.
2. **Always read `CONTEXT.md`** and related documents before generating/modifying code inside a module.
3. When creating new files:

   * Use **snake\_case for shell scripts**, **PascalCase for Go structs**, and **kebab-case for markdown/doc files**.
   * Add proper headers and comments where necessary.
4. **Update relevant docs automatically** if you make meaningful changes in logic or structure.
5. Do **not hardcode values**; respect environment variables or configuration files.
6. **Avoid circular dependencies** in Go packages.
7. Use **structured logging** and respect existing logger interfaces.
8. **Do not touch `/go.mod` or `/pkg/`** unless required for feature implementation and approved by prompt/context.

---

## 🧪 Testing Expectations

* Any new handler, service, or repository must be accompanied by a corresponding **unit test**.
* Place test files next to the original file with `_test.go` suffix.
* Use mock interfaces from the project or `testify/mock` where applicable.

---

## 📜 Documentation Rules

* If you generate a new API, auto-generate the Swagger annotations in the handler.
* Update or create `.md` files inside `/devdocs/AI` with reasoning or traceability if AI-based logic is implemented.
* If modifying anything inside `/scripts/`, ensure it runs on Unix environments and contains proper shebangs.

---

## 🔁 Suggested Improvements (Optional for AI Agents)

These are ideas to propose or implement after team approval:

1. **Add AI changelog**: Maintain an `AI_CHANGELOG.md` with auto-generated summaries of what AI touched.
2. **Introduce validations registry**: Centralize validations in a package under `/pkg/validation/`.
3. **AI Test Plan Generator**: For any new API or service, automatically generate a stub QA test plan in `/devdocs/QA/`.
4. **Code ownership metadata**: Add a `.codeowners`-like registry or YAML to tag folders with responsible modules or maintainers for better traceability.
5. **Lint and Static Checks**: Propose or add `golangci-lint` configuration updates if style errors are frequent.
6. **Dev CLI Assistant**: Optionally create a script like `/scripts/dev-helper.sh` to assist devs in generating new services/routes using templates.

---

## 🧠 AI Context Expansion via `/devdocs/AI`

The `/devdocs/AI` folder acts as the **working memory and long-term knowledge base** for AI agents. This allows agents to avoid re-reading the entire codebase repeatedly and instead operate from structured, pre-digested information.

### Suggested Structure:

```
/devdocs/AI
├── OVERVIEW.md             # High-level summary of the entire repo
├── services/
│   └── erm-transfer-bank.md    # Summary of service logic, API flows, key models
├── flows/
│   ├── transaction_flow.md     # Key business logic flows (e.g., bank transfers)
│   └── auth_flow.md            # Auth/session logic
├── interfaces/
│   ├── services.md             # Service interface map
│   └── repositories.md         # Repo layer responsibilities
├── glossary.md                # Domain-specific terms (e.g., "transfer ledger", "debit lock")
├── api-cheatsheet.yaml        # Minimal OpenAPI or endpoint index for fast ref
├── prompts/
│   ├── refactor_guidelines.md  # How to do safe refactors
│   ├── doc_style.md            # Preferred doc style (tone, format)
│   └── test_template.md        # Testing patterns AI should follow
└── memory/
    ├── known_bugs.md           # Bugs AI should be aware of
    └── previous_tasks.json     # Past work logs with context
```

### Smart Practices:

* Use `/memory/previous_tasks.json` to track changes with rationale.
* Summarize code flows in `/flows/` rather than letting agents analyze raw files.
* Keep `/prompts/` folder to store reusable AI behavior instructions.
* Add frontmatter metadata (e.g., tags, updated timestamps) to markdown files to help agents query effectively.
* Maintain a central `knowledge.json` index that maps concepts to files.

AI agents must **update this folder** if they perform meaningful work, such as creating a new service, altering business logic, or fixing significant bugs.

