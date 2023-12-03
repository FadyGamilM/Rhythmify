package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectToPostgres(t *testing.T) {
	// host := "0.0.0.0:5432"
	// username := "auth_db_user"
	// password := "auth_db_pass"
	// sslmode := false

	dbConn, err := SetupConnection()
	assert.NoError(t, err)
	assert.NotNil(t, dbConn)
}
