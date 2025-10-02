package handler

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vinicosta1/golab/data"
	"github.com/gin-gonic/gin"
)

func CreateProduto(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var produto data.Produto
		sql := `
            INSERT INTO produtos (nome, preco)
			VALUES ($1, $2)
			RETURNING id
        `
		
		err := c.ShouldBindJSON(&produto)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		err = pool.QueryRow(context.Background(), sql, produto.Nome, produto.Preco).Scan(&produto.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		
		c.JSON(http.StatusCreated, gin.H{"data": produto})
	}
}

func GetProdutos(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		sql := `SELECT id, nome, preco FROM produtos`
		
		rows, err := pool.Query(context.Background(), sql)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var produtos []data.Produto
		for rows.Next() {
			var produto data.Produto
			err := rows.Scan(
				&produto.ID,
				&produto.Nome,
				&produto.Preco,
			)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}

			produtos = append(produtos, produto)
		}

		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"data": produtos})
	}
}