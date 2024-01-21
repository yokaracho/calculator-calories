package app

import (
	"blocker/handler"
	hellopage "blocker/templates/hello"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func HTML(c echo.Context, comp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return comp.Render(c.Request().Context(), c.Response().Writer)
}

func App() {
	e := echo.New()
	e.Static("/static", "static")

	e.GET("/", func(c echo.Context) error {
		return HTML(c, hellopage.HelloPage())
	})
	e.POST("/hello", handler.CalculateHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
