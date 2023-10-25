# CHI Rest API :computer:

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/breno5g/chi-rest-api/blob/main/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/breno5g/chi-rest-api)](https://github.com/breno5g/chi-rest-api/issues)
[![GitHub stars](https://img.shields.io/github/stars/breno5g/chi-rest-api)](https://github.com/breno5g/chi-rest-api/stargazers)
![Golang](https://img.shields.io/badge/Go-1.21.1-blue)
![SQLite](https://img.shields.io/badge/SQLite-3.36.0-blue)
![Go-Chi](https://img.shields.io/badge/Go--Chi-4.1.2-blue)
![Swagger](https://img.shields.io/badge/Swagger-2.0-blue)

This is a simple RESTful API project developed using Golang :rocket:, Swagger :bookmark_tabs:, Chi :construction_worker:, and SQLite :file_folder:. The purpose of this API is to create and manage users :busts_in_silhouette: and products :package:. It provides endpoints for creating, retrieving, updating, and deleting users and products.

## 📂 Project Structure

```css
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

- **cmd/server/**: Contains the main application file for your server.

- **configs/**: Houses configuration files or settings for your Go application.

- **docs/**: May include documentation files or assets related to your project.

- **internal/**: Organizes your internal code, with separate folders for DTO, entities, and infrastructure.

- **pkg/entity/**: Contains code related to entities used in your application.

- **pkg/infra/**: Infrastructure code with subdirectories for database and webserver.

- **test/**: May include test files or assets for your project.

## 🛠️ Prerequisites

Before running this project, you'll need to have Golang and SQLite installed on your system.

## 🚀 Running the Project

To run the project:

1. Clone this repository to your local machine.

2. Open your terminal and navigate to the project's root directory:

   ```bash
   cd chi-rest-api
   ```

3. Run the following command to start the application:

   ```bash
   go run ./cmd/server/main.go
   ```

4. The API will be available at `http://localhost:3001`.

## Usage :computer:

You can use tools like [curl](https://curl.se/) or [Postman](https://www.postman.com/) to interact with the API. The API provides endpoints for creating, retrieving, updating, and deleting users and products.

## Swagger Documentation :book:

The API is documented using Swagger, providing comprehensive details about endpoints, request parameters, and response formats. Access the Swagger documentation at `http://localhost:3001/swagger/index.html` when the API is running.

## Database :file_folder:

This project uses SQLite as the database for simplicity. You can find the database file in the `data` directory.

## 🤝 Contributing

If you wish to contribute to this project, please follow these guidelines:

1. Fork the repository.

2. Create a new branch for your feature or bug fix.

3. Implement your changes.

4. Test your changes thoroughly.

5. Create a pull request with a clear description of your changes.

## 📄 License

This project is licensed under the [MIT License](LICENSE). Feel free to use and modify it according to your needs.

Enjoy using this CHI Rest API with Golang, Swagger, Go-Chi, and SQLite for managing users and products! :tada:

You can further customize this README with your specific details and make any additional changes as needed.
