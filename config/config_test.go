package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInitialize(t *testing.T) {
	tests := []struct {
		envVar string
		value  string
		getVar *string
	}{
		{"DB_HOST", "localhost", &DBHost},
		{"DB_USER", "testuser", &DBUser},
		{"DB_PASSWORD", "testpass", &DBPassword},
		{"DB_NAME", "testdb", &DBName},
		{"DB_PORT", "5432", &DBPort},
		{"DB_SSL_MODE", "disable", &DBSSLMode},
		{"APP_SERVER_PORT", "8080", &PORT},
	}

	for _, tt := range tests {
		err := os.Setenv(tt.envVar, tt.value)
		require.NoError(t, err, "failed to set env var: %s", tt.envVar)
	}

	Initialize()

	// Verify that all config variables are set correctly
	for _, tt := range tests {
		require.Equal(t, tt.value, *tt.getVar, "expected %s to be %s", tt.envVar, tt.value)
	}
}
