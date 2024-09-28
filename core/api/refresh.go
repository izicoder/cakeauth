package api

import "github.com/labstack/echo/v4"

func HandleRefresh(c echo.Context) error {
	return c.HTML(200, "wtf")
}
