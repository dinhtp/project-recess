package echo

import (
    "context"
    "net/http"

    "github.com/dinhtp/project-recess/domain/user"
    "github.com/dinhtp/project-recess/util"
    "github.com/labstack/echo/v4"
    "gorm.io/gorm"
)

type UserController struct {
    server *echo.Echo
    db     *gorm.DB
}

func NewUserController(db *gorm.DB, server *echo.Echo) *UserController {
    return &UserController{db: db, server: server}
}

func (c *UserController) RegisterHandler() {
    c.server.GET("/users/:id", c.Get)
}

func (c *UserController) Get(e echo.Context) error {
    // TODO: handle validation
    userId := util.StringToInt(e.Param("id"))

    result, err := user.NewService(c.db).Get(context.Background(), uint(userId))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    return e.JSONPretty(http.StatusOK, result, "  ")
}
