package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// accessing password from environment
	pwd := os.Getenv("MYSQL_PASSWORD")
	if pwd == "" {
		fmt.Println("not accessed")
	}
	fmt.Println("password is : ", pwd)
	// making connection with db
	// sql.Open("drivername", "datasourcename")
	// datasourcename = username:password@tcp(host:port)/dbname
	db, err := sql.Open("mysql", "root:"+pwd+"@tcp(127.0.0.1:3306)/godb")
	defer db.Close()
	if err != nil {
		fmt.Println("connection")
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("connection viability error")
		log.Fatal(err)
	}
	fmt.Println("connection established")
	createTable(db)
	insertData(db)
	selectRow(db)
}

func createTable(db *sql.DB) {
	query := `CREATE TABLE employee (
	id INT PRIMARY KEY,
	firstname VARCHAR(50) NOT NULL,
	lastname VARCHAR(50) NOT NULL,
	mobile INT,
	address VARCHAR(100)
)`

	create, err := db.Query(query)
	defer create.Close()
	if err != nil {
		fmt.Println("Error in creating table")
		log.Fatal(err)
	}
	fmt.Println("Employee table created successfully")
}

func insertData(db *sql.DB) {
	query := `INSERT INTO employee(id, firstname, lastname, mobile, address)
	VALUES (1, "Ram", "Sharma", 0123456789, "Gautam nagar delhi")`
	insert, err := db.Query(query)
	defer insert.Close()
	if err != nil {
		fmt.Println("error in inserting data")
		log.Fatal(err)
	}
	fmt.Println("inserted successfully")
}
func selectRow(db *sql.DB) {
	query := `SELECT * FROM employee`
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		fmt.Println("error in selecting row")
		log.Fatal(err)
	}
	fmt.Println("selected successfully")
}
