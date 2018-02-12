package lib

import (
	"log"

	"upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

var config = mysql.ConnectionURL{
	Host:     "localhost",
	User:     "zabbix",
	Password: "saciqwe123",
	Database: "membros",
}

var Connection db.Database

func init() {
	var erro error

	Connection, erro = mysql.Open(config)
	if erro != nil {
		log.Fatal(erro.Error())
	}
}
