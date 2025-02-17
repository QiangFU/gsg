package gcp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigPath(t *testing.T) {
	t.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "test_path")
	assert.Equal(t, "test_path", ConfigPath())
}

func TestEuqalCRC32C(t *testing.T) {
	assert.True(t, equalCRC32C("invalid", "invalid", "invalid"))
	assert.False(t, equalCRC32C("gcp.go", "invalid", "invalid"))
	// assert.True(t, equalCRC32C("usa.geojson", "maaas", "borders/usa.geojson"))
	// assert.False(t, equalCRC32C("invalid", "maaas", "borders/usa.geojson"))
}
