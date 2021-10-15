package query

import (
	_ "mysql_mate_go/test/unit/suite"
)

import (
	"github.com/stretchr/testify/assert"
	"mysql_mate_go/pkg/db"
	"testing"
)

func TestPerformance(t *testing.T) {
	pfData := db.GetPerformance()
	assert.NotNil(t, pfData)
	assert.Greater(t, len(pfData), 0)
	assert.Greater(t, len(pfData["Innodb_buffer_pool_pages_free"]), 0)
	assert.Greater(t, len(pfData["Innodb_buffer_pool_pages_total"]), 0)
}
