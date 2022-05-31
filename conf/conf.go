package conf

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type webConf struct {
	Address string `yaml:"address"`
}

type mysqlConf struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Db       string `yaml:"db"`
}

type conf struct {
	Web   webConf
	Mysql mysqlConf
}

var Config conf

func LoadConf() error {
	confFile, err := os.Open("conf/conf.yaml")
	if err != nil {
		log.Fatal(err)
	}
	yamlDecoder := yaml.NewDecoder(confFile)
	err = yamlDecoder.Decode(&Config)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
