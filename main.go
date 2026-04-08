package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

type Transaction struct {
	ID        int    `json:"id"`
	Tanggal   string `json:"tanggal"`
	Jenis     string `json:"jenis"`
	Kategori  string `json:"kategori"`
	Nominal   uint    `json:"nominal"`
	Keterangan string `json:"keterangan"`
}

func main() {
	db, err := sql.Open("sqlite", "./cashflow.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

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

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})


	r.GET("/transactions", getTransactionsHandler(db))
	r.GET("/transactions/:id", getTransactionByIDHandler(db))
	r.POST("/transactions", createTransactionHandler(db))
	r.PUT("/transactions/:id", updateTransactionHandler(db))
	r.DELETE("/transactions/:id", deleteTransactionHandler(db))

	log.Println("Server starting on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func getAllTransactions(db *sql.DB, filters map[string]string, sortColumn string, sortOrder string, page int, limit int) ([]Transaction, int64, error) {
	whereClause := " WHERE 1=1"
	query := "SELECT id, tanggal, jenis, kategori, nominal, keterangan FROM transactions"
	args := []interface{}{}
	argCount := 0

	if jenis, ok := filters["jenis"]; ok {
		argCount++
		whereClause += fmt.Sprintf(" AND jenis = $%d", argCount)
		args = append(args, jenis)
	}

	if kategori, ok := filters["kategori"]; ok {
		argCount++
		whereClause += fmt.Sprintf(" AND kategori = $%d", argCount)
		args = append(args, kategori)
	}

	if tanggal, ok := filters["tanggal"]; ok {
		argCount++
		whereClause += fmt.Sprintf(" AND tanggal = $%d", argCount)
		args = append(args, tanggal)
	}

	if startDate, ok := filters["startDate"]; ok {
		argCount++
		whereClause += fmt.Sprintf(" AND tanggal >= $%d", argCount)
		args = append(args, startDate)
	}

	if endDate, ok := filters["endDate"]; ok {
		argCount++
		whereClause += fmt.Sprintf(" AND tanggal <= $%d", argCount)
		args = append(args, endDate)
	}

	query += whereClause

	usePagination := (page > 0 && limit > 0)
	if usePagination {
		validColumns := map[string]bool{
			"id": true, "tanggal": true, "jenis": true,
			"kategori": true, "nominal": true, "keterangan": true,
		}
		if !validColumns[sortColumn] {
			sortColumn = "tanggal"
		}

		if sortOrder != "ASC" && sortOrder != "DESC" {
			sortOrder = "DESC"
		}

		query += fmt.Sprintf(" ORDER BY %s %s", sortColumn, sortOrder)

		argCount++
		query += fmt.Sprintf(" LIMIT $%d", argCount)
		args = append(args, limit)

		argCount++
		offset := (page - 1) * limit
		query += fmt.Sprintf(" OFFSET $%d", argCount)
		args = append(args, offset)
	} else {
		query += " ORDER BY id DESC"
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var t Transaction
		err := rows.Scan(&t.ID, &t.Tanggal, &t.Jenis, &t.Kategori, &t.Nominal, &t.Keterangan)
		if err != nil {
			return nil, 0, err
		}
		transactions = append(transactions, t)
	}

	countQuery := "SELECT COUNT(*) FROM transactions" + whereClause
	countArgs := make([]interface{}, len(args)-2)
	if usePagination {
		copy(countArgs, args[:len(args)-2])
	} else {
		copy(countArgs, args)
	}

	var total int64
	err = db.QueryRow(countQuery, countArgs...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return transactions, total, nil
}

func getTransactionByID(db *sql.DB, id int) (*Transaction, error) {
	var t Transaction
	err := db.QueryRow(
		"SELECT id, tanggal, jenis, kategori, nominal, keterangan FROM transactions WHERE id = $1",
		id,
	).Scan(&t.ID, &t.Tanggal, &t.Jenis, &t.Kategori, &t.Nominal, &t.Keterangan)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func createTransaction(db *sql.DB, t Transaction) (int64, error) {
	result, err := db.Exec(
		"INSERT INTO transactions (tanggal, jenis, kategori, nominal, keterangan) VALUES ($1, $2, $3, $4, $5)",
		t.Tanggal, t.Jenis, t.Kategori, t.Nominal, t.Keterangan,
	)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func updateTransaction(db *sql.DB, t Transaction) error {
	result, err := db.Exec(
		"UPDATE transactions SET tanggal=$1, jenis=$2, kategori=$3, nominal=$4, keterangan=$5 WHERE id=$6",
		t.Tanggal, t.Jenis, t.Kategori, t.Nominal, t.Keterangan, t.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("transaction not found")
	}

	return nil
}

func deleteTransaction(db *sql.DB, id int) error {
	result, err := db.Exec("DELETE FROM transactions WHERE id=$1", id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("transaction not found")
	}

	return nil
}

func validateTransaction(t Transaction) error {
	if t.Tanggal == "" {
		return errors.New("tanggal is required")
	}
	if t.Jenis != "Pemasukan" && t.Jenis != "Pengeluaran" {
		return errors.New("jenis must be 'Pemasukan' or 'Pengeluaran'")
	}
	if t.Kategori == "" {
		return errors.New("kategori is required")
	}
	if t.Nominal <= 0 {
		return  errors.New("nominal must be greater than 0")
	}
	return nil
}

func sendSuccessResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{
		"success": true,
		"data":    data,
	})
}

func sendErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Printf("[ERROR] %s", message)
	c.JSON(statusCode, gin.H{
		"success": false,
		"error":   message,
	})
}

func getPaginationMetadata(page int, limit int, total int64) map[string]interface{} {
	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	hasNext := page < totalPages
	hasPrev := page > 1

	return map[string]interface{}{
		"page":       page,
		"limit":      limit,
		"total":      total,
		"totalPages": totalPages,
		"hasNext":    hasNext,
		"hasPrev":    hasPrev,
	}
}

func getTransactionsHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		filters := make(map[string]string)

		if jenis := c.Query("jenis"); jenis != "" {
			filters["jenis"] = jenis
		}
		if kategori := c.Query("kategori"); kategori != "" {
			filters["kategori"] = kategori
		}
		if tanggal := c.Query("tanggal"); tanggal != "" {
			filters["tanggal"] = tanggal
		}
		if startDate := c.Query("startDate"); startDate != "" {
			filters["startDate"] = startDate
		}
		if endDate := c.Query("endDate"); endDate != "" {
			filters["endDate"] = endDate
		}

		page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "0"))

		usePagination := (page > 0 && limit > 0)

		var sortColumn, sortOrder string
		if usePagination {
			sortColumn = c.DefaultQuery("sortColumn", "tanggal")
			sortOrder = c.DefaultQuery("sortOrder", "DESC")

			if page < 1 {
				page = 1
			}
			if limit < 1 || limit > 100 {
				limit = 10
			}
		} else {
			sortColumn = "id"
			sortOrder = "DESC"
		}

		transactions, total, err := getAllTransactions(db, filters, sortColumn, sortOrder, page, limit)
		if err != nil {
			sendErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Failed to get transactions: %v", err))
			return
		}

		response := gin.H{"transactions": transactions}

		if usePagination {
			response["pagination"] = getPaginationMetadata(page, limit, total)
		}

		log.Printf("[INFO] GET /transactions - Retrieved %d transactions", len(transactions))
		sendSuccessResponse(c, http.StatusOK, response)
	}
}

func getTransactionByIDHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			sendErrorResponse(c, http.StatusBadRequest, "Invalid transaction ID")
			return
		}

		transaction, err := getTransactionByID(db, id)
		if err != nil {
			sendErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Failed to get transaction: %v", err))
			return
		}

		if transaction == nil {
			sendErrorResponse(c, http.StatusNotFound, "Transaction not found")
			return
		}

		log.Printf("[INFO] GET /transactions/%d - Retrieved transaction", id)
		sendSuccessResponse(c, http.StatusOK, gin.H{
			"transaction": transaction,
		})
	}
}

func createTransactionHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var t Transaction
		if err := c.ShouldBindJSON(&t); err != nil {
			sendErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid request body: %v", err))
			return
		}

		if err := validateTransaction(t); err != nil {
			sendErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Validation error: %v", err))
			return
		}

		id, err := createTransaction(db, t)
		if err != nil {
			sendErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Failed to create transaction: %v", err))
			return
		}

		log.Printf("[INFO] POST /transactions - Created transaction ID: %d", id)
		sendSuccessResponse(c, http.StatusCreated, gin.H{
			"id": id,
		})
	}
}

func updateTransactionHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			sendErrorResponse(c, http.StatusBadRequest, "Invalid transaction ID")
			return
		}

		var t Transaction
		if err := c.ShouldBindJSON(&t); err != nil {
			sendErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid request body: %v", err))
			return
		}

		t.ID = id

		if err := validateTransaction(t); err != nil {
			sendErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Validation error: %v", err))
			return
		}

		if err := updateTransaction(db, t); err != nil {
			if err.Error() == "transaction not found" {
				sendErrorResponse(c, http.StatusNotFound, "Transaction not found")
				return
			}
			sendErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Failed to update transaction: %v", err))
			return
		}

		log.Printf("[INFO] PUT /transactions/%d - Updated transaction", id)
		sendSuccessResponse(c, http.StatusOK, gin.H{
			"id": id,
		})
	}
}

func deleteTransactionHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			sendErrorResponse(c, http.StatusBadRequest, "Invalid transaction ID")
			return
		}

		if err := deleteTransaction(db, id); err != nil {
			if err.Error() == "transaction not found" {
				sendErrorResponse(c, http.StatusNotFound, "Transaction not found")
				return
			}
			sendErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Failed to delete transaction: %v", err))
			return
		}

		log.Printf("[INFO] DELETE /transactions/%d - Deleted transaction", id)
		sendSuccessResponse(c, http.StatusOK, gin.H{
			"id": id,
		})
	}
}
