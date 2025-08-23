package app

import (
	"web/example/internal/db"
	"web/example/internal/http"

	"go.uber.org/zap"
)

func Launch() {
	zap.S().Info("App lauched")
	db_conn, err := db.NewSQLite3()

	if err != nil {
		zap.S().Errorln("Could not connect to DB: ", err.Error())
	}

	http.StartServer(db_conn)
}
