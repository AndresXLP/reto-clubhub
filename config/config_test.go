package config_test

import (
	"os"
	"sync"
	"testing"

	"franchises-system/config"
	"github.com/stretchr/testify/assert"
)

func TestEnvironments_WhenAllVarsAreSet(t *testing.T) {
	config.Environments()

	assert.Equal(t, "default", config.Cfg.Postgres.DbName)
}

func TestEnvironments_WhenNoVarsAreSet(t *testing.T) {
	errExpected := "Error parsing environment vars &errors.errorString{s:\"Key: 'Config.Server.Port' Error:Field validation for 'Port' failed on the 'required' tag;\\nKey: 'Config.Postgres.Host' Error:Field validation for 'Host' failed on the 'required' tag;\\nKey: 'Config.Postgres.Port' Error:Field validation for 'Port' failed on the 'required' tag;\\nKey: 'Config.Postgres.User' Error:Field validation for 'User' failed on the 'required' tag;\\nKey: 'Config.Postgres.Password' Error:Field validation for 'Password' failed on the 'required' tag;\\nKey: 'Config.Postgres.DbName' Error:Field validation for 'DbName' failed on the 'required' tag\"}"
	config.Once = sync.Once{}

	os.Clearenv()

	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, errExpected, r)
		}
	}()

	config.Environments()
}
