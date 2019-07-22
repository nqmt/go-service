package con

import (
	"github.com/jmoiron/sqlx"
	_ "gopkg.in/goracle.v2"
)

func ConnectOracle() *sqlx.DB {
	db, err := sqlx.Open("goracle", "SYSTEM/password@127.0.0.1:1521/TEST01")
	if err != nil {
		panic(err)
	}

	return db
}