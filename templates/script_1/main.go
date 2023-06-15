package main

import (
	"context"
	"github.com/derv-dice/gstk/pkg/conf"
	"github.com/derv-dice/gstk/pkg/iox"
	"github.com/derv-dice/gstk/pkg/pgdb"
	"github.com/derv-dice/gstk/pkg/wpool"
	log "github.com/derv-dice/gstk/pkg/zerologx"
	"github.com/jmoiron/sqlx"
)

var dbFirst *sqlx.DB

func main() {
	var err error

	err = log.ConfigureLogger(log.Config{Level: log.DebugLevel, Pretty: true, Colored: true, ToFile: true})
	if err != nil {
		log.Fatal().Err(err).Msg("Ошибка конфигурации логгера")
	}

	log.Debug().Msg("Генерация шаблона конфига")
	if err = conf.GenerateConfigTemplate(Config{}, conf.YamlConfType); err != nil {
		log.Fatal().Err(err).Msg("Ошибка генерации шаблона конфига")
	}

	log.Debug().Msg("Загрузка конфига")
	if err = conf.LoadConfig(config, conf.YamlConfType); err != nil {
		log.Fatal().Err(err).Msg("Ошибка загрузки конфига")
	}

	log.Debug().Msgf("Подключение к БД %s", config.DbFirstName)
	if dbFirst, err = pgdb.Pool().AddConn(config.DbFirstName, config.DbFirstDSN); err != nil {
		log.Fatal().Err(err).Msgf("Ошибка подключения к БД %s", config.DbFirstName)
	}
	defer pgdb.Pool().CloseAll()

	log.Debug().Msgf("Чтение файла %s", config.InputFilename)
	var ids []string
	if ids, err = iox.ReadFileLineByLine(config.InputFilename, true); err != nil {
		log.Fatal().Err(err).Msgf("Ошибка чтения файла %s", config.InputFilename)
	}

	log.Debug().Int("go_count", config.GoCount).Msg("Инициализация wPool")
	wPool := wpool.NewWPool(context.Background(), config.GoCount).Start()
	defer wPool.Stop(false)

	log.Debug().Msg("Запущена обработка задач")
	for _, id := range ids {
		err = wPool.Put(&WpTask{
			Id: id,
		})

		if err != nil {
			log.Fatal().Err(err).Msgf("Ошибка постановки задачи в wPool")
		}
	}
}

type WpTask struct {
	Id string
}

func (c *WpTask) Do(ctx context.Context) {
	// todo: some work
}
