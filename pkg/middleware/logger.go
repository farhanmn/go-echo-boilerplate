package middleware

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"time"
)

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		end := time.Now()

		date := time.Now().Format("2006/01/02 15:04:05")
		method := c.Request().Method
		uri := c.Request().URL.Path
		status := c.Response().Status
		duration := end.Sub(start)

		var statusColor func(a ...interface{}) string
		switch {
		case status >= 500:
			statusColor = color.New(color.FgRed).SprintFunc() // Error
		case status >= 400:
			statusColor = color.New(color.FgYellow).SprintFunc() // Warning
		case status >= 300:
			statusColor = color.New(color.FgBlue).SprintFunc() // Redirect
		default:
			statusColor = color.New(color.FgGreen).SprintFunc() // Success
		}

		logMessage := fmt.Sprintf(
			"[%s] %s %s → %s | %s ms",
			date,
			statusColor(method),
			statusColor(uri),
			statusColor(status),
			color.YellowString(fmt.Sprintf("%v", duration.Milliseconds())),
		)

		logrus.Info(logMessage)

		return err
	}
}
