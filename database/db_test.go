package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenConnection(t *testing.T) {

	db := OpenConnection()

	defer CloseConnection(db)

	assert.NotNil(t, db)
}
