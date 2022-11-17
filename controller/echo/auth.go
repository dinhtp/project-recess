package echo

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "strings"
    "time"

    "github.com/casbin/casbin/v2"
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

    claims := jwt.StandardClaims{
        ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
        Id:        fmt.Sprintf("%d", result.ID),
        Subject:   result.CasbinUser,
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err := token.SignedString([]byte(TokenKey))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    response := &message.LoginUserResponse{ID: result.ID, Email: result.Email, Token: signedToken}
    return e.JSONPretty(http.StatusOK, response, "  ")
}

type Enforcer struct {
    enforcer *casbin.Enforcer
}

func newEnforcer(e *casbin.Enforcer) *Enforcer {
    return &Enforcer{enforcer: e}
}

func (e *Enforcer) Enforce(proceed echo.HandlerFunc) echo.HandlerFunc {
    return func(ctx echo.Context) error {
        jwtToken := strings.TrimPrefix(ctx.Request().Header.Get("Authorization"), "Bearer ")
        segments := strings.Split(jwtToken, ".")

        if len(segments) != 3 {
            return echo.ErrForbidden
        }

        decodedSegment, err := jwt.DecodeSegment(segments[1])
        if err != nil {
            return echo.ErrForbidden
        }

        var payload *jwt.StandardClaims
        err = json.Unmarshal(decodedSegment, &payload)
        if err != nil || payload == nil {
            return echo.ErrForbidden
        }

        passed, err := e.enforcer.Enforce(payload.Subject, ctx.Request().URL.Path, ctx.Request().Method)
        if passed && err == nil {
            return proceed(ctx)
        }

        return echo.ErrForbidden
    }
}
