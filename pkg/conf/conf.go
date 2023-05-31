package conf

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	YamlConfType = iota
	JsonConfType

	defaultTmplFilename = "config.tmpl"
	defaultFilename     = "config"
)

type FileType int

func LoadConfig(dst any, confType int) (err error) {
	if dst == nil {
		return errors.New("LoadConfig: dst is nil")
	}

	switch confType {
	case JsonConfType:
		return loadJsonConf(dst)
	case YamlConfType:
		return loadYamlConf(dst)
	default:
		return fmt.Errorf("GenerateConfigTemplate: incorrect confType %d", confType)
	}
}

func loadJsonConf(dst any) (err error) {
	var r io.Reader
	if r, err = os.Open(defaultFilename + ".json"); err != nil {
		return
	}

	e := json.NewDecoder(r)
	return e.Decode(dst)
}

func loadYamlConf(dst any) (err error) {
	var r io.Reader
	if r, err = os.Open(defaultFilename + ".yaml"); err != nil {
		return
	}

	e := yaml.NewDecoder(r)
	return e.Decode(dst)
}

func GenerateConfigTemplate(confStruct any, confType int) (err error) {
	if confStruct == nil {
		return errors.New("GenerateConfigTemplate: confStruct is nil")
	}

	switch confType {
	case JsonConfType:
		return generateJsonTmpl(confStruct)
	case YamlConfType:
		return generateYamlTmpl(confStruct)
	default:
		return fmt.Errorf("GenerateConfigTemplate: incorrect confType %d", confType)
	}
}

func generateJsonTmpl(confStruct any) (err error) {
	var w io.Writer
	if w, err = os.Create(defaultTmplFilename + ".json"); err != nil {
		return
	}

	e := json.NewEncoder(w)
	e.SetIndent("", "\t")
	return e.Encode(confStruct)
}

func generateYamlTmpl(confStruct any) (err error) {
	var w io.Writer
	if w, err = os.Create(defaultTmplFilename + ".yaml"); err != nil {
		return
	}

	e := yaml.NewEncoder(w)
	e.SetIndent(2)
	return e.Encode(confStruct)
}
