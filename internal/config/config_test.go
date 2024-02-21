package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
  t.Run("Success", func (t *testing.T) {
    filename := "../../cmd/api/config.yaml"
    err := LoadConfig(filename)

    require.Nil(t, err)
  })
}
