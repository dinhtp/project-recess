package mux

import (
    "context"
    "encoding/json"
    "net/http"

    "github.com/dinhtp/project-recess/domain/message"
    "github.com/dinhtp/project-recess/domain/user"
    "github.com/gorilla/mux"
    "gorm.io/gorm"
)

type AuthController struct {
    router *mux.Router
    db     *gorm.DB
}

func NewAuthController(db *gorm.DB, router *mux.Router) *AuthController {
    return &AuthController{db: db, router: router}
}

func (c *AuthController) RegisterHandler() {
    c.router.HandleFunc("/login", c.Login).Methods(http.MethodPost)
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
    request := new(message.LoginUserRequest)
    err := json.NewDecoder(r.Body).Decode(&request)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    result, err := user.NewService(c.db).Login(context.Background(), request)
    if err != nil {
        http.Error(w, err.Error(), http.StatusForbidden)
        return
    }

    _ = json.NewEncoder(w).Encode(result)
}
