package app

import (
	"Effective_Mobile/internal/api"
	"context"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type App struct {
	Log        *slog.Logger
	HTTPServer *http.Server
	Router     *mux.Router
}

func New(log *slog.Logger, port int) *App {
	app := &App{
		Log: log,
	}

	app.Router = mux.NewRouter()
	api.SetupRoutes(app.Router)
	app.HTTPServer = &http.Server{
		Addr:         ":" + strconv.Itoa(port),
		Handler:      app.Router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return app
}

func (a *App) MustRun() {
	a.Log.Info("starting HTTP server", slog.String("addr", a.HTTPServer.Addr))
	if err := a.HTTPServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		a.Log.Error("failed to start HTTP server", slog.String("error", err.Error()))
		panic(err)
	}
}

func (a *App) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.HTTPServer.Shutdown(ctx); err != nil {
		a.Log.Error("failed to shutdown HTTP server", slog.String("error", err.Error()))
	} else {
		a.Log.Info("HTTP server stopped gracefully")
	}
}
