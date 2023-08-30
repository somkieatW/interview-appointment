package config

import (
	"time"
)

type Config struct {
	Mysql  Mysql
	Secret Secret
}

type Mysql struct {
	Host         string
	Port         int
	Database     string
	Timeout      time.Duration
	MaxIdleConns int
	MaxOpenConns int
	MaxLifetime  time.Duration
}

type Secret struct {
	MysqlUsername string
	MysqlPassword string
}
