package main

import "net/http"
import "Hug/framework"

func main() {
	server := &http.Server{
		Handler: framework.NewEngine(),
		Addr:    ":9393",
	}
	server.ListenAndServe()
}
