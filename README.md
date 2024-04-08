# Go Bookstore

This project is a simple RESTful API for a bookstore application built with Go. It uses GORM for ORM functionality and Gorilla Mux for routing. The API allows for managing books in the bookstore, supporting operations such as creating, retrieving, updating, and deleting books.

## Features

- List all books
- Get details of a specific book
- Create a new book
- Update an existing book
- Delete a book

## Prerequisites

Before you begin, ensure you have met the following requirements:
- Go (version 1.16 or later)
- MySQL

## Setting Up

### Clone the Repository

First, clone the repository to your local machine:

git clone https://github.com/Vishak-V/golang-with-mysql.git
cd go-bookstore


### Database Configuration

Create a MySQL database for the application and note the credentials. You'll need to update the connection string in `config/app.go` with your database details:

```go
// Example connection string (update with your details)
dsn := "root:password@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
```
## Install Dependencies

Run the following command to install the required Go modules:
```
go mod tidy
```

## Running the Application

To start the server, run:
```
go run main.go
```

# Go Bookstore API

## API Endpoints

The Go Bookstore API provides the following RESTful endpoints:

- **List All Books**
  - `GET /books`
  - Retrieves a list of all books in the bookstore.

- **Get Book Details**
  - `GET /book/{id}`
  - Retrieves details of a specific book by its ID.

- **Create a New Book**
  - `POST /book`
  - Creates a new book with the provided book details in the request body.

- **Update an Existing Book**
  - `PUT /book/{id}`
  - Updates the details of an existing book. The updated book details should be provided in the request body.

- **Delete a Book**
  - `DELETE /book/{id}`
  - Deletes a book from the bookstore by its ID.

## Contributing

We welcome contributions to the Go Bookstore project! If you have suggestions for improvement or want to contribute code, please:

1. Fork the repository.
2. Create a new branch for your feature or fix.
3. Commit your changes.
4. Push to your branch.
5. Open a pull request.

For major changes or if you have any questions, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License - see the `LICENSE` file for details.
