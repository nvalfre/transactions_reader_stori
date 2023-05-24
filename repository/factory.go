package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	_ "modernc.org/sqlite"
)

// NewDatabaseRepo creates a new instance of DatabaseRepo
func NewDatabaseRepo() *DatabaseRepo {
	return openSqlite()
	//dsn := "user:password@/dbname"
	//db, err := sql.Open("mysql", dsn) //TODO implement rds.
	//if err != nil {
	//	log.Println(err)
	//	return openSqlite()
	//}
	//
	//err = db.Ping()
	//if err != nil {
	//	log.Println(err)
	//	return openSqlite()
	//}
	//
	//err = executeMockMigrationDDL(db)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//return &DatabaseRepo{
	//	Db: db,
	//}
}

func executeMockMigrationDDL(db *sql.DB) error {
	_, err := db.Exec(queryDDL)
	return err
}

func openSqlite() *DatabaseRepo {
	db, err := sql.Open("sqlite", "transactions.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = executeMockMigrationDDL(db)

	if err != nil {
		log.Fatal(err)
	}
	return &DatabaseRepo{
		Db: db,
	}
}
