package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "qweqwe"
	dbname   = "postgres"
)

type vpnRepImpl struct {
	db *sql.DB
}

func NewvpnRepImpl() *vpnRepImpl {
	db := connectDB()
	rep := &vpnRepImpl{db: db}
	return rep
}

func connectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println("db connection error")
		return nil
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Cannot ping db")
	} else {
		fmt.Println("db ping ping")
	}

	return db
}

func (*vpnRepImpl) Get() {

}
func (*vpnRepImpl) Post() {

}
func (*vpnRepImpl) Delete() {

}
func (*vpnRepImpl) Update() {

}
