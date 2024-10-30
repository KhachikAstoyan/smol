package db

import (
	"fmt"
	"net/url"
	"smol/core"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB(conf *core.Config) *sqlx.DB {
	dsn := url.URL{
		User:     url.UserPassword(conf.DB.Username, conf.DB.Password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%s", conf.DB.Host, conf.DB.Port),
		Path:     conf.DB.Name,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	conn := sqlx.MustConnect("postgres", dsn.String())
	return conn
}
