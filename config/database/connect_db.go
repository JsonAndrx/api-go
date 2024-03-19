package database

import (
	"api-rest/helper"

	// "time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

)

func ConectDb() (*sqlx.DB){
    db, err := sqlx.Connect("mysql", "root:@/doublev")
	if err != nil {
		helper.ErrorPanic(err)
	}

    return db
}
