package mux

import (
    "fmt"
    "net/http"

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
    fmt.Println(r.Method, r.RequestURI)
}

func (c *UserController) List(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Method, r.RequestURI)
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Method, r.RequestURI)
}

func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Method, r.RequestURI)

}

func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Method, r.RequestURI)
}
