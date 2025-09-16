# Go Microservice: Books, Authors & Quotes

This project is a **microservice written in Go**, following **Domain-Driven Design (DDD)** principles.  
The domain consists of three independent aggregates: `Book`, `Author`, and `Quote`.

The main idea is to keep each aggregate isolated and testable, allowing clients to consume the API in a flexible way.  
The design aims for simplicity and clarity so that any developer can easily identify the domain logic and the project layers.

---

## 🚀 Main Features
- Layered architecture (starting with the domain).
- Three main aggregates: **Book**, **Author**, **Quote**.
- Unit tests included for the domain layer.
- Ready to evolve with domain events.

---

## Basic Structure

```
├── internal/
│ └── domain/ # Aggregates, value objects, domain logic
├── docs/ # Documentation and decisions
│ └── decisions.md
└── README.md
```

---

## Dependencies
- [Go 1.25+](https://go.dev/dl/)  
- Main libraries:
  - `github.com/google/uuid` (value objects / IDs)
  - `github.com/brianvoe/gofakeit/v7` (testing)
  - `github.com/stretchr/testify` (testing)

---

## Documentation
- [Architecture Decisions](./docs/decisions.md)

---

## How to Run the Project

### Run tests
```bash
make local-test-unit
```