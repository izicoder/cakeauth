package core

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func TestDB() {
	cfg, _ := pgx.ParseConfig("user=postgres password=admin host=localhost port=5432 dbname=appdb")
	conn, err := pgx.ConnectConfig(context.Background(), cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	var greeting string
	err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)

	createTable(conn)
}

func createTable(conn *pgx.Conn) {
	rows, err := conn.Query(context.Background(), "create table users (id int primary key, name text, password text)")
	if err != nil {
		fmt.Printf("cant create table %v\n", err)
	} else {
		values, err := rows.Values()
		if err != nil {
			fmt.Println("bruh")
		}
		fmt.Println("table created ")
		fmt.Println(values...)
	}
}
