# GeoBill Golang Versions for Backend
<<<<<<< HEAD

GeoBill is a billing software website tailored for those launching hosting businesses. It aims to provide a cost-effective alternative to expensive commercial hosting software. GeoBill is built using Angular 20 for the frontend and Java Spring Boot 3 for the backend.

**Main Geobill Project Java Versions:** [https://github.com/alexistdev/geobill](https://github.com/alexistdev/geobill)
=======

GeoBill is a billing software website tailored for those launching hosting businesses. It aims to provide a cost-effective alternative to expensive commercial hosting software. GeoBill is built using Angular 20 for the frontend and Java Spring Boot 3 for the backend.


>>>>>>> 097380dc2591089b1df47422f8dce9b6e9cda434

## Prerequisites
- MySQL Database running at `localhost:3306`
- Database `mybill` exists.
- Table `users` exists. (If not, run the SQL below)

```sql
CREATE TABLE IF NOT EXISTS users (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

## How to Run

1.  Navigate to the project directory:
    ```powershell
    cd C:\Users\alexi\.gemini\antigravity\scratch\go-rbac-backend
    ```
2.  Run the application:
    ```powershell
    go run main.go
    ```

## Verification Steps

### 1. Registration
**Request:**
```bash
curl -X POST http://localhost:8080/v1/auth/registration \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "password123", "role": "ADMINISTRATOR"}'
```
**Expected Response:** `201 Created`

### 2. Login
**Request:**
```bash
curl -X POST http://localhost:8080/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "password123"}'
```
**Expected Response:** `200 OK`

### 3. Access Protected Route (Admin)
**Request:**
```bash
curl -v -u admin:password123 http://localhost:8080/v1/admin/dashboard
```
**Expected Response:** `200 OK` with JSON welcome message.

### 4. Access Denied (Wrong Role)
Create a user with role `USER`, then try to access `admin/dashboard`.
**Expected Response:** `403 Forbidden`
