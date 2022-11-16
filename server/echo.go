package server

import (
    echoCtrl "github.com/dinhtp/project-recess/controller/echo"
    "github.com/labstack/echo/v4"
    "gorm.io/gorm"
)

type EchoServer struct {
    db      *gorm.DB
    Address string
}

func (s *EchoServer) Serve() {
    server := echo.New()

    echoCtrl.NewAuthController(s.db, server).RegisterHandler()
    echoCtrl.NewUserController(s.db, server).RegisterHandler()

    server.Logger.Fatal(server.Start(":8080"))
}
