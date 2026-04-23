package config

import "time"

type ServerConfig struct {
	Env         string // "prod", "staging", "dev"
	Port        int
	ReadTimeout time.Duration
}

type DatabaseConfig struct {
	Host         string
	User         string
	Password     string
	Name         string
	Port         int
	MaxOpenConns int
	MaxIdleConns int
}
