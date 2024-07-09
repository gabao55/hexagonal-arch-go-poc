package main

import (
	"database/sql"
	"github.com/gabao55/hexagonal-arch-go/application"
	"github.com/gabao55/hexagonal-arch-go/adapters/db"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	Db, _ := sql.Open("sqlite3", "db.sqlite3")
	productDbAdapter := db.NewProductDb(Db)
	productService := application.NewProductService(productDbAdapter)

	product, _ := productService.Create("Example Product", 50)
	productService.Enable(product)
}