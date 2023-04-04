package main

import (
	"periodic-timestamps/app"
	"periodic-timestamps/settings"
)

func main() {

	app.RunWithEngine(settings.ENGINE)

}
