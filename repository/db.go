package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func getInstance() *sql.DB {
	const (
		host     = "ec2-52-203-118-49.compute-1.amazonaws.com"
		port     = "5432"
		user     = "jvysslbstxwefn"
		password = "f5eaea64092a2fcb94b231c6c7e8a47f7493c4f5bfd6d8d98bc0b5561d17f82a"
		dbname   = "d3hms1ojq1sklj"
	)
	if db == nil {
		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbname)
		res, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}
		db = res
	}
	return db
}
