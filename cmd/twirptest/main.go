package main

import (
	"net/http"
	server "twirptest/internal/twirptestserver"
	pb "twirptest/rpc/twirptest"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main(){
    server := &server.HelloWorldServer{}
	twirpHandler := pb.NewHelloWorldServer(server)
    // You can use any mux you like - NewHelloWorldServer gives you an http.Handler.
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(func(h http.Handler) http.Handler {
        fn := func(w http.ResponseWriter, r *http.Request) {
            r.Header.Set("Access-Control-Allow-Origin", "*")
            
            h.ServeHTTP(w, r)
        }
        return http.HandlerFunc(fn)
    })
    r.Mount(twirpHandler.PathPrefix(), twirpHandler)
    

  
    http.ListenAndServe(":8081", r)
}