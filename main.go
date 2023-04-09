package main

import (
	"flag"
	"os"
	"periodic-timestamps/app"
	"periodic-timestamps/settings"
)

type flagInputs struct {
	host string
	port string
}

func (f flagInputs) flagsEnvVarsMap() map[string]string {
	return map[string]string{
		settings.ApiHostEnvVarName: f.host,
		settings.ApiPortEnvVarName: f.port,
	}
}

func getFlags() flagInputs {
	host := flag.String("host", "0.0.0.0", "Host address where the service will run. Default: 0.0.0.0")
	port := flag.String("port", "8080", "Port number on which the service will be exposed. Default: 8080")
	flag.Parse()
	return flagInputs{
		host: *host,
		port: *port,
	}
}

func setEnvVarsIfNotSet(envVarsMap map[string]string) {
	for name, value := range envVarsMap {
		if os.Getenv(name) == "" {
			os.Setenv(name, value)
		}
	}
}

func main() {
	flags := getFlags()
	setEnvVarsIfNotSet(flags.flagsEnvVarsMap())
	app.RunWithEngine(settings.ENGINE)
}
