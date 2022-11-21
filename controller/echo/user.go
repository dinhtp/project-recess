package echo

import (
    "context"
    "net/http"

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
    jwtConfig := middleware.JWTConfig{Claims: &jwt.StandardClaims{}, SigningKey: []byte(user.TokenKey)}
    group.Use(middleware.JWTWithConfig(jwtConfig))

    //enforcer, err := casbin.NewEnforcer("/var/rbac/model.conf", "/var/rbac/policy.csv")
    //if err != nil {
    //    return
    //}
    //
    //group.Use(newEnforcer(enforcer).Enforce)
}

// Get godoc
// @Summary      Get a user detail
// @Description  Get a user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  message.User
// @Failure      500  {object}  interface{} "{"error":"error_code", "message":"error_description"}"
// @Router       /users/{id} [get]
func (c *UserController) Get(e echo.Context) error {
    userId := util.StringToInt(e.Param("id"))

    result, err := user.NewService(c.db).Get(context.Background(), uint(userId))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    return e.JSONPretty(http.StatusOK, result, "  ")
}

// List godoc
// @Summary      Get user list
// @Description  Get user list by page size and limit
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        page       query		string	false	"current page"
// @Param        per_page   query		string	false	"page limit"
// @Success      200  {object}  message.ListUserResponse
// @Failure      500  {object}  interface{} "{"error":"error_code", "message":"error_description"}"
// @Router       /users [get]
func (c *UserController) List(e echo.Context) error {
    request := &message.ListUserRequest{
        Page:    uint(util.StringToInt(e.QueryParam("page"))),
        PerPage: uint(util.StringToInt(e.QueryParam("per_page"))),
    }

    if request.Page == 0 {
        request.Page = 1
    }

    if request.PerPage == 0 {
        request.PerPage = 10
    }

    result, err := user.NewService(c.db).List(context.Background(), request)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    return e.JSONPretty(http.StatusOK, result, "  ")
}

// Create godoc
//
//	@Summary		Create a user
//	@Description	Create user by json
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		message.User	true	"Create Account"
//	@Success		204		{object}	message.User
//  @Failure        500     {object}    interface{}  "{"error":"error_code", "message":"error_description"}"
//	@Router			/users [post]
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

// Update godoc
//
//	@Summary		Update a user
//	@Description	Update user by json
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		message.User	true	"Update User"
//  @Param          id      path        int             true    "User ID"
//	@Success		200		{object}	message.User
//  @Failure        500     {object}    interface{} "{"error":"error_code", "message":"error_description"}"
//	@Router			/users/{id} [put]
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

// Delete godoc
//
//	@Summary		Delete a user
//	@Description	Delete by user ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"user ID"
//	@Success		200	{object}   interface{} "{}"
//  @Failure        500  {object}  interface{} "{"error":"error_code", "message":"error_description"}"
//	@Router			/users/{id} [delete]
func (c *UserController) Delete(e echo.Context) error {
    userId := util.StringToInt(e.Param("id"))

    err := user.NewService(c.db).Delete(context.Background(), uint(userId))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    return e.NoContent(http.StatusNoContent)
}
