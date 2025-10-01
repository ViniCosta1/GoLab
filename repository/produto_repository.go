package repository

import (
	"database/sql"
	"github.com/vinicosta1/golab/model"
)

type ProdutoRepository struct {
	// DB é a conexão com o banco de dados.
    // Usamos *sql.DB da biblioteca padrão do Go.
	DB *sql.DB
}

// cria uma nova instância do nosso repositório.
func NewProdutoRepository(db *sql.DB) *ProdutoRepository {
	return &ProdutoRepository{DB: db}
}