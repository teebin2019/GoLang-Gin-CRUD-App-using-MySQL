# GoLang Gin CRUD App using MySQL

## Installation

```bash
go get github.com/pasan1/GoLang-Gin-CRUD-App-using-MySQL
```

## Usage

```bash
go run main.go
```

## API Endpoints

### Get All Users

```bash
GET /users
```

### Save User

```bash
POST /users
```

## Author

[Pasan Abeysekara](https://pasanabeysekara.com/)

## Structure

```bash
📦app
 ┣ 📂controllers
 ┃ ┣ 📜user_controller.go
 ┃ ┗ 📜vegetable_controller.go
 ┣ 📂models
 ┃ ┣ 📜user_model.go
 ┃ ┗ 📜vegetable_model.go
 ┣ 📂repositories
 ┃ ┣ 📜user_repository.go
 ┃ ┗ 📜vegetable_repository.go
 ┣ 📂routes
 ┃ ┗ 📜router.go
 ┗ 📜main.go
```

### Usage of the structure

1. `app/controllers`: This folder contains the controllers responsible for handling HTTP requests and defining the application's behavior. In this example, we have a `user_controller.go` file.
2. `app/models`: This folder contains the model definitions for your application. Here, we have a `user_model.go` file.
3. `app/repositories`: This folder contains the repository layer that interacts with the database. In this example, we have a `user_repository.go` file.
4. `app/routes`: This folder contains the router configuration for your application. Here, we have a `router.go` file.
5. `app/main.go`: This is the entry point of your application. It initializes the server and sets up the routes.
6. `config/database`.go: This file contains the configuration details for connecting to your MySQL database.
7. `go.mod` and `go.sum`: These files manage the Go modules and their dependencies.

## Contact

Email: [contact@pasanabeysekara.com](mailto:contact@pasanabeysekara.com)<br>
Website: [https://pasanabeysekara.com](https://pasanabeysekara.com)<br>
LinkedIn: [https://www.linkedin.com/in/pasanabeysekara](https://www.linkedin.com/in/pasanabeysekara)<br>
