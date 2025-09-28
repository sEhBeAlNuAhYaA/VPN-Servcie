package app

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	diContainer *di
	server      *http.Server
}

func New() *App {
	a := &App{}
	a.initDepend()

	return a
}

func (a *App) initDepend() error {
	errDi := a.initDI()
	errHttp := a.initServer()

	return errors.Join(errDi, errHttp)
}

func (a *App) initDI() error {
	a.diContainer = NewDI()
	fmt.Println("DI init")
	return nil
}

func (a *App) initServer() error {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)

	router.Post("/api/popa", a.diContainer.GetApi().Post)
	router.Get("/api/popa", a.diContainer.GetApi().Get)
	router.Delete("/api/popa", a.diContainer.GetApi().Delete)
	router.Patch("/api/popa", a.diContainer.GetApi().Update)

	a.server = &http.Server{
		Addr:              "localhost:8000",
		ReadHeaderTimeout: time.Duration(time.Duration.Seconds(1)),
		Handler:           router,
	}
	fmt.Println("Server init")
	return nil
}

func (a *App) Run() error {
	go func() {
		err := a.runHttpServer()
		if err != nil {
			fmt.Println("пизда нахуй")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("GG")
	return nil
}

func (a *App) runHttpServer() error {
	err := a.server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
