package database

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	// db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_database")
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "Serverh5n&*#"
	dbName := "celestialdevtest"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// gunakan DB
}
