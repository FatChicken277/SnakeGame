package storage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDBConection(t *testing.T) {
	c := require.New(t)

	dbSource := "postgresql://admin@localhost:26257/snake_test?sslmode=disable"
	dbConn, err := DBConection(dbSource)
	c.NoError(err)
	dbConn.Close(context.Background())
}
