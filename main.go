package main

import (
	"cariinkos/bootstrap"
)

func main() {
	app := bootstrap.Boot()

	app.Start()
}
