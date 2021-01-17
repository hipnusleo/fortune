package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

const defaultOptions = `
{
	"mysqlOptions":{
		"user":"johndoe",
		"password":"Qwer@1234",
		"host":"localhost",
		"port":"3306",
		"database":"default_db",
		"tb_prefix":"default"
	}
	"redisOptions":{
		"host":"localhost",
		"port":"6379",
	}
	"influxdbOptions":{
		"host":"localhost",
		"port":"8086",
		"token":"my-token"
	}
}
`

var defaultPathToConfig string
var defaultConfigFile = "conf.json"
var cfg = &Config{Type: "JSON"}
var options = &Options{}

type Config struct {
	Type         string // JSON / YAML /Ini
	PathToConfig string
}

type Options struct {
	MysqlOption    `json:mysqlOptions`
	RedisOption    `json:redisOptions`
	InfluxdbOption `json:influxdbOptions`
}

type MysqlOption struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	TBPrefix string `json:"tb_prefix"`
}

type RedisOption struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

type InfluxdbOption struct {
	Host  string `json:"host"`
	Port  string `json:"port"`
	Token string `json:"token"`
}

func init() {
	// find $PWD
	pwd, err := os.Getwd()
	if err != nil {
		logrus.Fatalf("failed to get pwd: %s", err)
	}
	defaultPathToConfig = path.Join(pwd, defaultConfigFile)
}

func loadJSONConfig() {
	if cfg.PathToConfig == "" {
		cfg.PathToConfig = defaultPathToConfig
	}
	content, err := ioutil.ReadFile(cfg.PathToConfig)
	if err != nil {
		logrus.Errorf("encounter errors when reading %s: %s", cfg.PathToConfig, err)
		return
	}
	err = json.Unmarshal(content, options)
	if err != nil {
		logrus.Errorf("failed to unmarshal config file: %s", err)
	}
	logrus.Info(options)
}

func loadYAMLConfig() {

}

func InitConfig() {
	switch cfg.Type {
	case "JSON":
		loadJSONConfig()

	case "YAML":
		loadYAMLConfig()

	default:
		logrus.Error("incorrect config file")
	}
}
