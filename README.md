# CHI Rest API :computer:

[![Golang](https://img.shields.io/badge/Golang-1.16%2B-blue.svg)](https://golang.org/)
[![SQLite](https://img.shields.io/badge/SQLite-3.x-blue.svg)](https://sqlite.org/)
[![Go-Chi](https://img.shields.io/badge/Go--Chi-4.1.2%2B-blue.svg)](https://pkg.go.dev/github.com/go-chi/chi)
[![Swagger](https://img.shields.io/badge/Swagger-2.0-blue.svg)](https://swagger.io/)

A simple RESTful API project developed using Golang :rocket:, Swagger :bookmark_tabs:, Chi :construction_worker:, and SQLite :file_folder:. This API serves the purpose of managing users :busts_in_silhouette: and products :package:. It provides endpoints for creating, retrieving, updating, and deleting users and products.

...

## Project Structure :open_file_folder:

Here is the folder structure of the project:

```
chi-rest-api/
├─ cmd/
│  ├─ server/
│  │  ├─ main.go
├─ configs/
│  ├─ configs.go
├─ docs/
├─ internal/
│  ├─ dto/
│  ├─ entity/
│  ├─ infra/
│  │  ├─ database/
│  │  ├─ webserver/
├─ pkg/
│  ├─ entity/
├─ test/
```

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Swagger Documentation](#swagger-documentation)
- [Database](#database)
- [Run the Project](#run-the-project)
- [Contributing](#contributing)
- [License](#license)

## Installation :wrench:

To get started with this project, follow these installation steps:

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/chi-rest-api.git
   cd chi-rest-api
   ```

2. **Install Dependencies:**

   ```bash
   go mod download
   ```

## Usage :computer:

You can use tools like [curl](https://curl.se/) or [Postman](https://www.postman.com/) to interact with the API. The API provides endpoints for managing users and products.

## Swagger Documentation :book:

This API is documented using Swagger, providing comprehensive details about endpoints, request parameters, and response formats. Access the Swagger documentation by navigating to:

```
http://localhost:8080/swagger/index.html
```

Make sure the API is running to access this documentation.

## Database :file_folder:

The project utilizes SQLite as the database :file_folder:, with the database file located in the `data` directory.

## Run the Project :running:

To run the project, use the following command:

```bash
make run
# or
go run ./cmd/server/main.go
```

The API server will start, and you can access it at `http://localhost:8080`. You can modify the port in the `main.go` file if needed.

## Contributing :hammer_and_wrench:

If you wish to contribute to this project, please follow these guidelines:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Implement your changes.
4. Test your changes thoroughly.
5. Create a pull request with a clear description of your changes.

## License :page_with_curl:

This project is licensed under the [MIT License](LICENSE). Feel free to use, modify, and distribute it according to your needs.

---

Feel free to further customize this README with emojis or additional information about your project. Good luck with your "CHI Rest API" project! :tada:
