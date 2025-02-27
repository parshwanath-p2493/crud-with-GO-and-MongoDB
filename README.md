# Go CRUD API with Gin and MongoDB

This is a simple CRUD API built using the **Gin** framework and **MongoDB** in **Go**. It allows users to perform Create, Read, Update, and Delete operations on a MongoDB collection.

## Features
- **CRUD operations** (Create, Read, Update, Delete)
- **MongoDB Integration**
- **Gin Framework**
- **Swagger Documentation**

---

## Installation and Setup

### 1. Install Go Modules
Ensure that **Go** is installed on your system. Then, initialize the project:
```sh
mkdir go-crud
cd go-crud
go mod init go-crud
```

### 2. Install Required Packages
```sh
go get -u github.com/gin-gonic/gin
go get go.mongodb.org/mongo-driver
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

### 3. Start MongoDB
Ensure MongoDB is running locally on **port 27017** or update the connection string accordingly in `database.go`.

---

## Project Structure
```
/go-crud
│── docs/                 # Swagger documentation
│── database.go           # MongoDB connection
│── handlers.go           # CRUD operations
│── models.go             # Data model
│── main.go               # Server setup
│── go.mod                # Go module dependencies
│── README.md             # Project documentation
```

---

## Running the API

### 1. Start the API Server
```sh
go run main.go
```
The server will start at: `http://localhost:9090`

### 2. Access Swagger Documentation
Once the server is running, open:
```
http://localhost:9090/swagger/index.html
```

---

## API Endpoints

### **1. Get All Data**
**Endpoint:** `GET /data`

**Response:**
```json
{
  "message": "Data retrieved",
  "data": [
    {
      "id": "61a8c3a0c9a4d7eb6c4a4d6d",
      "name": "John Doe",
      "email": "john@example.com"
    }
  ]
}
```

### **2. Add Data**
**Endpoint:** `POST /data`

**Request Body:**
```json
{
  "name": "Jane Doe",
  "email": "jane@example.com"
}
```

**Response:**
```json
{
  "message": "Data inserted",
  "data": {
    "id": "61a8c3a0c9a4d7eb6c4a4d6e",
    "name": "Jane Doe",
    "email": "jane@example.com"
  }
}
```

### **3. Update Data**
**Endpoint:** `PUT /data?id={id}`

**Request Body:**
```json
{
  "name": "Jane Updated",
  "email": "jane.updated@example.com"
}
```

**Response:**
```json
{
  "message": "Data updated"
}
```

### **4. Delete Data**
**Endpoint:** `DELETE /data?id={id}`

**Response:**
```json
{
  "message": "Data deleted"
}
```

---

## Swagger Integration

### 1. Generate Swagger Documentation
Run the following command to generate the Swagger docs:
```sh
swag init --output docs
```

### 2. Import Swagger in `main.go`
```go
import (
    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
    _ "./docs"
)
```

### 3. Add Swagger Route
```go
r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
```

Now, you can access Swagger UI at:
```
http://localhost:9090/swagger/index.html
```

---

## Contributing
Feel free to contribute to this project by submitting issues or pull requests.

---

## License
This project is open-source and available under the **MIT License**.

