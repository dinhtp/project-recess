package server

import "gorm.io/gorm"

type Server interface {
    Serve()
}

func NewServer(db *gorm.DB, address string) Server {
    return &EchoServer{Address: address, db: db}
}
