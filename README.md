# golang-nutrition-server

This project is a Go-based web application using the Gin framework. It provides a RESTful API for managing products and their associated allergies.

## Features

- CRUD operations for products.
- Adding and removing allergies associated with products.
- Endpoints for user registration and authentication.

## Prerequisites

To run this application, you'll need:

- Go (version 1.16 or later)
- PostgreSQL
- Docker (optional, for containerization)

## Installation

Clone the repository:

```bash
git clone https://github.com/yourusername/your-repo-name.git
cd your-repo-name
```

## Running

To run the application locally:

```bash
go run main.go
```
The application will start and listen for requests on localhost:8080.

Running with Docker

To build and run the application using Docker:

```bash
docker build -t your-app-name .
docker run -p 8080:8080 your-app-name
```

## API Endpoints

The application provides the following RESTful endpoints:

    POST /products: Add a new product.
    GET /products/{productID}: Retrieve a product by ID.
    PUT /products: Update a product.
    DELETE /products/{productID}: Delete a product.
    POST /register: Register a new user.
    POST /login: Authenticate a user.

## Testing

Run the automated tests for this system:

```bash
go test ./...
```

