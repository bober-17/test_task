package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", Handler, MiddleWare)

	e.Logger.Fatal(e.Start(":1323"))
}

func Handler(ctx echo.Context) error {

	var (
		Year       = 2027
		Month      = time.January
		Day        = 1
		Hour       = 0
		Minute     = 0
		Second     = 0
		Nanosecond = 0
		Location   = time.UTC
	)

	futureDate := time.Date(Year, Month, Day, Hour, Minute, Second, Nanosecond, Location)
	now := time.Now()
	diff := futureDate.Sub(now)
	days := int(diff.Hours() / 24)

	s := fmt.Sprintf("Количество дней до %d.%d.%d года: %d", Day, Month, Year, days)

	return ctx.String(http.StatusOK, s)
}

func MiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		userRole := ctx.Request().Header.Get("User-Role")

		if userRole == "admin" {
			log.Println("red button user detected")
		}

		return next(ctx)
	}
}
