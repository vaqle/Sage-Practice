package database

import (
	"database/sql"
	"fmt"
	"github.com/df-mc/dragonfly/server"
	"github.com/go-sql-driver/mysql"
	"go.uber.org/atomic"
	"time"
)

var Srv *server.Server

func Main() {
	initTables()
}

var DATABASE *sql.DB

func initTables() {
	update()
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

var Players atomic.Int32
//so to @prim learned what atomics are
func update() {
	go func() {
		for {
			Players.Store(int32(len(Srv.Players())))
			time.Sleep(1 * time.Second)
		}
	}()
}

func GetDatabase() *sql.DB {
	return DATABASE
}
