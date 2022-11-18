package echo

import (
    "encoding/json"
    "strings"

    "github.com/casbin/casbin/v2"
    "github.com/golang-jwt/jwt"
    "github.com/labstack/echo/v4"
)

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
