package server

import (
    "net/http"
    "time"

    muxCtrl "github.com/dinhtp/project-recess/controller/mux"
    "github.com/gorilla/mux"
    "github.com/sirupsen/logrus"
    "gorm.io/gorm"
)

type MuxServer struct {
    db      *gorm.DB
    Address string
}

func (s *MuxServer) Serve() {
    router := mux.NewRouter()

    muxCtrl.NewUserController(s.db, router).RegisterHandler()
    muxCtrl.NewAuthController(s.db, router).RegisterHandler()

    server := &http.Server{
        Handler:      router,
        Addr:         s.Address,
        WriteTimeout: 30 * time.Second,
        ReadTimeout:  60 * time.Second,
    }

    logrus.Fatal(server.ListenAndServe())
}
