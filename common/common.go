package common

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Config config
	DB     *sql.DB
)

type config struct {
	Env string
	db  string
}

func init() {
	Config = config{
		Env: "dev", // os.Getenv("ENV"),
		db:  "root:testpass@/shipt?parseTime=true",
	}

	DB = getDb()
}

func getDb() *sql.DB {
	log.Println("connecting to DB......")
	db, err := sql.Open("mysql", Config.db)
	if err != nil {
		log.Panicln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Panicln(err)
	}
	log.Println("connected!")
	return db
}
