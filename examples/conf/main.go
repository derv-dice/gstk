package main

import (
	"fmt"
	"github.com/derv-dice/gstk/pkg/conf"
)

type Conf struct {
	Addr  string
	Port  int
	Token string
}

var Config = new(Conf)

func main() {
	err := conf.GenerateConfigTemplate(&Conf{}, conf.YamlConfType)
	if err != nil {
		panic(err)
	}

	err = conf.LoadConfig(Config, conf.YamlConfType)
	if err != nil {
		panic(err)
	}

	fmt.Println(Config)
}
