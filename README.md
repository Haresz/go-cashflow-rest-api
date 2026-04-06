# 💰 Cashflow Backend API - Go & Gin

REST API lengkap untuk monitoring keuangan menggunakan Go, framework Gin, dan database SQLite.

## 📋 Table of Contents

- [📖 Deskripsi Proyek](#-deskripsi-proyek)
- [👨‍💻 Panduan untuk Pengembang JS](#-panduan-untuk-pengembang-js)
- [🛠️ Prasyarat](#️-prasyarat)
- [🚀 Cara Menjalankan](#-cara-menjalankan)
- [📡 Dokumentasi API/Fungsi](#-dokumentasi-apifungsi)
- [⚙️ Penjelasan Konsep Go](#️-penjelasan-konsep-go)
- [📁 Struktur Folder](#-struktur-folder)
- [🐛 Troubleshooting](-troubleshooting)
- [📚 Lanjutkan Belajar](-lanjutkan-belajar)

---

## 📖 Deskripsi Proyek

**Cashflow Backend API** adalah backend REST API untuk monitoring keuangan (cashflow) yang dibangun dengan bahasa Go. Proyek ini dirancang untuk membantu developer JavaScript bertransisi ke Go sambil membangun aplikasi backend yang fungsional.

### 🎯 Fungsi Teknis

- ✅ **REST API Server**: HTTP server yang melayani request/response dengan format JSON
- ✅ **Database SQLite**: Database relasional ringan untuk menyimpan data transaksi keuangan
- ✅ **CORS Middleware**: Cross-Origin Resource Sharing untuk integrasi frontend
- ✅ **Auto Database Creation**: Database dan tabel dibuat otomatis jika belum ada
- ✅ **Error Handling**: Pattern error handling yang eksplisit dan predictable

### 🛠️ Tech Stack

| Komponen | Teknologi | Versi | Fungsi |
|----------|----------|-------|--------|
| **Bahasa** | Go | 1.25+ | Bahasa pemrograman backend |
| **Web Framework** | Gin | 1.12.0 | HTTP router & middleware |
| **Database** | SQLite | 3.x | Penyimpanan data |
| **HTTP Library** | net/http | Standard library | HTTP server/client |
| **CORS** | gin-contrib/cors | 1.7.7 | Cross-origin handling |

### 🎨 Fitur Utama

- 🚀 **Single Binary Deployment**: Hasil compile adalah satu executable file tanpa dependencies
- 📊 **Structured Logging**: Log output yang jelas dan mudah dipantau
- 🔒 **Secure CORS**: Konfigurasi CORS yang aman untuk production
- 💾 **SQLite Database**: Database tanpa server, mudah di-setup
- 📚 **Learning-Friendly**: Kode dengan komentar detail untuk belajar Go
- 🔄 **Type-Safe**: Compile-time type checking (beda dengan JavaScript!)

---

## 👨‍💻 Panduan untuk Pengembang JS

Selamat datang di Go! Sebagai developer JavaScript yang berpengalaman, kamu mungkin terbiasa dengan struktur Node.js. Berikut perbandingan penting:

### 📦 Package.json vs Go Module

| JavaScript (Node.js) | Go | Keterangan |
|----------------------|-----|------------|
| `package.json` | `go.mod` | Daftar dependencies |
| `package-lock.json` | `go.sum` | Checksum & version lock |
| `npm install` | `go mod download` | Download dependencies |
| `npm run dev` | `go run .` | Run project |
| `npm run build` | `go build` | Compile executable |

**Perbedaan Utama:**
- **JavaScript**: Interpreted language, perlu Node.js runtime
- **Go**: Compiled language, hasilnya single binary (tidak butuh runtime!)

### 📁 Struktur Folder: Node.js vs Go

```
Node.js Project Structure          │  Go Project Structure
─────────────────────────────────────┼─────────────────────────────────────
my-nodejs-app/                      │  cashflow-backend/
├── src/                            │  ├── main.go (entry point)
│   ├── index.js (entry point)      │  ├── go.mod (dependencies)
│   ├── controllers/                │  ├── go.sum (checksums)
│   │   └── userController.js      │  └── cashflow.db (database)
│   ├── models/                     │
│   │   └── user.js                 │
│   └── routes/                     │
│       └── userRoutes.js           │
├── package.json                    │
└── node_modules/ (dependencies)    │  // Dependencies di:
                                     │  // $GOPATH/pkg/mod/ (global cache)
```

### 🎯 Entry Point: App.listen() vs main()

**JavaScript (Express.js):**
```javascript
// src/index.js atau app.js
const express = require('express');
const app = express();

app.listen(8081, () => {
  console.log('Server started on port 8081');
});
```

**Go:**
```go
// main.go
package main

import (
  "log"
  "net/http"
)

func main() {  // ← Entry point, HARUS ada
  log.Println("Server starting...")
  // Setup code here...
}
```

### 🔧 Dependency Management: npm vs go mod

**JavaScript:**
```bash
npm install express
npm install --save-dev jest
```

**Go:**
```bash
go get github.com/gin-gonic/gin
go get -u github.com/gin-gonic/gin  # update
go mod tidy  # clean unused deps
```

### 🚀 Running Project: npm start vs go run

**JavaScript:**
```bash
npm start           # start production
npm run dev         # start development
node index.js       # run directly
```

**Go:**
```bash
go run .            # run without compile (development)
go build            # compile executable
./cashflow-backend  # run compiled binary
```

### 📦 Import: require/import vs import

**JavaScript (ES6):**
```javascript
import express from 'express';
import { getUser } from './controllers/userController';
const sqlite3 = require('sqlite3');
```

**Go:**
```go
import (
  "database/sql"
  "net/http"
  "github.com/gin-gonic/gin"
  _ "modernc.org/sqlite"  // blank import untuk side effects
)
```

### 🔍 Error Handling: try/catch vs if err != nil

**JavaScript:**
```javascript
try {
  const result = await riskyOperation();
  return result;
} catch (error) {
  console.error(error);
  throw error;
}
```

**Go:**
```go
result, err := riskyOperation()
if err != nil {
  log.Fatal(err)  // atau return err
}
return result
```

**Catatan:** Go TIDAK punya try/catch! Error handling sangat eksplisit.

### 🏗️ Data Structures: Objects vs Structs

**JavaScript:**
```javascript
const transaction = {
  id: 1,
  tanggal: "2026-04-06",
  jenis: "Pemasukan",
  nominal: 100000
};
```

**Go:**
```go
type Transaction struct {
  ID        int    `json:"id"`
  Tanggal   string `json:"tanggal"`
  Jenis     string `json:"jenis"`
  Nominal   int    `json:"nominal"`
}
```

**Perbedaan:**
- **JavaScript**: Objects dinamis, bisa tambah properti sembarang
- **Go**: Structs punya tipe tetap (type-safe)

### 🌐 HTTP Server: app.get() vs r.GET()

**JavaScript (Express.js):**
```javascript
app.get('/ping', (req, res) => {
  res.json({ message: 'pong' });
});
```

**Go (Gin):**
```go
r.GET("/ping", func(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "message": "pong",
  })
})
```

---

## 🛠️ Prasyarat

Sebelum memulai, pastikan Anda sudah memenuhi prasyarat berikut:

### 1. Go Installation

**Cek versi Go yang sudah terinstall:**
```bash
go version
```

**Output yang diharapkan:**
```
go version go1.25.6 linux/amd64
```

### 2. Jika Belum Install Go

**Windows:**
```bash
# Download installer dari: https://go.dev/dl/
# Run installer dan follow wizard
# Verify installation:
go version
```

**Linux/WSL:**
```bash
# Download Go
wget https://go.dev/dl/go1.25.6.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.25.6.linux-amd64.tar.gz

# Add to PATH (tambahkan ke ~/.bashrc atau ~/.zshrc)
export PATH=$PATH:/usr/local/go/bin

# Verify
go version
```

**macOS:**
```bash
# Install dengan Homebrew
brew install go

# Verify
go version
```

### 3. Environment Setup (Optional tapi Recommended)

```bash
# Set GOPATH (Go 1.11+ sudah pakai Go modules, tapi tetap useful)
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# Verify Go environment
go env
```

### 4. Development Tools (Optional)

**VS Code Extensions:**
- Go (golang.go)
- SQLite (alexcvzz.vscode-sqlite)

---

## 🚀 Cara Menjalankan

Ikuti langkah-langkah ini untuk menjalankan proyek:

### 📝 Langkah 1: Navigate ke Project Directory

```bash
cd /mnt/d/._show-case/go-cahsflow-rest-api
```

### 📦 Langkah 2: Install Dependencies

Mirip dengan `npm install`, tapi di Go kita pakai `go get`:

```bash
# Install semua dependencies
go get github.com/gin-gonic/gin
go get modernc.org/sqlite
go get github.com/gin-contrib/cors

# Clean up unused dependencies
go mod tidy
```

**Apa yang terjadi:**
- `go get`: Download package ke cache lokal (`$GOPATH/pkg/mod/`)
- `go.mod`: File yang track semua dependencies (mirip `package.json`)
- `go.sum`: File dengan cryptographic hashes (mirip `package-lock.json`)
- `go mod tidy`: Remove dependencies yang tidak dipakai

### 🔨 Langkah 3: Build Application

```bash
# Build executable untuk sistem saat ini
go build -o cashflow-backend.exe .
```

**Apa yang terjadi:**
- `go build`: Compile Go code ke binary executable
- `-o cashflow-backend.exe`: Nama output file
- `.`: Source directory (current directory)
- Hasilnya adalah **single binary** - tidak butuh Node.js atau dependencies lain!

**Output:**
```
cashflow-backend.exe (Windows)
cashflow-backend (Linux/macOS)
```

### ▶️ Langkah 4: Run Server

Ada 3 cara untuk menjalankan server:

**Cara 1: Run dari Binary (Production)**
```bash
# Windows
./cashflow-backend.exe

# Linux/macOS
./cashflow-backend
```

**Cara 2: Run Tanpa Compile (Development)**
```bash
go run .
```

**Cara 3: Run di Background (Linux/macOS)**
```bash
# Run di background
./cashflow-backend &

# Atau dengan nohup (tetap jalan setelah logout)
nohup ./cashflow-backend &

# Check logs
tail -f server.log
```

### ✅ Langkah 5: Verify Server Running

```bash
# Test ping endpoint
curl http://localhost:8081/ping

# Expected response:
# {"message":"pong"}
```

**Output yang diharapkan:**
```
2026/04/06 11:00:00 Database initialized successfully
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
[GIN-debug] Listening and serving HTTP on :8081
```

### 🛑 Langkah 6: Stop Server

**Stop dengan Ctrl+C** (jika jalan di foreground)

```bash
# Atau kill process
pkill cashflow-backend

# Atau find dan kill PID
ps aux | grep cashflow-backend
kill <PID>
```

---

## 📡 Dokumentasi API/Fungsi

Proyek ini saat ini memiliki endpoint dasar. Berikut dokumentasinya:

### 🔍 GET /ping

**Deskripsi:** Health check endpoint untuk memverifikasi server berjalan.

**Request:**
```bash
curl http://localhost:8081/ping
```

**Response:**
```json
{
  "message": "pong"
}
```

**Status Code:**
- `200 OK`: Server berjalan normal

### 📊 Database Schema

Tabel `transactions` dibuat otomatis dengan schema berikut:

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

**Field Mapping (Struct Go):**

| Field Go | Type | JSON Key | Deskripsi |
|----------|------|----------|-----------|
| `ID` | `int` | `"id"` | Primary key (auto-increment) |
| `Tanggal` | `string` | `"tanggal"` | Tanggal transaksi (format: YYYY-MM-DD) |
| `Jenis` | `string` | `"jenis"` | Jenis transaksi (Pemasukan/Pengeluaran) |
| `Kategori` | `string` | `"kategori"` | Kategori transaksi |
| `Nominal` | `int` | `"nominal"` | Jumlah nominal (integer) |
| `Keterangan` | `string` | `"keterangan"` | Deskripsi tambahan |

---

## ⚙️ Penjelasan Konsep Go

Proyek ini menggunakan beberapa konsep penting di Go. Berikut penjelasannya:

### 1. 📦 Package System

**Konsep:**
- `package main`: Package khusus untuk membuat executable
- Harus ada `func main()` di package main untuk titik masuk program

**Contoh:**
```go
package main  // ← Package declaration

func main() {  // ← Entry point
  // Code here
}
```

**JavaScript Analogi:**
- Mirip dengan `index.js` atau `app.js` sebagai entry point

### 2. 🏗️ Structs

**Konsep:**
- Go menggunakan structs, bukan classes
- Structs mengelompokkan field data yang terkait
- Tipe data sangat strict (type-safe)

**Contoh:**
```go
type Transaction struct {
  ID        int    `json:"id"`
  Tanggal   string `json:"tanggal"`
  Jenis     string `json:"jenis"`
}
```

**JavaScript Analogi:**
```javascript
// JavaScript Object (loose typing)
const transaction = { id: 1, tanggal: "2026-04-06" };

// JavaScript Class (OOP)
class Transaction {
  constructor(id, tanggal) {
    this.id = id;
    this.tanggal = tanggal;
  }
}
```

### 3. 🔄 Multiple Returns

**Konsep:**
- Functions di Go bisa mengembalikan multiple values
- Pattern umum: `(value, error)` - return value dan error

**Contoh:**
```go
db, err := sql.Open("sqlite", "./cashflow.db")
if err != nil {
  log.Fatal(err)
}
```

**JavaScript Analogi:**
```javascript
// JavaScript tidak punya multiple returns
// Biasanya pakai object atau array
const result = await riskyOperation();
if (result.error) {
  console.error(result.error);
}
```

### 4. ❌ Error Handling

**Konsep:**
- Go TIDAK punya try/catch
- Error handling sangat eksplisit dengan `if err != nil`
- Prevents hidden errors

**Contoh:**
```go
// Go - explicit error checking
result, err := someFunction()
if err != nil {
  log.Fatal(err)
}
// Process result...
```

**JavaScript Analogi:**
```javascript
// JavaScript - try/catch
try {
  const result = await someFunction();
  // Process result...
} catch (error) {
  console.error(error);
}
```

### 5. ⏳ Defer Statement

**Konsep:**
- `defer` menjadwalkan function untuk dijalankan saat function exits
- LIFO order (Last In, First Out)
- Dijalankan bahkan jika panic() terjadi

**Contoh:**
```go
func main() {
  db, err := sql.Open("sqlite", "./cashflow.db")
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()  // ← Ditutup saat main() exits

  // Code here...
}
```

**JavaScript Analogi:**
```javascript
// JavaScript - finally block
try {
  const db = await connectDB();
  // Code here...
} catch (error) {
  console.error(error);
} finally {
  await db.close();  // ← Selalu jalan
}
```

### 6. 🌐 HTTP Server (Gin)

**Konsep:**
- Gin adalah HTTP web framework performa tinggi
- Router menangani URL routing dan request/response
- Middleware untuk request processing pipeline

**Contoh:**
```go
r := gin.Default()  // Router dengan default middleware

r.GET("/ping", func(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "message": "pong",
  })
})
```

**JavaScript Analogi:**
```javascript
// Express.js
const app = express();

app.get('/ping', (req, res) => {
  res.json({ message: 'pong' });
});
```

### 7. 🗄️ Database Operations

**Konsep:**
- `database/sql`: Standard SQL interface
- `sql.Open()`: Buka koneksi database
- `db.Exec()`: Eksekusi SQL tanpa return rows
- `db.Query()`: Eksekusi SQL yang return rows

**Contoh:**
```go
db, err := sql.Open("sqlite", "./cashflow.db")
_, err = db.Exec("CREATE TABLE IF NOT EXISTS ...")
```

**JavaScript Analogi:**
```javascript
// Node.js sqlite3
const db = new sqlite3.Database('./cashflow.db');
db.run("CREATE TABLE IF NOT EXISTS ...");
```

---

## 📁 Struktur Folder

```
go-cahsflow-rest-api/
├── main.go                  # Entry point aplikasi (mirip index.js)
├── go.mod                   # Module dependencies (mirip package.json)
├── go.sum                   # Dependency checksums (mirip package-lock.json)
├── cashflow.db              # SQLite database file (auto-created)
├── cashflow-backend.exe     # Compiled binary (Windows)
├── server.log               # Server logs (opsional)
└── README.md                # Dokumentasi ini
```

### 📂 Penjelasan Setiap File

| File | Deskripsi | Mirip di JavaScript |
|------|----------|---------------------|
| `main.go` | Entry point aplikasi dan semua logic | `index.js` atau `server.js` |
| `go.mod` | Module dependencies & version | `package.json` |
| `go.sum` | Dependency checksums & security | `package-lock.json` |
| `cashflow.db` | SQLite database file | `database.sqlite` |
| `cashflow-backend.exe` | Compiled executable | Tidak ada (JS interpreted) |
| `README.md` | Dokumentasi proyek | `README.md` |

### 🔄 Perbandingan Workflow

**JavaScript (Node.js):**
```
npm install → node_modules/ dependencies
npm run dev → Start development server
npm run build → Bundle (optional)
```

**Go:**
```
go get → $GOPATH/pkg/mod/ dependencies (global cache)
go run . → Run without compile
go build → Compile to single binary
./binary → Run compiled binary
```

---

## 🐛 Troubleshooting

### Error 1: Port Already in Use

**Problem:**
```
listen tcp :8081: bind: address already in use
```

**Solution:**
```bash
# Find process using port 8081
netstat -ano | findstr :8081  # Windows
lsof -ti:8081                  # Linux/macOS

# Kill the process
taskkill /PID <PID> /F         # Windows
kill -9 <PID>                  # Linux/macOS
```

### Error 2: Module Not Found

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

### Error 3: Connection Refused

**Problem:**
```
curl: (7) Failed to connect to localhost port 8081: Connection refused
```

**Solution:**
```bash
# Verify server is running
ps aux | grep cashflow-backend

# Restart server
./cashflow-backend.exe
```

### Error 4: CORS Issues

**Problem:** Frontend tidak bisa akses API

**Solution:**
```go
// Check CORS config di main.go
// Pastikan origin frontend kamu ada di AllowOrigins
r.Use(cors.New(cors.Config{
    AllowOrigins: []string{"http://localhost:5173"},
}))
```

---

## 📚 Lanjutkan Belajar

### Next Steps untuk CRUD Endpoints

Setelah `GET /ping` berjalan, implementasikan endpoint CRUD:

1. **GET /transactions** - Ambil semua transactions
2. **POST /transactions** - Buat transaction baru
3. **PUT /transactions/:id** - Update transaction
4. **DELETE /transactions/:id** - Hapus transaction

### Resources untuk Belajar Lanjut

**📖 Dokumentasi Resmi:**
- [Go Documentation](https://go.dev/doc/)
- [Gin Framework](https://gin-gonic.com/docs/)
- [SQL Database](https://go.dev/doc/database/sql-overview)

**🎓 Tutorial:**
- [Go by Example](https://gobyexample.com/)
- [A Tour of Go](https://tour.go.dev/welcome/1)

### Useful Go Commands

```bash
# Module management
go mod init <module-name>    # Initialize module
go mod tidy                  # Clean dependencies
go mod download              # Download dependencies

# Building
go build                     # Build executable
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

### JavaScript vs Go Quick Reference

| Konsep | JavaScript | Go |
|--------|-----------|-----|
| Variable | `let x = 10` | `x := 10` |
| Function | `const add = (a,b) => a+b` | `func add(a,b int) int { return a+b }` |
| Object | `{ name: "John" }` | `struct { Name string }` |
| Array | `[1, 2, 3]` | `[]int{1, 2, 3}` |
| Error | `try/catch` | `if err != nil` |
| Async | `async/await` | `goroutines + channels` |
| Import | `import express` | `import "github.com/gin-gonic/gin"` |
| Package | NPM packages | Go modules |

---

## 📊 Project Status

- ✅ Server berjalan di `localhost:8081`
- ✅ Database SQLite dibuat otomatis
- ✅ CORS middleware enabled
- ✅ GET /ping endpoint berfungsi
- ⏳ GET /transactions endpoint (belum diimplementasi)
- ⏳ POST /transactions endpoint (belum diimplementasi)
- ⏳ PUT /transactions endpoint (belum diimplementasi)
- ⏳ DELETE /transactions endpoint (belum diimplementasi)

---

**🎉 Selamat Belajar Go!** 

Proyek ini dirancang khusus untuk membantu developer JavaScript bertransisi ke Go. Kode sudah dilengkapi dengan komentar detail yang menjelaskan setiap konsep Go dengan analogi JavaScript.

Happy coding! 🚀
