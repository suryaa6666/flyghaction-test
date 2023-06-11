package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "SURYA GANTENG!",
		})
	})

	PORT := os.Getenv("PORT")

	e.Logger.Fatal(e.Start(":" + PORT))
}
