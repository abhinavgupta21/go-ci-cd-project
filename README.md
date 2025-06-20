# Books API

A simple RESTful API for managing books, built with Go, Gin framework, and GORM ORM.

---

## Table of Contents

- [Overview](#overview)  
- [Features](#features)  
- [Project Structure](#project-structure)  
- [Environment Variables](#environment-variables)  
- [Installation](#installation)  
- [Usage](#usage)  
- [API Endpoints](#api-endpoints)  
- [Postman Collection](#postman-collection)
- [Continuous Integration and Release Workflows](#continuous-integration-and-release-workflows)
- [Contributing](#contributing)

---

## Overview

This project implements a basic CRUD API for books. Each book record contains the following attributes:

- **ID** (unique identifier)  
- **Title** (required)  
- **Author**  
- **PublishedYear**  
- **Price** (required)  

The API is built using the Gin web framework and uses GORM for ORM with database support.

---

## Features

- Create, Read, Update, and Delete (CRUD) operations for books  
- Validation for required fields (\`Title\` and \`Price\`)  
- Database schema migrations and updates with GORM  
- Graceful shutdown with proper closing of database connections  
- Environment-based configuration  

---

## Project Structure

```

├── .github/
│   └── workflows/
│       ├── ci.yml           # CI workflow for building Go project
│       └── release.yml      # Workflow for manual release deployment
├── config/
│   └── config.go            # Load and initialize environment variables
├── controllers/
│   └── book_controller.go   # Controller functions for CRUD operations
├── db/
│   └── connection.go        # DB connection and schema migration logic
├── models/
│   └── model.go             # Book model definition
├── routes/
│   └── routes.go            # API routing definitions
├── .env                     # Environment variable definitions
├── book-api.postman_collection.json  # Postman collection for testing API
└── main.go                  # Application entry point and server setup
```

---

## Environment Variables

The project uses environment variables for configuration, loaded from \`.env\` file:

| Variable         | Description                |
|------------------|----------------------------|
| \`DB_HOST\`      | Database host              |
| \`DB_PORT\`      | Database port              |
| \`DB_USER\`      | Database username          |
| \`DB_PASSWORD\`  | Database password          |
| \`DB_NAME\`      | Database name              |
| \`SERVER_PORT\`  | Port where the API runs    |

> **Note:** Add other variables as per your database and environment setup.

---

## Installation

1. **Clone the repository**

   ```
   git clone git@github.com:abhinavgupta21/go-ci-cd-project.git
   cd go-ci-cd-project
   ```

2. **Set up \`.env\` file**

   Created a `.env` file in the root folder with the necessary environment variables.

3. **Install dependencies**

   ```
   go mod tidy
   ```

4. **Run the API server**

   ```
   go run main.go
   ````

   The API server will start on the configured port (default from \`.env\`).

---

## Usage

- The API exposes REST endpoints to manage books.
- Use any API client (e.g., Postman) or \`curl\` to interact with the API.
- For convenience, a Postman collection (\`book-api.postman_collection.json\`) is provided.

---

## API Endpoints

| Method | Endpoint         | Description           |
|--------|------------------|-----------------------|
| POST   | \`/books\`       | Create a new book     |
| GET    | \`/books\`       | Get all books         |
| GET    | \`/books/:id\`   | Get a book by ID      |
| PUT    | \`/books/:id\`   | Update a book by ID   |
| DELETE | \`/books/:id\`   | Delete a book by ID   |

---

## Postman Collection

Import the `book-api.postman_collection.json` file in Postman to quickly test the API endpoints.

---

## Continuous Integration and Release Workflows

This project uses GitHub Actions workflows to automate build, test, lint, and release processes.

### CI Workflow (`.github/workflows/ci.yml`)

- **Triggered on:** Push to the `main` branch  
- **Runs:**
  - Checks out the code  
  - Sets up Go 1.24.3 environment  
  - Runs `golangci-lint` for code linting  
  - Builds the entire project  
  - Executes all Go tests  

This workflow ensures that every change pushed to the main branch is validated through linting, building, and testing, helping maintain code quality and stability.

---

### Release Workflow (`.github/workflows/release.yml`)

- **Triggered on:** Creation of a GitHub Release (manual release event)  
- **Runs:**
  - Checks out the code  
  - Sets up Go 1.24.3 environment  
  - Runs tests to verify code correctness  
  - Builds a Linux AMD64 binary named `booksApp`  
  - Uploads the binary as a release asset using [softprops/action-gh-release](https://github.com/softprops/action-gh-release)  

This workflow automates the packaging of this Go application binary and attaches it to your GitHub release, simplifying distribution for users.

---

### Notes

- The workflows use [golangci-lint](https://golangci-lint.run/) for comprehensive static analysis.
- The release workflow currently builds a Linux binary; you can extend this to build for other OS/architectures if needed.
- Make sure to create releases in your GitHub repository manually to trigger the release workflow.

---

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests for improvements.

---
