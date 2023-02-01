package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/pflag"
)

const html = `<html>
<head>
	<meta charset="UTF-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1.0" />
	<script src="https://cdn.tailwindcss.com"></script>
</head>
<body>
	<div class="h-screen flex items-center justify-center">
		<div class="text-6xl">Hello, world!</div>
	</div>
</body>
</html>`

// flags
var (
	port = pflag.Uint32("port", 80, "Port to bind web server")
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, html)
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", *port)))
}
