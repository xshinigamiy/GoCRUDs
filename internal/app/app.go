package app

import (
	"GoPractice/GoCRUDs/config"
	h "GoPractice/GoCRUDs/internal/health-check"
	"GoPractice/GoCRUDs/pkg/utils"
	"context"
	"errors"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	Router *mux.Router
}

func (a *App) SetupService() {
	log.Info("Server initialization started...")
	config.InitConfigs()
	a.Router = a.InitializeRoutes()
	log.Info("Server initialization successful")
}

func (a *App) InitializeRoutes() *mux.Router {
	r := mux.NewRouter()
	log.Info("Initializing routes")
	h.HealthCheckRoutes(r)
	return r
}

func (a *App) StartHttpServer() {
	port := ":" + config.GetConfigs().ServiceConfig.AppPort
	log.Info("port: " + port)
	server := &http.Server{
		Addr:         port,
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  10 * time.Second,
		Handler:      a.Router,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
		log.Info("Server started successfully")
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	_ = server.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	//log.Info("shutting down server....")
	os.Exit(0)
}

func reportPanics(w http.ResponseWriter) {
	// Capture the panic if there is one, and report it.
	if penic := recover(); penic != nil {
		var err error
		switch x := penic.(type) {
		case string:
			err = errors.New(x)
		case error:
			err = x
		default:
			err = errors.New("unknown panic")
		}
		utils.RespondWithJSON(w, http.StatusInternalServerError, nil, err.Error(), "data")
	}
}

/*
Introducing a panic reporting middleware. The middleware handler would defer reportPanics() before calling
the next handler
*/
func panicReporterFunc(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer reportPanics(w)
	next(w, r)
}
