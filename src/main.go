/*
Very simple static file http server in go
Usage:
	-p="8080": listening port, default port 8080
	-d=".": static files directory, default current directory
*/

package main

import (
	"log"
        "flag"
        "net/http"
)

func main() {

	port := flag.String("p", "8080", "listening port")
	directory := flag.String("d", "/wwwroot", "static files directory")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
