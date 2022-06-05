package main

import (
	"database/sql"
	"log"
	// "fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Person struct {
	Name string
	Age int
}

// database + sqlite3
// テーブル作成

var Db *sql.DB

func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql")
	defer Db.Close()

	// cmd := `CREATE TABLE IF NOT EXISTS person(
	// 			name STRING,
	// 			age INT)`
	
	// _, err := Db.Exec(cmd)

	// cmd := "INSERT INTO person (name, age) VALUES (?, ?)"
	// _, err := Db.Exec(cmd, "tarou", 20)

	// cmd := "UPDATE person SET age = ? WHERE name = ?"
	// _, err := Db.Exec(cmd, 25, "tarou")

	// cmd := "INSERT INTO person (name, age) VALUES (?, ?)"
	// _, err := Db.Exec(cmd, "hanako", 19)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// cmd := "SELECT * FROM person where age = ?"
	// row := Db.QueryRow(cmd, 20)
	// var p Person
	// err := row.Scan(&p.Name, &p.Age)

	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Println("No row")
	// 	} else {
	// 		log.Println(err)
	// 	}
	// }
	// fmt.Println(p.Name, p.Age)

	// cmd := "SELECT * FROM person"
	// rows, _ := Db.Query(cmd)
	// defer rows.Close()
	// var pp []Person
	// for rows.Next() {
	// 	var p Person
	// 	err := rows.Scan(&p.Name, &p.Age)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	pp = append(pp, p)
	// }

	// for _, p := range pp {
	// 	fmt.Println(p.Name, p.Age)
	// }

	cmd := "DELETE FROM person WHERE name = ?"
	_, err := Db.Exec(cmd, "hanako")
	if err != nil {
		log.Fatalln(err)
	}

}