package server

import (
    echoCtrl "github.com/dinhtp/project-recess/controller/echo"
    _ "github.com/dinhtp/project-recess/docs"
    "github.com/labstack/echo/v4"
    echoSwagger "github.com/swaggo/echo-swagger"
    "gorm.io/gorm"
)

type EchoServer struct {
    db      *gorm.DB
    Address string
}

func (s *EchoServer) Serve() {
    server := echo.New()

    server.GET("/swagger/*", echoSwagger.WrapHandler)

    echoCtrl.NewAuthController(s.db, server).RegisterHandler()
    echoCtrl.NewUserController(s.db, server).RegisterHandler()

    server.Logger.Fatal(server.Start(s.Address))
}
