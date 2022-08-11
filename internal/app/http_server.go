package app

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/borisbbtest/ya-dr/internal/config"
	"github.com/borisbbtest/ya-dr/internal/handlers"
	midd "github.com/borisbbtest/ya-dr/internal/middleware"
	"github.com/borisbbtest/ya-dr/internal/model"
	"github.com/borisbbtest/ya-dr/internal/storage"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("context", "system_loyalty")

type serviceSystemLoyalty struct {
	wrapp  handlers.WrapperHandler
	middle midd.WrapperMiddleware
}

func New(cfg *config.MainConfig) *serviceSystemLoyalty {
	sessionHTTP := &storage.SessionHTTP{DBSession: map[string]model.DataUser{}}
	return &serviceSystemLoyalty{
		wrapp: handlers.WrapperHandler{
			ServerConf: cfg,
			Session:    sessionHTTP,
		},
		middle: midd.WrapperMiddleware{
			Session: sessionHTTP,
		},
	}
}

func (hook *serviceSystemLoyalty) Start() (err error) {

	// Launch the listening thread
	log.Println("Initializing HTTP server")
	r := chi.NewRouter()

	hook.wrapp.Storage, err = storage.NewPostgreSQLStorage(hook.wrapp.ServerConf.DatabaseURI)
	if err != nil {
		log.Error(err)

	}
	defer hook.wrapp.Storage.Close()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(midd.GzipHandle)
	//r.Use(middleware.Compress(5, "gzip"))
	r.Use(middleware.Recoverer)

	r.Post("/api/user/register", hook.wrapp.PostJSONRegisterHandler)
	r.Post("/api/user/login", hook.wrapp.PostJSONLoginHandler)

	serviceLogic := r.Group(nil)
	serviceLogic.Use(hook.middle.MiddleSetSessionCookie)
	//serviceLogic.Use(midd.GzipHandle)
	//serviceLogic.Use(middleware.Compress(5, "gzip"))

	serviceLogic.Post("/api/user/orders", hook.wrapp.PostOrderHandler)
	serviceLogic.Get("/api/user/orders", hook.wrapp.GetJSONOrdersHandler)

	serviceLogic.Get("/api/user/balance", hook.wrapp.GetJSONOrdersHandler)

	serviceLogic.Post("/api/user/balance/withdraw", hook.wrapp.GetJSONOrdersHandler)

	serviceLogic.Post("/api/user/withdrawals", hook.wrapp.GetJSONOrdersHandler)

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "web"))
	hook.wrapp.FileServer(r, "/form", filesDir)

	server := &http.Server{
		Addr:         hook.wrapp.ServerConf.RunAddress,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("can't start the listening thread: %s", err)
	}

	log.Info("Exiting")
	return nil
}
