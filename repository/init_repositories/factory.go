package init_repositories

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	_ "modernc.org/sqlite"
)

// NewDatabaseRepo creates a new instance of DatabaseRepo
func NewDatabaseRepo(db *sql.DB) *DatabaseRepo {
	return openSqlite(db)
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

func openSqlite(db *sql.DB) *DatabaseRepo {
	err := executeMockMigrationDDL(db)

	if err != nil {
		log.Fatal(err)
	}
	return &DatabaseRepo{
		Db: db,
	}
}
