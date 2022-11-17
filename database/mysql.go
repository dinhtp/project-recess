package database

import (
    "time"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type MySqlConnector struct {
    Lifetime        time.Duration
    IdleConnections int
    OpenConnections int
    Dsn             string
}

func (c *MySqlConnector) Connect() (*gorm.DB, error) {
    orm, err := gorm.Open(mysql.Open(c.Dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    db, err := orm.DB()
    if nil != err {
        return nil, err
    }

    db.SetConnMaxLifetime(c.Lifetime)
    db.SetMaxIdleConns(c.IdleConnections)
    db.SetMaxOpenConns(c.OpenConnections)

    return orm, nil
}
