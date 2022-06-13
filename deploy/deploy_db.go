package main

import (
	"bytelite/etc"
	"database/sql"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"io/ioutil"
	"os"
)

// MustLoadSqlScript  loads the sql script from the file.
func MustLoadSqlScript(kind string) string {
	var fileName string
	switch kind {
	case "mysql":
		fileName = "./deploy/app_mysql.sql"
	case "TiDB":
		fileName = "./deploy/app_TiDB.sql"
	default:
		panic("unknown kind")
	}
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	// read all content
	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	// execute the content
	return string(content)
}

var f = flag.String("conf", "etc/config.yaml", "config file")
var kind = flag.String("kind", "mysql", "mysql or TiDB")

func main() {
	var c etc.Config
	conf.MustLoad(*f, &c)
	conn := sqlx.NewMysql(c.DSN)
	db, err := conn.RawDB()
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
	// load sql script
	script := MustLoadSqlScript(*kind)
	// execute the script
	_, err = db.Exec(script)
	if err != nil {
		panic(err)
	}
}
