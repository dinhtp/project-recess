package echo

import (
    "context"
    "net/http"

    "github.com/casbin/casbin/v2"
    "github.com/dinhtp/project-recess/domain/message"
    "github.com/dinhtp/project-recess/domain/user"
    "github.com/dinhtp/project-recess/util"
    "github.com/golang-jwt/jwt"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
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
    group := c.server.Group(user.PathPrefix)

    c.registerMiddleware(group)

    group.GET("/:id", c.Get)
    group.GET("", c.List)
    group.POST("", c.Create)
    group.PUT("/:id", c.Update)
    group.DELETE("/:id", c.Delete)
}

func (c *UserController) registerMiddleware(group *echo.Group) {
    jwtConfig := middleware.JWTConfig{Claims: &jwt.StandardClaims{}, SigningKey: []byte(TokenKey)}
    group.Use(middleware.JWTWithConfig(jwtConfig))

    enforcer, err := casbin.NewEnforcer("model.conf", "policy.csv")
    if err != nil {
        return
    }

    group.Use(newEnforcer(enforcer).Enforce)
}

func (c *UserController) Get(e echo.Context) error {
    userId := util.StringToInt(e.Param("id"))

    result, err := user.NewService(c.db).Get(context.Background(), uint(userId))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    return e.JSONPretty(http.StatusOK, result, "  ")
}

func (c *UserController) List(e echo.Context) error {
    request := &message.ListUserRequest{
        Page:    uint(util.StringToInt(e.QueryParam("page"))),
        PerPage: uint(util.StringToInt(e.QueryParam("per_page"))),
    }

    result, err := user.NewService(c.db).List(context.Background(), request)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    return e.JSONPretty(http.StatusOK, result, "  ")
}

func (c *UserController) Create(e echo.Context) error {
    request := new(message.User)
    if err := e.Bind(request); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    result, err := user.NewService(c.db).Create(context.Background(), request)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    return e.JSONPretty(http.StatusCreated, result, "  ")
}

func (c *UserController) Update(e echo.Context) error {
    request := new(message.User)
    if err := e.Bind(request); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    result, err := user.NewService(c.db).Update(context.Background(), request)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    return e.JSONPretty(http.StatusOK, result, "  ")
}

func (c *UserController) Delete(e echo.Context) error {
    userId := util.StringToInt(e.Param("id"))

    err := user.NewService(c.db).Delete(context.Background(), uint(userId))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    return e.NoContent(http.StatusNoContent)
}
