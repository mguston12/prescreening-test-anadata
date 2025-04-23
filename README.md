# Screening Test - CRUD API with JWT (Go + PostgreSQL)

## 🛠️ Installation

### 1. Clone Repository

```bash
git clone (https://github.com/mguston12/prescreening-test-anadata.git)
cd go-project
go mod init
go mod tidy
```

### 2. Buat .env file
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=cruddb
DB_PORT=5432
DB_HOST=db
JWT_SECRET=your_jwt_secret

### 3. Run with Docker Compose 
```bash
docker-compose up --build
```

### 4. Buka Terminal lalu jalankan program
```bash
go run cmd/main.go
```

### 5. Program Berhasil dijalankan


Testing API Menggunakan POSTMAN

BASEURL = http://localhost:8080

1. Register (/register)
URL: http://localhost:8080/register
Method: POST
Body : JSON
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "secret123"
}


2. Login (/login)
URL: http://localhost:8080/login
Method: POST
Body : JSON 
{
  "email": "john@example.com",
  "password": "secret123"
}

API Login akan mengembalikan token yang akan digunakan untuk API lainnya yang digunakan untuk header Authorization <JWT_TOKEN>

3. GetAllUsers
URL: http://localhost:8080/api/users
Method: GET
Headers:
Authorization: Bearer <JWT_TOKEN>

4. GetUserByID
URL: http://localhost:8080/api/users/{id}
Method: GET
Headers:
Authorization: Bearer <JWT_TOKEN>

5. UpdateUser 
URL: http://localhost:8080/api/users/{id}
Method: PUT
Headers:
Authorization: Bearer <JWT_TOKEN>
Body : JSON
{
  "name": "John New",
  "email": "johnnew@example.com"
}

6. DeleteUser
URL: http://localhost:8080/api/users/{id}
Method: DELETE
Headers:
Authorization: Bearer <JWT_TOKEN>
