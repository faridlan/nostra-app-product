package app

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/faridlan/nostra-api-product/helper"
)

func NewDatabase() *sql.DB {

	helper.LoadEnv()
	port := helper.GetEnvWithKey("PORT_DB")
	host := helper.GetEnvWithKey("HOST_DB")
	name := helper.GetEnvWithKey("NAME_DB")
	pass := helper.GetEnvWithKey("PASS_DB")

	db, err := sql.Open("mysql", fmt.Sprintf("root:%s@tcp(%s:%s)/%s", pass, host, port, name))
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(time.Minute * 10)
	db.SetConnMaxLifetime(time.Minute * 60)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(60)

	return db
}
