package database

import (
	"test-golang-muhamad-isro/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenConnection(t *testing.T) {

	config := config.InitConfig()

	DATABASE_URL := config.GetString("DATABASE_URL")

	db := OpenConnection(DATABASE_URL)

	defer CloseConnection(db)

	assert.NotNil(t, db)
}
