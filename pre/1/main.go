package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	connStr := "dbname=20241029c sslmode=disable host=localhost"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Mengatur parameter pool koneksi
	//db.SetMaxOpenConns(25)                  // Jumlah maksimal koneksi terbuka
	//db.SetMaxIdleConns(25)                  // Jumlah maksimal koneksi idle
	//db.SetConnMaxLifetime(30 * time.Minute) // Durasi maksimal koneksi (misalnya, 30 menit)
	//db.SetConnMaxIdleTime(5 * time.Minute)  // Durasi maksimal koneksi idle (misalnya, 5 menit)

	// Test the connection to the database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}
}
