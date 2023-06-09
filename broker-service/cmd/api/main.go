package main

import (
	"fmt"
	"log"
	"net/http"
)

// webPort is the port for the web server
const webPort = "80"

type Config struct{}

// starting broker service
func main() {
	app := Config{}

	log.Printf("Starting broker service on port %s\n", webPort)

	//define server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes()}

	// start server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
