package config

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/go_simple_ecommerce?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Gagal membuka koneksi database %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Database tidak merespon: %v", err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute * 5)

	log.Println("Berhasil terhubung ke database MySQL!")

	return db

}
