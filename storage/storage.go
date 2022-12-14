package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once // a traves de esta estructura crearemos el patron de diseño Singleton
)

func NewPostgresDB() {
	once.Do(func() {
		// * Lo que este aqui dentro se ejecutará una sola vez
		// Conexion a la DB
		var err error
		db, err = sql.Open("postgres", "postgres://edteam:edteam@localhost:7530/godb?sslmode=disable")
		if err != nil {
			log.Fatalf("can´t open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can´t do ping: %v", err)
		}

		fmt.Println("conectado a postgres")
	})
}

// Pool return a unique instance of db
func Pool() *sql.DB {
	return db
}
