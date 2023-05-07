package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "Serverh5n"
	dbname   = "celestialposgre"
)

func GetConnection() (db *sql.DB, error error) {
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
	db, err := sql.Open("postgres", "postgresql://postgres:Serverh5n@localhost:5433/celestialposgre?sslmode=disable")
	if err != nil {
		log.Fatalf("could not connect to postgres database: %v", err)
	}

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
