package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/celpostgress-api/utils"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "Serverh5n"
	dbname   = "celestialposgre"
)

func GetConnection(cfg utils.Config) (db *sql.DB, error error) {
	// dbDriver := "postgress"
	// dbUser := "root"
	// dbPass := "Serverh5n&*#"
	// dbName := "celestiallivedb"
	// // datasource := "root:@tcp(localhost:3306)/celestialdev?parseTime=true"
	// // db, err := sql.Open("mysql",dataSource)
	// db, err := sql.Open("postgres", psqlInfo)
	// // db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// return db, err
	// postgresql://admin:password123@localhost:6500/golang_postgres?sslmode=disable
	pgDSN := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.PgDBUser, cfg.PgDBPass, cfg.PgDBHost, cfg.PgDBPort, cfg.PgDBName)
	db, err := sql.Open("postgres", pgDSN)
	if err != nil {
		log.Fatalf("could not connect to postgres database: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("could not connect to %s %v", pgDSN, err)
		return nil, err
	}

	log.Default().Print("database connected to", pgDSN)

	// db = dbConn.New(conn)

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)
	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }
	return db, err
}

// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// "password=%s dbname=%s sslmode=disable",
// host, port, user, password, dbname)
// db, err := sql.Open("postgres", psqlInfo)
// if err != nil {
// panic(err)
// }
// defer db.Close()

// err = db.Ping()
// if err != nil {
// panic(err)
// }

// fmt.Println("Successfully connected!")

func GetConnectionTest() (db *sql.DB, error error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "Serverh5n&*#"
	dbName := "celestialdev"
	// datasource := "root:@tcp(localhost:3306)/celestialdev?parseTime=true"
	// db, err := sql.Open("mysql",dataSource)
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db, err
}

func GetConnectionDev() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "Serverh5n&*#"
	dbName := "celestialdev"
	// datasource := "root:@tcp(localhost:3306)/celestialdev?parseTime=true"
	// db, err := sql.Open("mysql",dataSource)
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}
