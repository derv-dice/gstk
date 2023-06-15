package main

import (
	"github.com/derv-dice/gstk/pkg/conf"
	"github.com/derv-dice/gstk/pkg/pgdb"
	log "github.com/derv-dice/gstk/pkg/zerologx"
)

func init() {
	var err error
	err = log.ConfigureLogger(log.Config{Level: log.DebugLevel, Pretty: true, Colored: true, ToFile: true})
	if err != nil {
		log.Fatal().Err(err).Msg("Ошибка конфигурации логгера")
	}

	if err = conf.GenerateConfigTemplate(Config{}, conf.YamlConfType); err != nil {
		log.Fatal().Err(err).Msg("Ошибка генерации шаблона конфига")
	}
}

func main() {
	var err error

	log.Debug().Msg("Загрузка конфига")
	if err = conf.LoadConfig(config, conf.YamlConfType); err != nil {
		log.Fatal().Err(err).Msg("Ошибка загрузки конфига")
	}

}
