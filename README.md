# Prakerja_Golang

Prakerja_Golang is a repository containing a Golang application that serves as an example project for the Prakerja program. It demonstrates how to fetch employee data from a database and provide a REST API for retrieving, creating, updating, and deleting employee information. Additionally, the application includes authentication using JWT for admin users.

# Features
- Fetch and display all employee data from the database.
- Create, update, and delete employee data using REST API endpoints.
- Secure API endpoints with authentication using JWT for admin users.

# Installation
To get this project up and running on your local machine, follow these steps:
1. Clone the repository:

```sql
git clone https://github.com/rndanugerah/Prakerja_Golang.git
cd Prakerja_Golang
```

2. Install the dependencies:

Make sure you have Golang installed on your machine. Then, run:
```sql
go mod tidy
```

3. Set up the database:

This project requires a database to store employee data. Make sure you have a compatible database installed ( MySQL) and configure the connection details in the db.CreateCon() function.

4. Run the application:
```sql
go run main.go
```

# Usage
Once the application is up and running, you can access the API endpoints using API testing tools like Postman.

# API Endpoints
- GET /employees: Retrieve a list of all employees.
- GET /employees/{id}: Retrieve information about a specific employee by their ID.
- POST /employees: Create a new employee. (Requires admin authentication)
- PUT /employees/{id}: Update information for a specific employee by their ID. (Requires admin authentication)
- DELETE /employees/{id}: Delete an employee by their ID. (Requires admin authentication)
- GET /generate-hash: generate a secure hash value for a given input data, making it suitable for data integrity verification and password hashing purposes.
- POST /login: Authenticate admin user and obtain a JWT token.

# Database

This application uses a relational database to store employee data. Ensure that you have a compatible database installed (MySQL) and configure the connection details in the db.CreateCon() function.

# Authentication
The application supports authentication for admin users using JSON Web Tokens (JWT). The /login endpoint is used to obtain a JWT token by providing valid admin credentials. The generated token should be included in the Authorization header for all endpoints requiring admin access.

# Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request.

1. Fork the repository.
2. Create a new branch: git checkout -b my-feature-branch.
3. Make your changes and commit them: git commit -m 'Add some feature'.
4. Push to the branch: git push origin my-feature-branch.
5. Submit a pull request

