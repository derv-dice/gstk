package zerologx

import (
	"github.com/rs/zerolog"
	"io"
	"os"
	"path"
	"time"
)

const (
	TraceLevel = zerolog.TraceLevel
	DebugLevel = zerolog.DebugLevel
	InfoLevel  = zerolog.InfoLevel
	WarnLevel  = zerolog.WarnLevel
	ErrorLevel = zerolog.ErrorLevel
	FatalLevel = zerolog.FatalLevel
	PanicLevel = zerolog.PanicLevel

	_defaultLogsDir = "./logs"
)

// Config - Конфигуратор логов
// - Level: Уровень логов. По умолчанию - DEBUG
// - Pretty: В stdout будет выводиться лог формате '1:01PM LEVEL MESSAGE'
// - Colored: В stdout будет выводиться цветной лог (Работает только если Pretty=true)
// - ToFile: Писать ли логи в файл
// - Dir: В какую директорию положить файл с логами (Работает только если ToFile=true). По умолчанию - ./logs/
// - CodeLine: Логирование строки кода
type Config struct {
	Level    zerolog.Level
	Pretty   bool
	Colored  bool
	ToFile   bool
	Dir      string
	CodeLine bool
}

func ConfigureLogger(conf Config) (err error) {
	var writers []io.Writer

	if conf.ToFile {
		if conf.Dir == "" {
			conf.Dir = _defaultLogsDir
		}

		if _, err = os.Stat(conf.Dir); os.IsNotExist(err) {
			if err = os.MkdirAll(conf.Dir, 0775); err != nil {
				return
			}
		}

		var f io.Writer
		if f, err = os.Create(path.Join(conf.Dir, time.Now().Format("2006-01-02_15:04:05.log"))); err != nil {
			return
		}

		writers = append(writers, f)
	}

	if conf.Pretty {
		cw := zerolog.ConsoleWriter{Out: os.Stdout}
		if !conf.Colored {
			cw.NoColor = true
		}

		writers = append(writers, cw)
	} else {
		writers = append(writers, os.Stdout)
	}

	Logger = zerolog.New(zerolog.MultiLevelWriter(writers...)).Level(conf.Level).With().Timestamp().Logger()

	if conf.CodeLine {
		Logger = Logger.With().Caller().Logger()
	}

	return
}
