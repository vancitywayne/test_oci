package main

import (
    "log"
    "net/http"
    "proyek1-be/routes"
)

func main() {
    r := routes.RegisterRoutes()
    log.Println("Server started on: http://localhost:3714")
    log.Fatal(http.ListenAndServe(":3714", r))
    
}

func Handler(w http.ResponseWriter, r *http.Request) {
	routes.RegisterRoutes().ServeHTTP(w, r)
}
