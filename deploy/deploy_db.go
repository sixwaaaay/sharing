package main

import (
	"bytelite/etc"
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/zeromicro/go-zero/core/conf"
)

var f = flag.String("conf", "etc/config_test.yaml", "config file")

func main() {
	var c etc.Config
	conf.MustLoad(*f, &c)
	db, _ := sql.Open("mysql", c.DSN)
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://./deploy/app_mysql.sql",
		"test",
		driver,
	)
	m.Steps(2)
}
