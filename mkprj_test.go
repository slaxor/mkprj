package mkprj

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_mkprj(t *testing.T) {
  err := mkprj()
  assert.NoError(t, err)
}
