package echo

import (
    "context"
    "net/http"
    "time"

    "github.com/dinhtp/project-recess/domain/message"
    "github.com/dinhtp/project-recess/domain/user"
    "github.com/dinhtp/project-recess/util"
    "github.com/golang-jwt/jwt"
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
    group := c.server.Group("/users")

    group.GET("/:id", c.Get)
    group.GET("", c.List)
    group.POST("", c.Create)
    group.PUT("/:id", c.Update)
    group.DELETE("/:id", c.Delete)
    group.POST("/login", c.Login)
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

func (c *UserController) Login(e echo.Context) error {
    request := new(message.LoginUserRequest)
    if err := e.Bind(request); err != nil {
        return echo.NewHTTPError(http.StatusUnauthorized, "invalid email or password")
    }

    result, err := user.NewService(c.db).Login(context.Background(), request)
    if err != nil {
        return echo.NewHTTPError(http.StatusUnauthorized, "invalid email or password")
    }

    claims := jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 72).Unix()}
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    t, err := token.SignedString([]byte("secret"))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    response := &message.LoginUserResponse{
        ID:    result.ID,
        Email: result.Email,
        Token: t,
    }

    return e.JSONPretty(http.StatusOK, response, "  ")
}
