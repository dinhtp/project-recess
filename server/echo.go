package server

import (
    "github.com/labstack/echo/v4"
    "gorm.io/gorm"
)

type EchoServer struct {
    db      *gorm.DB
    Address string
}

func (s *EchoServer) Serve() {
    server := echo.New()
    server.Logger.Fatal(server.Start(":8080"))
}
