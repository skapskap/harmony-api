package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

func main()  {

	listenAddr := flag.String("listenAddr", ":4869", "server listen address")

	srv := &http.Server{
		Addr:         *listenAddr,
		Handler:      routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Printf("Starting server in port %s...\n", *listenAddr)

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}