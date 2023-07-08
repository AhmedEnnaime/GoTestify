package config

import (
	"fmt"

	"github.com/AhmedEnnaime/GoTestify/entities"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func ConnectDB(config *entities.Config) *sqlx.DB {
	DB = getDBConnection(config)

	return DB
}

func getDBConnection(config *entities.Config) *sqlx.DB {
	var dbConnectionStr string

	dbConnectionStr = fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		config.Database.Host,
		config.Database.Port,
		config.Database.DbName,
		config.Database.DbName,
		config.Database.Password,
	)

	db, err := sqlx.Open("postgres", dbConnectionStr)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(1)
	db.SetMaxIdleConns(5)

	fmt.Println("Connected to DB")

	return db
}
