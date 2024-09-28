package api

import "github.com/labstack/echo/v4"

func HandleFresh(c echo.Context) error {
	return c.HTML(200, "bruh")
}
