To-Do App Go REST CRUD
A RESTful API built in Go for managing a to-do list. It uses PostgreSQL for data storage and Redis for rate limiting.

Reference: https://roadmap.sh/projects/todo-list-api

Features:
✅ User authentication with JWT
👤 Admin + User roles
📋 Full CRUD operations on to-do tasks
🛢️ PostgreSQL database
🚦 Redis-based rate limiting
🔐 Secure API design
📁 Configurable via config.yaml

Prerequisites:

Install the following before getting started:
Go
PostgreSQL
Redis
Git

Setup Instructions:
1. Clone the Repository:
git clone https://github.com/sathish2000k/to-do-app-go-rest-crud.git
cd to-do-app-go-rest-crud

2. Install Go Modules:
go mod tidy

3. Configure config.yaml
In the root directory, create or update config.yaml:

server:
  port: 8080

database:
  host: localhost
  port: 5432
  user: your_db_user
  password: your_db_password
  name: todoapp
  sslmode: disable

redis:
  address: localhost:6379
  password: ""         # optional

Database Setup
Start PostgreSQL and create a database:

CREATE DATABASE todoapp;
Apply migrations if any are included or auto-generated.

Redis Setup
Make sure Redis is running locally:

redis-server

Running the Application

go run main.go
API will be available at http://localhost:8080

API Endpoints
🔐 Auth
POST /token – Authenticate and get JWT
POST /setPassword - Setting password after new user created
POST /resetPassword - Resetting password for a user

Admin



📝 To-Dos
GET /todos – List to-dos (JWT required)
POST /todos – Create a to-do
GET /todos/{id} – Get a specific to-do
PUT /todos/{id} – Update to-do
DELETE /todos/{id} – Delete to-do

🔒 Admin API Endpoints (/admin)
These routes are protected by:

JWT token
Admin role check (models.Admin)
Redis rate limiting

adminRouter := router.Group("/admin", 
  authMiddleWare.AuthMiddlewareToken(models.Admin), 
  authMiddleWare.RateLimitMiddleWare())
  
Admin Endpoints:
POST /admin/createUser – Create a new user
DELETE /admin/deleteUser/:id – Delete user by ID
GET /admin/getUser/:id – Fetch user details by ID
GET /admin/getAllUsers – Get list of all users

⚠️ All endpoints are protected with JWT and may be subject to rate limiting via Redis.

🧪 Testing
Use Postman or curl.

Example:
curl -H "Authorization: <admin_token>" http://localhost:8080/admin/getAllUsers

License
MIT
