# gstk - golang scripts tool kit

Всякие полезные приколюхи, ускоряющие и облегчающие написание скриптов почти как с питоном, но без дебильных лесенок.

Доступные пакеты:
- `go get github.com/derv-dice/gstk/pkg/webpb` - Прогрессбар в браузере
- `go get github.com/derv-dice/gstk/pkg/conf` - Генерация и парсинг json и yaml конфигов
- `go get github.com/derv-dice/gstk/pkg/wpool` - workers pool Выполнение задач в несколько потоков
- `go get github.com/derv-dice/gstk/pkg/iox` - Утилиты для ввода/вывода. Чтение, запись файлов различных форматов (txt, csv, xlsx, ...)
- `go get github.com/derv-dice/gstk/pkg/pgdb` - Работа с БД PostgreSQL (`sqlx`+`pgx`)
- `go get github.com/derv-dice/gstk/pkg/zerologx` - Обвязка для работы с [zerolog](https://github.com/rs/zerolog)
- `go get github.com/derv-dice/gstk/pkg/interval` - Функции для работы с периодами и интервалами. В основном, нужно для разбиения sql запросов по
  времени

В планах:
- `pkg/mongodb` - Работа с БД MongoDB
- `pkg/fdb` - Работа с БД FoundationDB
