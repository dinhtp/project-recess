package database

import (
    "gorm.io/gorm"
)

type Connector interface {
    Connect() (*gorm.DB, error)
}

func NewConnector(dbType, dsn string) Connector {
    switch dbType {
    case DbTypeMySql:
        return &MySqlConnector{
            Lifetime:        DefaultLifeTime,
            IdleConnections: DefaultIdleConnection,
            OpenConnections: DefaultMaxConnection,
            Dsn:             dsn,
        }
    case DbTypeSqLite:
        return &SQLiteConnector{
            Lifetime:        DefaultLifeTime,
            IdleConnections: DefaultIdleConnection,
            OpenConnections: DefaultMaxConnection,
            Dsn:             dsn,
        }
    default:
        return nil
    }
}
