package config

import (
	"github.com/alecthomas/kong"
	"github.com/go-sql-driver/mysql"
)

type Config struct {
	AppAddr    string `env:"APP_ADDR" default:":8080"`
	DBHost     string `env:"NS_MARIADB_HOSTNAME" default:"localhost"`
	DBPort     string `env:"NS_MARIADB_PORT" default:"3306"`
	DBUser     string `env:"NS_MARIADB_USER" default:"root"`
	DBPassword string `env:"NS_MARIADB_PASSWORD" default:"password"`
	DBName     string `env:"NS_MARIADB_DATABASE" default:"mydb"`
}

func (c *Config) MySQLConfig() *mysql.Config {
	return &mysql.Config{
		User:      c.DBUser,
		Passwd:    c.DBPassword,
		Net:       "tcp",
		Addr:      c.DBHost + ":" + c.DBPort,
		DBName:    c.DBName,
		Collation: "utf8mb4_general_ci",
		ParseTime: true,
	}
}

func LoadConfig() *Config {
	var cfg Config

	kong.Parse(&cfg)

	return &cfg
}
