## Table of Contents

- [Requirements](#requirements)
- [Instalation](#instalation)



## Requirements

- Go version 1.15 or higher
- PostgreSQL (13 or higher)


## Installation
1. Create database 'privy'
   ```bash
   create database privy;
   ```
2. Create table 'employees'
   ```bash
   create table employees(id serial primary key, name varchar(255), age integer, salary integer);
   ```
   
3. Install the dependencies:
    ```bash
    go get
    go mod tidy
    ```

4. Configure your database connection APP_DB_USERNAME, APP_DB_PASSWORD, APP_DB_NAME in main.go file

5. Run the application:
    ```bash
    go run main.go
    ```
6. Run in postman with method GET
  ``` http://localhost:8080/employees ```

