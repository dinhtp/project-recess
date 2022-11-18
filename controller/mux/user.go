package mux

import (
    "context"
    "encoding/json"
    "net/http"

    "github.com/dinhtp/project-recess/domain/message"
    "github.com/dinhtp/project-recess/domain/user"
    "github.com/dinhtp/project-recess/util"
    "github.com/gorilla/mux"
    "gorm.io/gorm"
)

type UserController struct {
    router *mux.Router
    db     *gorm.DB
}

func NewUserController(db *gorm.DB, router *mux.Router) *UserController {
    return &UserController{db: db, router: router}
}

func (c *UserController) RegisterHandler() {
    c.router.HandleFunc("/users", c.List).Methods(http.MethodGet)
    c.router.HandleFunc("/users/{id:[0-9]+}", c.Get).Methods(http.MethodGet)
    c.router.HandleFunc("/users", c.Create).Methods(http.MethodPost)
    c.router.HandleFunc("/users/{id:[0-9]+}", c.Update).Methods(http.MethodPut)
    c.router.HandleFunc("/users/{id:[0-9]+}", c.Delete).Methods(http.MethodDelete)

    //c.registerMiddleware(group)
}

//func (c *UserController) registerMiddleware(group *echo.Group) {
//    jwtConfig := middleware.JWTConfig{Claims: &jwt.StandardClaims{}, SigningKey: []byte(TokenKey)}
//    group.Use(middleware.JWTWithConfig(jwtConfig))
//
//    enforcer, err := casbin.NewEnforcer("model.conf", "policy.csv")
//    if err != nil {
//        return
//    }
//
//    group.Use(newEnforcer(enforcer).Enforce)
//}

func (c *UserController) Get(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    userId := util.StringToInt(params["id"])

    result, err := user.NewService(c.db).Get(context.Background(), uint(userId))
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(result)
}

func (c *UserController) List(w http.ResponseWriter, r *http.Request) {
    request := &message.ListUserRequest{
        Page:    uint(util.StringToInt(r.URL.Query().Get("page"))),
        PerPage: uint(util.StringToInt(r.URL.Query().Get("per_page"))),
    }

    result, err := user.NewService(c.db).List(context.Background(), request)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(result)
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
    request := new(message.User)

    err := json.NewDecoder(r.Body).Decode(&request)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    result, err := user.NewService(c.db).Create(context.Background(), request)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(result)
}

func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
    request := new(message.User)

    err := json.NewDecoder(r.Body).Decode(&request)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    result, err := user.NewService(c.db).Update(context.Background(), request)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(result)
}

func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    userId := util.StringToInt(params["id"])

    err := user.NewService(c.db).Delete(context.Background(), uint(userId))
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
}
