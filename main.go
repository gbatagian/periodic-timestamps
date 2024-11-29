package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"periodic-timestamps/ptlist"
	"periodic-timestamps/settings"
)

type flagInputs struct {
	host *string
	port *string
}

func getFlags() flagInputs {
	host := flag.String("host", settings.DefaultHost, fmt.Sprintf("Server host address, default: %s", settings.DefaultHost))
	port := flag.String("port", settings.DefaultPort, fmt.Sprintf("Server listening port, default: %s", settings.DefaultPort))
	flag.Parse()
	return flagInputs{
		host: host,
		port: port,
	}
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	router.HandleFunc("GET /ptlist", ptlist.PtListGet())

	flags := getFlags()
	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", *flags.host, *flags.port),
		Handler: router,
	}

	log.Printf("Server starts listening at: %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
