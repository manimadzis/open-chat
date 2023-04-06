package interations

import (
	"github.com/jackc/pgx"
	"os"
)

func connect() *pgx.ConnPool {
	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     "localhost",
			Port:     5432,
			Database: "openchat",
			User:     "postgres",
			Password: "postgres",
		},
	},
	)
	if err != nil {
		panic(err)
	}
	return pool
}

func drop(pool *pgx.ConnPool) {
	text, err := os.ReadFile("../../migration/drop_tables.sql")
	if err != nil {
		panic(err)
	}

	if _, err := pool.Exec(string(text)); err != nil {
		panic(err)
	}
}

func create(pool *pgx.ConnPool) {
	text, err := os.ReadFile("../../migration/create_tables.sql")
	if err != nil {
		panic(err)
	}

	if _, err := pool.Exec(string(text)); err != nil {
		panic(err)
	}
}

func load(pool *pgx.ConnPool) {

}
