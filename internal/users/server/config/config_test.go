package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultConfig(t *testing.T) {
	conf := New()
	assert.Equal(t, conf.IdGen.URL, "localhost:8080")
	assert.Equal(t, conf.UsersServer.TCPPort, 7777)
}

func TestIdGenURL(t *testing.T) {
	os.Setenv("USERS_SERVER_IDGEN_URL", "localhost:8081")
	conf := New()
	assert.Equal(t, conf.IdGen.URL, "localhost:8081")
}

func TestUserServerTCPPort(t *testing.T) {
	os.Setenv("USERS_SERVER_PORT", "7777")
	conf := New()
	assert.Equal(t, conf.UsersServer.TCPPort, 7777)
}
