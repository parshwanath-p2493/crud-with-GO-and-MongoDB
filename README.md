# Go CRUD API with Gin & MongoDB

This project is a simple CRUD API built using the **Gin** framework and **MongoDB** as the database. It provides endpoints to **create, read, update, and delete** data.

---

## Features
- RESTful API with **Gin**
- **MongoDB** for data storage
- CRUD operations (**Create, Read, Update, Delete**)
- Well-structured modular codebase

---

## Prerequisites
Ensure you have the following installed:
- **Go** (https://go.dev/doc/install)
- **MongoDB** (https://www.mongodb.com/try/download/community)

---

## Installation & Setup
### 1. Clone the Repository
```sh
git clone https://github.com/your-username/go-crud.git
cd go-crud
```

### 2. Initialize Go Module
```sh
go mod init go-crud
```

### 3. Install Dependencies
```sh
go get -u github.com/gin-gonic/gin
go get go.mongodb.org/mongo-driver
```

### 4. Run MongoDB
Ensure MongoDB is running on `localhost:27017`.

For Linux/macOS:
```sh
mongod --dbpath /path/to/mongo/data
```
For Windows (using PowerShell):
```sh
mongod.exe --dbpath C:\path\to\mongo\data
```

---

## Project Structure
```
/go-crud
├── main.go          # Server setup & routes
├── database.go      # MongoDB connection
├── handlers.go      # CRUD logic
├── models.go        # Data model definition
└── go.mod           # Go module dependencies
```

---

## CRUD Endpoints

### **1. Create Data**
**Endpoint:** `POST /data`

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com"
}
```
**Response:**
```json
{
  "message": "Data inserted",
  "data": {
    "id": "65d4b6...",
    "name": "John Doe",
    "email": "john@example.com"
  }
}
```

---

### **2. Read All Data**
**Endpoint:** `GET /data`

**Response:**
```json
{
  "message": "Data retrieved",
  "data": [
    {
      "id": "65d4b6...",
      "name": "John Doe",
      "email": "john@example.com"
    }
  ]
}
```

---

### **3. Update Data (Fixed URL Param Issue)**
**Endpoint:** `PUT /data/:id`

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
  "message": "Data updated"
}
```

---

### **4. Delete Data (Fixed URL Param Issue)**
**Endpoint:** `DELETE /data/:id`

**Response:**
```json
{
  "message": "Data deleted"
}
```

---

## Running the Application
```sh
go run main.go
```
The server will start at: **`http://localhost:9090`**

---

## Fix for PUT & DELETE Issues
The original code used `c.Query("id")`, which did not work well for `PUT` and `DELETE` requests. The issue was resolved by extracting the `id` from the URL path using `c.Param("id")`.

Updated `main.go`:
```go
r.PUT("/data/:id", UpdateData)
r.DELETE("/data/:id", DeleteData)
```

Updated `handlers.go`:
```go
id := c.Param("id")
objectId, err := primitive.ObjectIDFromHex(id)
if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
    return
}
```

This ensures the correct `id` is extracted and converted properly.

---


