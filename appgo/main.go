package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	println("Hello, World!!")
	insertProducts()
}

func insertProducts() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3309)/checkin")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO products (id, name) VALUES ('1', 'Carrinho')")
	if err != nil {
		panic(err.Error())
	}
}