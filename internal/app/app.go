package app

import (
	handlers "Effective_Mobile/internal/handlers_api"
	"Effective_Mobile/internal/services"
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

func New(log *slog.Logger, port int, userService *services.UserService, worklogService *services.WorklogService) *App {
	app := &App{
		Log: log,
	}
	app.Router = mux.NewRouter()
	handlers.SetupRoutes(app.Router, userService, worklogService)
	app.HTTPServer = &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: app.Router,
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
