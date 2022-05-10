package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

func Main() {
	initTables()
}

var DATABASE *sql.DB

func initTables() {
	db, err := sql.Open("mysql", "db_383699:58a8761227@tcp(na02-db.cus.mc-panel.net:3306)/db_383699")
	if err != nil {
		panic(err)
	}
	_, err = db.Query("CREATE TABLE IF NOT EXISTS sessionPlayers (name VARCHAR(255) PRIMARY KEY, kills BIGINT, deaths BIGINT)")
	DATABASE = db
	if err != nil {
		sqlError := mysql.MySQLError{Number: 1, Message: err.Error()}
		fmt.Println(sqlError.Error())
		return
	}
	fmt.Println("Tables created")
}

func GetDatabase() *sql.DB {
	return DATABASE
}
