package con

import (
	"github.com/jmoiron/sqlx"
	_ "gopkg.in/goracle.v2"
)

func ConnectSql(driver, datasource string) *sqlx.DB {
	db, err := sqlx.Open(driver, datasource)
	if err != nil {
		panic(err)
	}

	return db
}