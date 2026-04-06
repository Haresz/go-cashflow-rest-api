// ============================================================================
// KONSEP: Package Declaration & Entry Point
// BELAJAR: Go mengorganisasi kode ke dalam packages. 'main' adalah package khusus untuk membuat executable.
// POIN PENTING:
//   - package main: Diperlukan untuk program executable (harus punya func main())
//   - func main(): Titik masuk - jalan pertama saat program mulai
// SAMA SEPERTI: Tidak ada di JavaScript, tapi mirip main() function di Node.js CLI atau入口点 di Node.js app
// CONTOH JAVASCRIPT:
//   // Tidak ada konsep package di JavaScript, tapi ini mirip dengan entry point:
//   // app.listen(3000, () => { console.log('Server started'); })
// ============================================================================

package main

// ============================================================================
// KONSEP: Imports - System dan Third-Party Packages
// BELAJAR: Import packages untuk menggunakan fungsionalitasnya. Go punya standard library + third-party.
// POIN PENTING:
//   - "database/sql": Standard SQL interface (mirip database driver di Node.js seperti pg, mysql2)
//   - "log": Logging utilities (mirip console.log di JavaScript)
//   - "net/http": HTTP server/client (mirip express.js atau http module di Node.js)
//   - _ "modernc.org/sqlite": Blank import - side effects saja (register driver)
// SAMA SEPERTI: import statements di JavaScript ES6 (import express from 'express')
// CONTOH JAVASCRIPT:
//   // JavaScript/Node.js
//   import express from 'express';
//   import sqlite3 from 'sqlite3';
//   const logger = require('./logger');
// ============================================================================

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

// ============================================================================
// KONSEP: Structs - Custom Data Types
// BELAJAR: Go menggunakan structs bukan classes. Structs mengelompokkan field data yang terkait.
// POIN PENTING:
//   - type struct: Mendefinisikan struktur data kustom (mirip class dengan hanya fields di JavaScript)
//   - PascalCase: Exported fields (public - bisa diakses dari package lain)
//   - json tags: Mengontrol serialisasi JSON (mirip JSON.parse/stringify di JavaScript)
//   - Konvensi Go: Public = PascalCase, Private = camelCase
// SAMA SEPERTI: Objects di JavaScript, tapi dengan tipe yang strict
// CONTOH JAVASCRIPT:
//   // JavaScript Object (loosely typed)
//   const transaction = {
//     id: 1,
//     tanggal: "2025-04-06",
//     jenis: "Pemasukan"
//   };
//
//   // JavaScript Class (OOP)
//   class Transaction {
//     constructor(id, tanggal, jenis) {
//       this.id = id;
//       this.tanggal = tanggal;
//       this.jenis = jenis;
//     }
//   }
// ============================================================================

type Transaction struct {
	ID        int    `json:"id"`
	Tanggal   string `json:"tanggal"`
	Jenis     string `json:"jenis"`
	Kategori  string `json:"kategori"`
	Nominal   int    `json:"nominal"`
	Keterangan string `json:"keterangan"`
}

// ============================================================================
// KONSEP: Main Function - Program Entry Point
// BELAJAR: Setiap program executable Go harus punya main() function di package main.
// POIN PENTING:
//   - Tidak ada parameters: Menggunakan command-line args via os.Args (mirip process.argv di Node.js)
//   - Tidak ada return value: Gunakan os.Exit() untuk exit codes (mirip process.exit di Node.js)
// SAMA SEPERTI: main() function di Node.js CLI, atau入口点 di aplikasi Node.js
// CONTOH JAVASCRIPT:
//   // JavaScript/Node.js entry point
//   const express = require('express');
//   const app = express();
//   
//   app.listen(8080, () => {
//     console.log('Server started');
//   });
// ============================================================================

func main() {
	// ============================================================================
	// KONSEP: Database Connection - sql.Open()
	// BELAJAR: Buka koneksi database. Connection pooling ditangani secara otomatis.
	// POIN PENTING:
	//   - sql.Open("driver", "dataSource"): Mengembalikan *sql.DB dan error
	//   - *sql.DB: Pointer ke database object (mirip connection pool di Node.js)
	//   - Multiple returns: (value, error) - pattern error handling Go
	//   - "./cashflow.db": SQLite file path (buat jika belum ada)
	// SAMA SEPERTI: new sqlite3.Database() di Node.js, atau new Pool() di pg
	// CONTOH JAVASCRIPT:
	//   // JavaScript dengan sqlite3
	//   const sqlite3 = require('sqlite3');
	//   const db = new sqlite3.Database('./cashflow.db', (err) => {
	//     if (err) console.error(err);
	//   });
	//   
	//   // JavaScript dengan pg (PostgreSQL)
	//   const { Pool } = require('pg');
	//   const pool = new Pool({
	//     connectionString: 'postgres://localhost/mydb'
	//   });
	// ============================================================================
	db, err := sql.Open("sqlite", "./cashflow.db")
	
	// ============================================================================
	// KONSEP: Error Handling Pattern
	// BELAJAR: Go menggunakan explicit error checking bukan try/catch.
	// POIN PENTING:
	//   - if err != nil: Cek jika error ada (mirip catching exceptions)
	//   - log.Fatal(err): Print error dan exit program (mirip console.error + process.exit)
	//   - Tidak ada try/catch: Explicit checking mencegah hidden errors
	// SAMA SEPERTI: try/catch di JavaScript, tapi lebih explicit dan predictable
	// CONTOH JAVASCRIPT:
	//   // JavaScript dengan try/catch
	//   try {
	//     const result = riskyOperation();
	//   } catch (error) {
	//     console.error(error);
	//     process.exit(1);
	//   }
	//   
	//   // JavaScript dengan promise
	//   riskyOperation()
	//     .then(result => { /* handle success */ })
	//     .catch(error => { console.error(error); });
	// ============================================================================
	if err != nil {
		log.Fatal(err)
	}

	// ============================================================================
	// KONSEP: Defer Statement - Cleanup Execution
	// BELAJAR: Defer menjadwalkan function untuk dijalankan saat function saat ini returns.
	// POIN PENTING:
	//   - defer db.Close(): Tutup koneksi saat main() exits
	//   - LIFO order: Multiple defers jalan urutan terbalik (Last In, First Out)
	//   - Guaranteed: Selalu jalan, bahkan jika panic() terjadi (mirip finally di try/catch)
	// SAMA SEPERTI: finally block di try/catch di JavaScript
	// CONTOH JAVASCRIPT:
	//   // JavaScript dengan finally
	//   try {
	//     const data = await fetchData();
	//   } catch (error) {
	//     console.error(error);
	//   } finally {
	//     db.close(); // selalu jalan
	//   }
	// ============================================================================

	defer db.Close()

	// ============================================================================
	// KONSEP: SQL Table Creation - db.Exec()
	// BELAJAR: Eksekusi SQL statements yang tidak return rows (INSERT, UPDATE, DELETE, CREATE).
	// POIN PENTING:
	//   - Backticks: Raw string literals (tidak butuh escape characters)
	//   - db.Exec(query, args...): Eksekusi SQL dengan optional parameters
	//   - _, err = db.Exec(): Underscore ignore return value (kita tidak butuh Result)
	//   - CREATE TABLE IF NOT EXISTS: Operasi idempotent (aman untuk jalan berkali-kali)
	// SAMA SEPERTI: db.run() di Node.js sqlite3, atau db.query() di pg
	// CONTOH JAVASCRIPT:
	//   // JavaScript dengan sqlite3
	//   db.run(`
	//     CREATE TABLE IF NOT EXISTS transactions (
	//       id INTEGER PRIMARY KEY AUTOINCREMENT,
	//       tanggal TEXT NOT NULL,
	//       jenis TEXT NOT NULL,
	//       kategori TEXT NOT NULL,
	//       nominal INTEGER NOT NULL,
	//       keterangan TEXT
	//     )
	//   `, (err) => {
	//     if (err) console.error(err);
	//   });
	// ============================================================================
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		tanggal TEXT NOT NULL,
		jenis TEXT NOT NULL,
		kategori TEXT NOT NULL,
		nominal INTEGER NOT NULL,
		keterangan TEXT
	);`
	
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database initialized successfully")

	// ============================================================================
	// KONSEP: Gin Router - HTTP Framework
	// BELAJAR: Gin adalah web framework untuk membangun HTTP servers.
	// POIN PENTING:
	//   - gin.Default(): Creates router dengan Logger dan Recovery middleware
	//   - Router: Menangani URL routing dan request/response handling
	//   - Alternative: gin.New() untuk router tanpa middleware
	// SAMA SEPERTI: express() di Express.js, atau Koa router
	// CONTOH JAVASCRIPT:
	//   // JavaScript dengan Express.js
	//   const express = require('express');
	//   const app = express();
	//   
	//   // Middleware default
	//   app.use(express.json());
	//   app.use(express.urlencoded({ extended: true }));
	// ============================================================================
	r := gin.Default()

	// ============================================================================
	// KONSEP: Middleware - Request Processing Pipeline
	// BELAJAR: Middleware functions jalan sebelum/sesudah handlers. Digunakan untuk CORS, auth, logging.
	// POIN PENTING:
	//   - r.Use(middleware): Register middleware untuk jalan di semua requests
	//   - CORS: Cross-Origin Resource Sharing (browser security)
	//   - Config: Define domains yang boleh akses API kamu
	//   - AllowOrigins: ["http://localhost:5173"] - Hanya boleh React app kamu
	// SAMA SEPERTI: app.use() di Express.js, middleware di Koa
	// CONTOH JAVASCRIPT:
	//   // JavaScript dengan Express.js
	//   const cors = require('cors');
	//   app.use(cors({
	//     origin: 'http://localhost:5173',
	//     methods: ['GET', 'POST', 'PUT', 'DELETE'],
	//     allowedHeaders: ['Content-Type']
	//   }));
	// ============================================================================
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// ============================================================================
	// KONSEP: Route Definition - HTTP Endpoints
	// BELAJAR: Define URL patterns dan handler functions mereka.
	// POIN PENTING:
	//   - r.GET(path, handler): Handle GET requests di path yang ditentukan
	//   - func(c *gin.Context): Handler function menerima request/response context
	//   - c.JSON(status, data): Return JSON response secara otomatis
	//   - http.StatusOK: Constant untuk HTTP 200 OK
	// SAMA SEPERTI: app.get() di Express.js, router.get() di Koa
	// CONTOH JAVASCRIPT:
	//   // JavaScript dengan Express.js
	//   app.get('/ping', (req, res) => {
	//     res.json({ message: 'pong' });
	//   });
	// ============================================================================
	r.GET("/ping", func(c *gin.Context) {
		// gin.H adalah shortcut untuk map[string]interface{} (JSON object)
		// Sama seperti object literal di JavaScript
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// ============================================================================
	// KONSEP: Server Startup - HTTP Server
	// BELAJAR: Start HTTP server untuk listen incoming requests.
	// POIN PENTING:
	//   - r.Run(":8081"): Listen di port 8081 (bind ke semua interfaces)
	//   - Blocks execution: Server jalan sampai stopped atau error terjadi
	//   - Error handling: Check jika server gagal start (port in use, dll)
	// SAMA SEPERTI: app.listen() di Express.js, server.listen() di Node.js
	// CONTOH JAVASCRIPT:
	//   // JavaScript dengan Express.js
	//   app.listen(8081, () => {
	//     console.log('Server running on http://localhost:8081');
	//   });
	// ============================================================================
	log.Println("Server starting on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// ============================================================================
// RINGKASAN KONSEP GO YANG DIPELAJARI:
// 
// 1. Package System: package main, import statements
// 2. Structs: Custom data types dengan JSON tags
// 3. Multiple Returns: (value, error) pattern
// 4. Error Handling: if err != nil pattern (tidak ada try/catch)
// 5. Defer: Cleanup statements (LIFO execution)
// 6. Database Operations: sql.Open, db.Exec
// 7. Web Framework: Gin router, handlers, middleware
// 8. HTTP Handling: Context, JSON responses
// 9. Server Configuration: Port binding, error handling
// ============================================================================
