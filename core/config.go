package core

import (
	"os"
	"encoding/json"
)

var Conf *Config

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Db Database `json:"database"`
}

type Database struct {
	Address string `json:"address" form:"address"`
	Port string `json:"port" form:"port"`
	Dbname string `json:"dbname" form:"dbname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func ParseConf(config string) error {
	var c Config

	conf,err := os.Open(config)
	if err!= nil {
		return err
	}
	err = json.NewDecoder(conf).Decode(&c)

	Conf = &c
	return err
}