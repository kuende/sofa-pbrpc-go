package discovery

import (
	"os"
	"strings"
)

// EnvProvider implements discovery provider based on env vars
type EnvProvider struct {
	key string
}

// NewEnvProvider returns a new env var discovery provider
func NewEnvProvider(key string) Provider {
	return &EnvProvider{key: key}
}

// Hosts returns a list of for a service
func (e *EnvProvider) Hosts() ([]string, error) {
	value := os.Getenv(e.key)

	return strings.Split(value, ","), nil
}
