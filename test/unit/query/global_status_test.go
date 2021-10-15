package query

import (
	_ "mysql_mate_go/test/unit/suite"
)

import (
	"github.com/stretchr/testify/assert"
	"mysql_mate_go/pkg/db"
	"testing"
)

func TestGetGlobalStatus(t *testing.T) {
	globalStatus := db.GetGlobalStatus()
	assert.NotNil(t, globalStatus)
	assert.Greater(t, len(globalStatus), 0)
}
