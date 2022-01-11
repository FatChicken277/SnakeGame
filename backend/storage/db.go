package storage

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// DBConection is used to establish connection with the database
func DBConection(source string) (*pgx.Conn, error) {
	config, err := pgx.ParseConfig(source)
	if err != nil {
		return nil, err
	}

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
