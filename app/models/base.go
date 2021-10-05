package models

import (
	"Golang_udemy/todo_app/config"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/lib/pq"
	//_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

/*
const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
)
*/

func init() {

	url := os.Getenv("DATABASE_URL")
	url = "postgres://psjgxogklewnxt:0039b9ed7714c7d4edfc4fa45a3a2736c497601087a82a1c5fd6253d39b8369c@ec2-54-204-148-110.compute-1.amazonaws.com:5432/ddt943as20osh3"
	connection, _ := pq.ParseURL(url)
	connection += "sslmode=require"
	Db, err = sql.Open(config.Config.SQLDriver, connection)
	if err != nil {
		log.Fatalln(err)
	}

	/*
		Db, err = sql.Open(config.Config.SQLDriver, config.Config.DBName)
		if err != nil {
			log.Fatalln(err)
		}

		cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			uuid STRING NOT NULL UNIQUE,
			name STRING,
			email STRING,
			password STRING,
			created_at DATETIME)`, tableNameUser)
		Db.Exec(cmdU)

		cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			content TEXT,
			user_id INTEGER,
			created_at DATETIME)`, tableNameTodo)

		Db.Exec(cmdT)

		cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			uuid STRING NOT NULL UNIQUE,
			email STRING,user_id INTEGER,
			created_at DATETIME)`, tableNameSession)

		Db.Exec(cmdS)
	*/
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
