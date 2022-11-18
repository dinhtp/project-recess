package echo

import (
    "context"
    "github.com/dinhtp/project-recess/domain/message"
    "github.com/dinhtp/project-recess/domain/user"
    "github.com/labstack/echo/v4"
    "gorm.io/gorm"
    "net/http"
)

type AuthController struct {
    server *echo.Echo
    db     *gorm.DB
}

func NewAuthController(db *gorm.DB, server *echo.Echo) *AuthController {
    return &AuthController{db: db, server: server}
}

func (c *AuthController) RegisterHandler() {
    c.server.POST("/login", c.Login)
}

func (c *AuthController) Login(e echo.Context) error {
    request := new(message.LoginUserRequest)
    if err := e.Bind(request); err != nil {
        return echo.NewHTTPError(http.StatusUnauthorized, "invalid email or password")
    }

    result, err := user.NewService(c.db).Login(context.Background(), request)
    if err != nil {
        return echo.NewHTTPError(http.StatusUnauthorized, "invalid email or password")
    }

    return e.JSONPretty(http.StatusOK, result, "  ")
}
