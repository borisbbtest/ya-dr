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
	"github.com/borisbbtest/ya-dr/internal/storage"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("context", "system_loyalty")

type serviceShortURL struct {
	wrapp handlers.WrapperHandler
}

func New(cfg *config.MainConfig) *serviceShortURL {
	return &serviceShortURL{
		wrapp: handlers.WrapperHandler{
			ServerConf: cfg,
		},
	}
}

func (hook *serviceShortURL) Start() (err error) {

	// Launch the listening thread
	log.Println("Initializing HTTP server")
	r := chi.NewRouter()

	hook.wrapp.Storage, err = storage.NewPostgreSQLStorage(hook.wrapp.ServerConf.DATABASE_URI)
	if err != nil {
		log.Error(err)

	}
	defer hook.wrapp.Storage.Close()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(midd.GzipHandle)
	r.Use(midd.MidSetCookie)
	//r.Use(middleware.Compress(5, "gzip"))
	r.Use(middleware.Recoverer)

	r.Post("/api/user/register", hook.wrapp.PostJSONRegisterHandler)
	r.Post("/api/user/login", hook.wrapp.PostJSONLoginHandler)

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "web"))
	hook.wrapp.FileServer(r, "/form", filesDir)

	server := &http.Server{
		Addr:         hook.wrapp.ServerConf.RUN_ADDRESS,
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
