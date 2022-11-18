package echo

import (
    "context"
    "fmt"
    "net/http"
    "time"

    "github.com/dinhtp/project-recess/domain/message"
    "github.com/dinhtp/project-recess/domain/user"
    "github.com/golang-jwt/jwt"
    "github.com/labstack/echo/v4"
    "gorm.io/gorm"
)

const (
    TokenKey = "secret"
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

    // create jwt standard claims
    claims := jwt.StandardClaims{
        ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
        Id:        fmt.Sprintf("%d", result.ID),
        Subject:   result.CasbinUser,
    }

    // create jwt token and sign the token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err := token.SignedString([]byte(TokenKey))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    response := &message.LoginUserResponse{ID: result.ID, Email: result.Email, Token: signedToken}
    return e.JSONPretty(http.StatusOK, response, "  ")
}
