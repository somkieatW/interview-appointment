package db

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/somkieatW/candidate/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mySqlDB struct {
	Client *gorm.DB
}

func NewMySqlDB(cfg config.Mysql, secrets config.Secret) (*mySqlDB, error) {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?checkConnLiveness=false&loc=Local&parseTime=true&readTimeout=%s&timeout=%s&writeTimeout=%s&maxAllowedPacket=0",
		secrets.MysqlUsername,
		secrets.MysqlPassword,
		addr,
		cfg.Database,
		cfg.Timeout,
		cfg.Timeout,
		cfg.Timeout,
	)
	//log.Infof(context.Background(), "[CONFIG] [MYSQL] address: %s", addr)

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, errors.Wrap(err, "gorm open")
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.Wrap(err, "get db")
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.MaxLifetime)

	return &mySqlDB{db}, nil
}

func (db *mySqlDB) Close() error {
	mysql, err := db.Client.DB()
	if err != nil {
		return errors.Wrap(err, "client db")
	}
	if err := mysql.Close(); err != nil {
		return errors.Wrap(err, "mysql close")
	}
	return nil
}

func (db *mySqlDB) Ping() error {
	sql, err := db.Client.DB()
	if err != nil {
		return errors.Wrap(err, "mysql")
	}
	err = sql.Ping()
	return errors.Wrap(err, "mysql")
}
