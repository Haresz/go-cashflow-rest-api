# 💰 Cashflow Backend API - Go & Gin

A REST API for financial transaction management built with Go, Gin framework, and SQLite.

## 📋 Table of Contents

- [About](#about)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [API Documentation](#api-documentation)
- [Database Schema](#database-schema)
- [Project Structure](#project-structure)
- [Configuration](#configuration)
- [Troubleshooting](#troubleshooting)

---

## About

Cashflow Backend API is a REST API for managing financial transactions. It provides a complete CRUD interface for tracking income and expenses with support for filtering, search, and comprehensive error handling.

---

## Features

- ✅ **RESTful API**: Full CRUD operations with standard HTTP methods
- ✅ **SQLite Database**: Lightweight, serverless database with auto-initialization
- ✅ **Query Filtering**: Filter transactions by type, category, and date range
- ✅ **CORS Support**: Configurable cross-origin resource sharing for frontend integration
- ✅ **Input Validation**: Basic validation for required fields and transaction types
- ✅ **Structured Responses**: Consistent JSON response format with proper status codes
- ✅ **Single Binary Deployment**: Compiled to a single executable without dependencies
- ✅ **Comprehensive Logging**: INFO and ERROR level logging for all operations
- ✅ **Error Handling**: Explicit error checking with meaningful error messages

---

## Tech Stack

| Component | Technology | Version | Description |
|-----------|------------|---------|-------------|
| **Language** | Go | 1.25+ | Backend programming language |
| **Web Framework** | Gin | 1.12.0 | HTTP router and middleware |
| **Database** | SQLite | 3.x | Relational database storage |
| **HTTP Library** | net/http | Standard | HTTP server/client |
| **CORS** | gin-contrib/cors | 1.7.7 | Cross-origin request handling |

---

## Prerequisites

- **Go 1.25 or higher**

### Check Go Version
```bash
go version
```

Expected output:
```
go version go1.25.6 linux/amd64
```

---

## Installation

### Step 1: Navigate to Project Directory
```bash
cd /path/to/go-cahsflow-rest-api
```

### Step 2: Install Dependencies
```bash
go mod download
```

This downloads all required dependencies listed in `go.mod`.

### Step 3: Build Application
```bash
go build -o cashflow-backend .
```

This compiles the Go code into a single executable binary.

**Output:**
- `cashflow-backend.exe` (Windows)
- `cashflow-backend` (Linux/macOS)

### Step 4: Run Server

**Option 1: Run from Binary**
```bash
# Windows
./cashflow-backend.exe

# Linux/macOS
./cashflow-backend
```

**Option 2: Run Without Compiling (Development)**
```bash
go run .
```

**Option 3: Run in Background (Linux/macOS)**
```bash
./cashflow-backend &
```

### Step 5: Verify Server

```bash
curl http://localhost:8080/ping
```

Expected response:
```json
{
  "message": "pong"
}
```

---

## Usage

### Running the Server

Start the server:
```bash
./cashflow-backend
```

The server will start on `http://localhost:8080` by default.

### Stopping the Server

Press `Ctrl+C` to stop the server if running in the foreground, or use:
```bash
pkill cashflow-backend
```

### Environment Variables

The application can be configured using the following approaches:

1. **Direct Configuration**: Edit values in `main.go`
2. **Environment Variables** (not yet implemented, but can be added)

---

## API Documentation

### Base URL
```
http://localhost:8080
```

### Response Format

**Success Response:**
```json
{
  "success": true,
  "data": { ... }
}
```

**Error Response:**
```json
{
  "success": false,
  "error": "Error message here"
}
```

### Endpoints

#### 1. Health Check

**GET /ping**

Check if the server is running.

**Response:**
```json
{
  "message": "pong"
}
```

**Status Codes:**
- `200 OK`: Server is running

---

#### 2. List All Transactions

**GET /transactions**

Retrieve all transactions with optional filtering.

**Query Parameters:**

| Parameter | Type | Description | Example |
|-----------|------|-------------|---------|
| `jenis` | string | Filter by transaction type | `?jenis=Pemasukan` |
| `kategori` | string | Filter by category | `?kategori=Gaji` |
| `tanggal` | string | Filter by exact date (YYYY-MM-DD) | `?tanggal=2026-04-06` |
| `startDate` | string | Filter by start date (YYYY-MM-DD) | `?startDate=2026-01-01` |
| `endDate` | string | Filter by end date (YYYY-MM-DD) | `?endDate=2026-12-31` |

**Examples:**

Get all transactions:
```bash
curl http://localhost:8080/transactions
```

Filter by transaction type:
```bash
curl "http://localhost:8080/transactions?jenis=Pemasukan"
```

Filter by date range:
```bash
curl "http://localhost:8080/transactions?startDate=2026-01-01&endDate=2026-12-31"
```

Multiple filters:
```bash
curl "http://localhost:8080/transactions?jenis=Pemasukan&kategori=Gaji"
```

**Response:**
```json
{
  "success": true,
  "data": {
    "transactions": [
      {
        "id": 1,
        "tanggal": "2026-04-06",
        "jenis": "Pemasukan",
        "kategori": "Gaji",
        "nominal": 5000000,
        "keterangan": "Monthly salary"
      },
      {
        "id": 2,
        "tanggal": "2026-04-07",
        "jenis": "Pengeluaran",
        "kategori": "Makan",
        "nominal": 50000,
        "keterangan": "Lunch"
      }
    ]
  }
}
```

**Status Codes:**
- `200 OK`: Successfully retrieved transactions
- `500 Internal Server Error`: Database error

---

#### 3. Get Transaction by ID

**GET /transactions/:id**

Retrieve a single transaction by its ID.

**Example:**
```bash
curl http://localhost:8080/transactions/1
```

**Response:**
```json
{
  "success": true,
  "data": {
    "transaction": {
      "id": 1,
      "tanggal": "2026-04-06",
      "jenis": "Pemasukan",
      "kategori": "Gaji",
      "nominal": 5000000,
      "keterangan": "Monthly salary"
    }
  }
}
```

**Status Codes:**
- `200 OK`: Successfully retrieved transaction
- `404 Not Found`: Transaction with specified ID does not exist
- `400 Bad Request`: Invalid transaction ID format
- `500 Internal Server Error`: Database error

---

#### 4. Create Transaction

**POST /transactions**

Create a new transaction.

**Request Body:**
```json
{
  "tanggal": "2026-04-06",
  "jenis": "Pemasukan",
  "kategori": "Gaji",
  "nominal": 5000000,
  "keterangan": "Monthly salary"
}
```

**Validation Rules:**
- `tanggal` (required): Transaction date (YYYY-MM-DD format)
- `jenis` (required): Must be "Pemasukan" or "Pengeluaran"
- `kategori` (required): Category of the transaction
- `nominal` (required): Transaction amount (integer)
- `keterangan` (optional): Additional description

**Example:**
```bash
curl -X POST http://localhost:8080/transactions \
  -H "Content-Type: application/json" \
  -d '{
    "tanggal": "2026-04-06",
    "jenis": "Pemasukan",
    "kategori": "Gaji",
    "nominal": 5000000,
    "keterangan": "Monthly salary"
  }'
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1
  }
}
```

**Status Codes:**
- `201 Created`: Transaction successfully created
- `400 Bad Request`: Validation error or invalid request body
- `500 Internal Server Error`: Database error

---

#### 5. Update Transaction

**PUT /transactions/:id**

Update an existing transaction.

**Request Body:**
```json
{
  "tanggal": "2026-04-06",
  "jenis": "Pemasukan",
  "kategori": "Gaji",
  "nominal": 6000000,
  "keterangan": "Monthly salary updated"
}
```

**Example:**
```bash
curl -X PUT http://localhost:8080/transactions/1 \
  -H "Content-Type: application/json" \
  -d '{
    "tanggal": "2026-04-06",
    "jenis": "Pemasukan",
    "kategori": "Gaji",
    "nominal": 6000000,
    "keterangan": "Monthly salary updated"
  }'
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1
  }
}
```

**Status Codes:**
- `200 OK`: Transaction successfully updated
- `400 Bad Request`: Validation error or invalid request body
- `404 Not Found`: Transaction with specified ID does not exist
- `500 Internal Server Error`: Database error

---

#### 6. Delete Transaction

**DELETE /transactions/:id**

Delete a transaction by its ID.

**Example:**
```bash
curl -X DELETE http://localhost:8080/transactions/1
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1
  }
}
```

**Status Codes:**
- `200 OK`: Transaction successfully deleted
- `404 Not Found`: Transaction with specified ID does not exist
- `400 Bad Request`: Invalid transaction ID format
- `500 Internal Server Error`: Database error

---

## Database Schema

### Table: transactions

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| `id` | INTEGER | PRIMARY KEY AUTOINCREMENT | Unique identifier |
| `tanggal` | TEXT | NOT NULL | Transaction date |
| `jenis` | TEXT | NOT NULL | Transaction type (Pemasukan/Pengeluaran) |
| `kategori` | TEXT | NOT NULL | Transaction category |
| `nominal` | INTEGER | NOT NULL | Transaction amount |
| `keterangan` | TEXT | - | Additional description |

**SQL Schema:**
```sql
CREATE TABLE IF NOT EXISTS transactions (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  tanggal TEXT NOT NULL,
  jenis TEXT NOT NULL,
  kategori TEXT NOT NULL,
  nominal INTEGER NOT NULL,
  keterangan TEXT
);
```

### Field Mapping

| JSON Key | Go Field | Go Type | Description |
|----------|----------|---------|-------------|
| `id` | `ID` | `int` | Primary key (auto-increment) |
| `tanggal` | `Tanggal` | `string` | Transaction date (YYYY-MM-DD) |
| `jenis` | `Jenis` | `string` | Transaction type (Pemasukan/Pengeluaran) |
| `kategori` | `Kategori` | `string` | Transaction category |
| `nominal` | `Nominal` | `int` | Transaction amount |
| `keterangan` | `Keterangan` | `string` | Additional description |

---

## Project Structure

```
go-cahsflow-rest-api/
├── main.go                  # Application entry point and main logic
├── go.mod                   # Module dependencies
├── go.sum                   # Dependency checksums and versions
├── cashflow.db              # SQLite database file (auto-created)
├── cashflow-backend.exe     # Compiled binary (Windows)
├── .gitignore               # Git ignore rules
└── README.md                # This file
```

### File Descriptions

| File | Description |
|------|-------------|
| `main.go` | Entry point with all application logic including handlers, database operations, and server setup |
| `go.mod` | Module definition and dependency list (similar to `package.json`) |
| `go.sum` | Dependency checksums for security (similar to `package-lock.json`) |
| `cashflow.db` | SQLite database file (created automatically on first run) |
| `cashflow-backend.exe` | Compiled Windows executable (generated by `go build`) |
| `.gitignore` | Git ignore patterns for binaries, databases, and temporary files |
| `README.md` | Project documentation |

---

## Configuration

### Port Configuration

Default port: `8080`

To change the port, edit `main.go`:

```go
// Find this line in main.go
if err := r.Run(":8080"); err != nil {

// Change to desired port
if err := r.Run(":3000"); err != nil {
```

### Database Configuration

Default database path: `./cashflow.db`

To change the database location, edit `main.go`:

```go
// Find this line in main.go
db, err := sql.Open("sqlite", "./cashflow.db")

// Change to desired path
db, err := sql.Open("sqlite", "./data/cashflow.db")
```

### CORS Configuration

Default allowed origin: `http://localhost:5173`

To configure allowed origins, edit `main.go`:

```go
// Find the CORS configuration
r.Use(cors.New(cors.Config{
    AllowOrigins: []string{"http://localhost:5173"},
    // ... other config
}))

// Add more origins
AllowOrigins: []string{"http://localhost:5173", "https://yourdomain.com"},
```

---

## Troubleshooting

### Error: Port Already in Use

**Problem:**
```
listen tcp :8080: bind: address already in use
```

**Solution:**
```bash
# Find process using port 8080
netstat -ano | findstr :8080  # Windows
lsof -ti:8080                  # Linux/macOS

# Kill the process
taskkill /PID <PID> /F         # Windows
kill -9 <PID>                  # Linux/macOS
```

### Error: Module Not Found

**Problem:**
```
module github.com/gin-gonic/gin: found in module cache but not in go.mod
```

**Solution:**
```bash
# Clean module cache
go clean -modcache

# Re-download dependencies
go mod download

# Update go.mod and go.sum
go mod tidy
```

### Error: Connection Refused

**Problem:**
```
curl: (7) Failed to connect to localhost port 8080: Connection refused
```

**Solution:**
```bash
# Verify server is running
ps aux | grep cashflow-backend

# Restart server
./cashflow-backend

# Check if port is correct
curl http://localhost:8080/ping
```

### Error: CORS Issues

**Problem:** Frontend cannot access the API

**Solution:** Verify CORS configuration in `main.go` and ensure your frontend origin is in `AllowOrigins`:

```go
r.Use(cors.New(cors.Config{
    AllowOrigins: []string{"http://localhost:5173", "http://localhost:3000"},
}))
```

### Error: Database Locked

**Problem:**
```
database is locked
```

**Solution:**
```bash
# Close all connections
pkill cashflow-backend

# Check if another process is holding the database
fuser cashflow.db  # Linux
```

---

## Development

### Useful Go Commands

```bash
# Module management
go mod init <module-name>    # Initialize module
go mod tidy                  # Clean dependencies
go mod download              # Download dependencies
go mod verify                # Verify dependencies

# Building
go build                     # Build executable
go build -o app.exe .        # Build with custom name
go run .                     # Run without building

# Testing
go test ./...                # Run all tests
go test -v ./...             # Verbose test output

# Formatting
go fmt ./...                 # Format code
go vet ./...                 # Check for common errors

# Environment
go env                       # Show Go environment
go version                   # Show Go version
```

---

## API Status

- ✅ Server running on `localhost:8080`
- ✅ SQLite database auto-initialized
- ✅ CORS middleware configured
- ✅ `GET /ping` - Health check endpoint
- ✅ `GET /transactions` - List all transactions with filters
- ✅ `GET /transactions/:id` - Get transaction by ID
- ✅ `POST /transactions` - Create transaction
- ✅ `PUT /transactions/:id` - Update transaction
- ✅ `DELETE /transactions/:id` - Delete transaction
- ✅ Input validation for all operations
- ✅ Comprehensive logging for all operations

---

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

## License

This project is licensed under the MIT License.

---

**Happy Coding! 🚀**
