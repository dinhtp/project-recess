package connection

import (
    "time"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type SQLiteConnector struct {
    Lifetime        time.Duration
    IdleConnections int
    OpenConnections int
    Dsn             string
}

func (c *SQLiteConnector) Connect() (*gorm.DB, error) {
    orm, err := gorm.Open(sqlite.Open(c.Dsn), &gorm.Config{})
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
