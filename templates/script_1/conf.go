package main

var config = new(Config)

type Config struct {
	// Количество горутин
	GoCount int `yaml:"go_count"`

	// Подключения к БД
	DbFirstName string `yaml:"db_first_name"`
	DbFirstDSN  string `yaml:"db_first_dsn"`

	InputFilename string `yaml:"input_filename"`
}
