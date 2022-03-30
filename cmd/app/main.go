package main

import (
	"context"
	"encoding/gob"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/Elren44/elog"
	"github.com/Elren44/elren_todo/internal/composites"
	"github.com/Elren44/elren_todo/internal/config"
	"github.com/Elren44/elren_todo/internal/model"
	"github.com/Elren44/elren_todo/pkg/client/postgres"
	"github.com/Elren44/elren_todo/pkg/middleware"
	"github.com/Elren44/elren_todo/pkg/sessions"
	"github.com/Elren44/elren_todo/pkg/shutdown"
	"github.com/Elren44/elren_todo/pkg/utils"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var (
	cfg    *config.Config
	router *mux.Router
	logger *zap.SugaredLogger
)

func main() {
	var cfgPath string
	flag.StringVar(&cfgPath, "cfg_path", "./config.yml", "set config file path")
	flag.StringVar(&cfgPath, "cfg", "./config.yml", "set config file path shorthand")
	flag.Parse()

	os.Setenv("CONF_PATH", cfgPath)

	err := run()
	if err != nil {
		log.Fatal(err)
	}

	startServ(router, logger, cfg)
}

func startServ(router *mux.Router, logger *zap.SugaredLogger, cfg *config.Config) {
	var server *http.Server
	var listener net.Listener

	logger.Infof("bind application to host: %s and port: %s", cfg.Listen.BindIP, cfg.Listen.Port)

	var err error

	listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	if err != nil {
		logger.Fatal(err)
	}

	server = &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go shutdown.Graceful([]os.Signal{syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM},
		server)

	logger.Info("application initialized and started")

	if err := server.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			logger.Warn("server shutdown")
		default:
			logger.Fatal(err)
		}
	}
}

func run() error {
	gob.Register(model.UserDTO{})

	session := sessions.NewSessions()
	//send session to middleware
	middleware.SendSessions(session)
	logger = elog.InitLogger(elog.JsonOutput)
	cfg = config.GetConfig()

	tc, err := utils.TemplateCache()
	if err != nil {
		logger.Fatal("failed to create cache")
		return err
	}

	cfg.UseCache = false
	cfg.Session = session
	cfg.TemplatesCache = tc

	//send cfg to render pkg
	utils.NewAppTemplates(cfg)

	logger.Debug("creating db client")
	client, err := postgres.NewClient(context.TODO(), 3, cfg.Storage)
	if err != nil {
		logger.Fatalf("failed to get client: %v", err)
		return err
	}

	taskComposite := composites.NewTaskComposite(client, logger, cfg)
	userComposite := composites.NewUserComposite(client, logger, cfg)

	router = mux.NewRouter()
	router.Use(middleware.NoSurf)
	router.Use(middleware.SessionLoad)

	taskComposite.Handler.Register(router)
	userComposite.Handler.Register(router)

	return nil
}
