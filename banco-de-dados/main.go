package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	stringConexao := "root:root@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", stringConexao)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		fmt.Println("Dentro do erro ping")
		log.Fatalln(err)
	}
	fmt.Println("Conexão está aberta")

	linhas, err := db.Query("select * from usuarios")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(linhas)
}
