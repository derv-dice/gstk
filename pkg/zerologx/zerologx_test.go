package zerologx_test

import (
	log "github.com/derv-dice/gstk/pkg/zerologx"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testMsg = "test msg"

func TestConfigureLogger(t *testing.T) {
	err := log.ConfigureLogger(log.Config{
		Pretty:   true,
		Colored:  true,
		CodeLine: true,
	})
	assert.NoError(t, err)
	log.Info().Msg(testMsg)

	err = log.ConfigureLogger(log.Config{
		Pretty:   true,
		Colored:  true,
		CodeLine: false,
	})
	assert.NoError(t, err)
	log.Info().Msg(testMsg)

	err = log.ConfigureLogger(log.Config{
		Pretty:   true,
		Colored:  false,
		CodeLine: false,
	})
	assert.NoError(t, err)
	log.Info().Msg(testMsg)

	err = log.ConfigureLogger(log.Config{
		Pretty:   false,
		Colored:  false,
		CodeLine: false,
	})
	assert.NoError(t, err)
	log.Info().Msg(testMsg)

	err = log.ConfigureLogger(log.Config{
		Pretty:   false,
		Colored:  false,
		CodeLine: false,
		Level:    log.PanicLevel,
	})
	assert.NoError(t, err)
	log.Info().Msg(testMsg)
}
