package main

import "github.com/labstack/echo"

func setStaticRoute(e *echo.Echo)  {
	e.Static("/static", "public")
}