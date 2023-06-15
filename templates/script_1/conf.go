package main

var config = new(Config)

type Config struct {
	GoCount int    `yaml:"go_count"`
	DbDSN   string `yaml:"db_dsn"`
}
