package main

import (
	"github.com/vinicosta1/golab/handler"
	"github.com/vinicosta1/golab/db"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

var pool *pgxpool.Pool
var router *gin.Engine

func init() {
	pool = db.OpenConn()

	_, err := pool.Exec(db.Ctx, `CREATE TABLE IF NOT EXISTS produtos (
		id SERIAL PRIMARY KEY,
		nome VARCHAR(100) NOT NULL,
		preco NUMERIC(10, 2) NOT NULL
	)`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}

func main() {
	// configuring routes
	router = gin.Default()
	router.POST("/produtos", handler.CreateProduto(pool))
	router.GET("/produtos", handler.GetProdutos(pool))

	router.Run()
}