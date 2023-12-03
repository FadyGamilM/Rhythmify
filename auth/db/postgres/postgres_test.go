package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectToPostgres(t *testing.T) {
	dbConn, err := SetupConnection()
	assert.NoError(t, err)
	assert.NotNil(t, dbConn)
}
