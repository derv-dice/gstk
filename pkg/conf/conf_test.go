package conf

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

type Foo struct {
	A string `json:"a" yaml:"a"`
	B int    `json:"b" yaml:"b"`
	C bool   `json:"c" yaml:"c"`
	D struct {
		E time.Time         `json:"e" yaml:"e"`
		F map[string]string `json:"f" yaml:"f"`
		G []int             `json:"g" yaml:"g"`
	} `json:"d" yaml:"d"`
}

func cleanAfterTests() (err error) {
	if err = os.Remove("config.json"); err != nil {
		return
	}

	if err = os.Remove("config.yaml"); err != nil {
		return
	}

	return
}

func TestGenerateConfigTemplate(t *testing.T) {
	t.Run("generate_json_tmpl", func(t *testing.T) {
		assert.NoError(t, GenerateConfigTemplate(&Foo{}, JsonConfType))
		assert.NoError(t, os.Rename("config.tmpl.json", "config.json"))
	})

	t.Run("generate_yaml_tmpl", func(t *testing.T) {
		assert.NoError(t, GenerateConfigTemplate(&Foo{}, YamlConfType))
		assert.NoError(t, os.Rename("config.tmpl.yaml", "config.yaml"))
	})
}

func TestLoadConfig(t *testing.T) {
	f := Foo{}

	t.Run("load_json_config", func(t *testing.T) {
		assert.NoError(t, LoadConfig(&f, JsonConfType))
	})

	t.Run("load_yaml_config", func(t *testing.T) {
		assert.NoError(t, LoadConfig(&f, YamlConfType))
	})

	t.Run("clean_after_tests", func(t *testing.T) {
		assert.NoError(t, cleanAfterTests())
	})
}
