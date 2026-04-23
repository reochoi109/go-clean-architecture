package config

import (
	"fmt"
	"time"
)

type Config struct {
	// server run options
	Server ServerConfig

	// RDBMS options
	Database DatabaseConfig
}

func Load() *Config {
	_ = LoadEnv("")
	return &Config{
		Server: ServerConfig{
			Env:         EnvString("APP_ENV", "dev"),
			Port:        EnvInt("PORT", 8080),
			ReadTimeout: 5 * time.Second,
		},
		Database: DatabaseConfig{
			Host:         MustEnvString("DB_HOST"),
			User:         MustEnvString("DB_USER"),
			Name:         MustEnvString("DB_NAME"),
			Password:     MustEnvString("DB_PASSWORD"),
			Type:         EnvString("DB_TYPE", "mysql"),
			Port:         MustEnvInt("DB_PORT"),
			MaxOpenConns: EnvInt("DB_MAX_OPEN_CONNS", 10),
			MaxIdleConns: EnvInt("DB_MAX_IDLE_CONNS", 5),
		},
	}
}

func (c *Config) String() string {
	return fmt.Sprintf("Config{Server: [Port: %d], DB: [Host: %s, User: %s, Name: %s, Pass: ****]}",
		c.Server.Port,
		maskString(c.Database.Host),
		maskString(c.Database.User),
		c.Database.Name,
	)
}

func maskString(s string) string {
	if len(s) <= 2 {
		return "****"
	}
	return s[:1] + "***"
}
