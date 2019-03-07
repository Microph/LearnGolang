package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql" // import with _ for only register driver
)

func main() {
	ctx := context.Background()
	check := func(err error) {
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
	// Open connection to MySQL Database
	db, err := sql.Open("mysql", "s4OU7DIrDc:tg5lI31Ypn@tcp(remotemysql.com)/s4OU7DIrDc?parseTime=true")
	check(err)

	//-----------------------------------------------------

	//insert
	stmt, err := db.PrepareContext(ctx, "INSERT INTO posts(title, body) values (?,?)")
	check(err)

	result, err := stmt.ExecContext(ctx, "abx", "sadfdf sf.")
	check(err)
	lastID, _ := result.LastInsertId()
	fmt.Println("New Record ID:", lastID)
	//insert

	//update
	stmt, err := db.PrepareContext(ctx, "UPDATE posts SET title = ? WHERE id = ?")
	check(err)

	_, err = stmt.ExecContext(ctx, "เรียนเขียน Go", 1)
	check(err)
	//update

	//select
	qry := "SELECT id, title, body, created_at, updated_at FROM posts"
	stmt, err = db.PrepareContext(ctx, qry)
	check(err)

	rows, err := stmt.QueryContext(ctx)
	check(err)

	for rows.Next() {
		var (
			id        int
			title     string
			body      string
			createdAt time.Time
			updatedAt time.Time
		)
		err := rows.Scan(&id, &title, &body, &createdAt, &updatedAt)
		check(err)
		fmt.Println(id, title, body, createdAt.Format(time.RFC3339), updatedAt.Format(time.RFC3339))
	}
	//select
}
