package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/vuduongtp/go-core/pkg/server/middleware/secure"
	"github.com/vuduongtp/go-logadapter"
)

// Config represents server specific config
type Config struct {
	Stage           string
	Port            int
	ReadTimeout     int
	WriteTimeout    int
	Debug           bool
	AllowOrigins    []string
	IsEnableSwagger bool
	SwaggerPath     string
}

var (
	// DefaultConfig for the API server
	DefaultConfig = Config{
		Stage:        "development",
		Port:         8080,
		ReadTimeout:  10,
		WriteTimeout: 5,
		Debug:        true,
		AllowOrigins: []string{"*"},
	}
)

func (c *Config) fillDefaults() {
	if c.Stage == "" {
		c.Stage = DefaultConfig.Stage
	}
	if c.Port == 0 {
		c.Port = DefaultConfig.Port
	}
	if c.ReadTimeout == 0 {
		c.ReadTimeout = DefaultConfig.ReadTimeout
	}
	if c.WriteTimeout == 0 {
		c.WriteTimeout = DefaultConfig.WriteTimeout
	}
	if c.AllowOrigins == nil && len(c.AllowOrigins) == 0 {
		c.AllowOrigins = DefaultConfig.AllowOrigins
	}
}

// New instantates new Echo server
func New(cfg *Config) *echo.Echo {
	cfg.fillDefaults()
	e := echo.New()
	e.Validator = NewValidator()
	e.HTTPErrorHandler = NewErrorHandler(e).Handle
	e.Binder = NewBinder()
	e.Debug = cfg.Debug
	e.Logger = logadapter.NewEchoLogger()
	e.Use(logadapter.NewEchoLoggerMiddleware())
	if e.Debug {
		e.Logger.SetLevel(log.DEBUG)
		e.Use(secure.BodyDump())
	} else {
		e.Logger.SetLevel(log.ERROR)
	}
	e.Server.Addr = fmt.Sprintf(":%d", cfg.Port)
	e.Server.ReadTimeout = time.Duration(cfg.ReadTimeout) * time.Minute
	e.Server.WriteTimeout = time.Duration(cfg.WriteTimeout) * time.Minute

	e.Use(middleware.Recover(), secure.Headers(), secure.CORS(&secure.Config{AllowOrigins: cfg.AllowOrigins}))

	return e
}

// Start starts echo server
func Start(e *echo.Echo, isDevelopment bool) {
	// hide verbose logs
	e.HideBanner = true

	// graceful shutdown
	// Start server
	go func() {
		if err := e.StartServer(e.Server); err != nil {
			if err == http.ErrServerClosed {
				logadapter.Info("shutting down the server")
			} else {
				logadapter.Error("shutting down the server", err)
			}
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		// Error from closing listeners, or context timeout:
		logadapter.Error(fmt.Sprintf("⇨ http server shutting down error: %v\n", err))
	}
}
