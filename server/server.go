package server

import (
    "github.com/dinhtp/project-recess/database"
    "gorm.io/gorm"
)

type Server interface {
    Serve()
}

func NewServer(db *gorm.DB, address, dbType string) Server {
    switch dbType {
    case database.DbTypeSqLite:
        return &EchoServer{Address: address, db: db}
    case database.DbTypeMySql:
        return &MuxServer{Address: address, db: db}
    default:
        return nil
    }
}
