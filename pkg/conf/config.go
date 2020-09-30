package conf

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Db struct {
	DbType     string `yaml:"db_type"`
	DbUser     string `yaml:"db_user"`
	DbPassword string `yaml:"db_password"`
	DbHost     string `yaml:"db_host"`
	DbPort     string `yaml:"db_port"`
	DbName     string `yaml:"db_name"`
}

type Redis struct {
	RedisPrefix      string        `yaml:"redis_prefix"`
	RedisHost        string        `yaml:"redis_host"`
	RedisPassword    string        `yaml:"redis_password"`
	RedisDB          int           `yaml:"redis_db"`
	RedisMaxIdle     int           `yaml:"redis_max_idle"`
	RedisMaxActive   int           `yaml:"redis_max_active"`
	RedisIdleTimeout time.Duration `yaml:"redis_idle_timeout"`
}

type Config struct {
	// server
	Mode         string        `yaml:"mode"`
	Port         int           `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_time_out"`
	WriteTimeout time.Duration `yaml:"write_time_out"`

	// database
	Db Db `yaml:"db"`

	// redis
	Redis Redis `yaml:"redis"`
}

var GlobalConfig = Config{}

func SetupConf() {
	file, err := os.Open("app.yaml")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'app.yaml': %v", err)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("fail to read 'app.yaml': %v", err)
	}
	if err := yaml.Unmarshal([]byte(content), &GlobalConfig); err != nil {
		log.Fatalf("fail to yaml Unmarshal config: %v", err)
	}
}