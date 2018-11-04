package main

import (
	"./setting"
	"context"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"html/template"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var templates map[string]*template.Template

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	t := &Template{}
	e.Renderer = t

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	setRoute(e)
	setStaticRoute(e)

	go func() {
		if err := e.Start(setting.Server.Port); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Info(err)
		e.Close()
	}
	time.Sleep(1 * time.Second)
}

func init() {
	setting.Load()
	loadTemplates()
}