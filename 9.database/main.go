package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	server   = "localhost"
	port     = 3306
	user     = "go"
	password = "1234"
	database = "go"
)

var connectStr = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
	user, password, server, port, database)

// var pool *sql.DB // Database connection pool.

func main() {
	// get a connection
	db, err := sql.Open("mysql", connectStr)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	// test connection
	err = db.Ping()
	if err != nil {
		log.Panic("Connect to database failure", err)
	}
	fmt.Println("Ping successed!")

	// select one line
	selectOne(1, db)

	// select multiple rows
	var (
		id       int
		name     string
		birthday string
	)
	rows, err := db.Query("select id, name, birthday from t_user where id > ?", 1)
	if err != nil {
		log.Fatal("Got a problem when query multiple rows: ", err)
	}
	for rows.Next() {
		err := rows.Scan(&id, &name, &birthday)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("id: %d, name: %s, birthday: %s\n", id, name, birthday)
	}

	// testInsert(db)

	// testUpdate(db)

	// testDelete(db)

	fmt.Println("Program exited.")
}

func selectOne(queryId int, db *sql.DB) {
	sqlStr := "select id, name, birthday from t_user where id = ?"
	var (
		id       int
		name     string
		birthday string
	)
	err := db.QueryRow(sqlStr, queryId).Scan(&id, &name, &birthday)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("id: %d, name: %s, birthday: %s\n", id, name, birthday)
}

func testInsert(db *sql.DB) {
	today := time.Now()
	rst, err := db.Exec("insert t_user (name, birthday) values(?, ?)", "John Doe", today)
	if err != nil {
		log.Fatal(err)
	}

	id, err := rst.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	selectOne(int(id), db)
}

func testUpdate(db *sql.DB) {
	exec, err := db.Exec("update t_user set name = ? where id = ?", "Caroline Channing", 3)
	if err != nil {
		log.Fatal(err)
		return
	}

	n, err := exec.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %d line(s).\n", n)

	// select
	selectOne(3, db)
}

func testDelete(db *sql.DB) {
	rst, err := db.Exec("delete from t_user where id = ?", 1)
	if err != nil {
		log.Fatal(err)
		return
	}

	n, err := rst.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %d line(s).", n)
}
