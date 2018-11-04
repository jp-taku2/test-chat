package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func setRoute(e *echo.Echo)  {
	e.GET("/", handleIndexGet)

	// auth
	e.GET("/login", handleLoginGet)
	e.POST("/login", handleLoginPost)
	e.GET("/signup", handleSignupGet)
	e.POST("/signup", handleSignupPost)
	e.GET("/signup_account", handleSignupAccountGet)
	e.POST("/signup_account", handleSignupAccountPost)



}

func handleIndexGet(c echo.Context) error {
	return c.String(http.StatusOK, "index")
}

func handleLoginGet(c echo.Context) error {
	return c.Render(http.StatusOK, "login", nil)
}

func handleLoginPost(c echo.Context) error {
	return c.String(http.StatusOK, "handleLoginPost")
}

func handleSignupGet(c echo.Context) error {
	return c.String(http.StatusOK, "handleSignupGet")
}

func handleSignupPost(c echo.Context) error {
	return c.String(http.StatusOK, "handleSignupPost")
}

func handleSignupAccountGet(c echo.Context) error {
	return c.String(http.StatusOK, "handleSignupAccountGet")
}

func handleSignupAccountPost(c echo.Context) error {
	return c.String(http.StatusOK, "handleSignupAccountPost")
}