package main

import (
	// "database/sql"
	// "log"
	"fmt"

	// _ "github.com/mattn/go-sqlite3"
	// _ "github.com/lib/pq"
	// "gopkg.in/go-ini/ini.v1"
	"github.com/google/uuid"
)

// type Person struct {
// 	Name string
// 	Age int
// }

// database + sqlite3
// テーブル作成

// var Db *sql.DB

// var err error

// type ConfigList struct {
// 	Port		int
// 	DbName		string
// 	SQLDriver	string
// }

// var Config ConfigList

// func init() {
// 	cfg, err := ini.Load("config.ini")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	Config = ConfigList{
// 		Port: cfg.Section("web").Key("port").MustInt(8080),

// 		DbName: cfg.Section("db").Key("name").MustString("example.sql"),

// 		SQLDriver: cfg.Section("db").Key("driver").String(),
// 	}
// }

func main() {
	uuidObj, _ := uuid.NewUUID()
	fmt.Println("  ", uuidObj.String())

	uuidObj2, _ := uuid.NewRandom()
	fmt.Println("  ", uuidObj2.String())


	// fmt.Printf("Port = %v\n", Config.Port)
	// fmt.Printf("DbName = %v\n", Config.DbName)
	// fmt.Printf("SQLDriver = %v\n", Config.SQLDriver)


	// Db, err := sql.Open("postgres", "user=test_user dbname=testdb password=password sslmode=disable")
	// if err != nil {
	// 	log.Println(err)
	// }
	// defer Db.Close()

	// C
	// cmd := "INSERT INTO persons (name, age) VALUES ($1, $2)"
	// _, err2 := Db.Exec(cmd, "Nancy", 20)
	// if err2 != nil {
	// 	log.Fatalln(err2)
	// }

	// R
	// cmd := "SELECT * FROM persons where age = $1"
	// row := Db.QueryRow(cmd, 20)
	// var p Person
	// err = row.Scan(&p.Name, &p.Age)

	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Println("No row")
	// 	} else {
	// 		log.Println(err)
	// 	}
	// }
	// fmt.Println(p.Name, p.Age)

	// cmd = "SELECT * FROM persons"
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

	// U
	// cmd := "UPDATE persons SET age = $1 WHERE name = $2"
	// _, err = Db.Exec(cmd, 25, "Nancy")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// D
	// cmd := "DELETE FROM persons WHERE name = $1"
	// _, err = Db.Exec(cmd, "Nancy")
	// if err != nil {
	// 	log.Fatalln(err)
	// }


	// Db, _ := sql.Open("sqlite3", "./example.sql")
	// defer Db.Close()

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

	// cmd := "DELETE FROM person WHERE name = ?"
	// _, err := Db.Exec(cmd, "hanako")
	// if err != nil {
	// 	log.Fatalln(err)
	// }


}