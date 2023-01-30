package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"postgre-dashboard/postgre"
	"time"
)

var postgreHost = flag.String("postgre-host", "localhost", "postgre host")
var postgrePort = flag.Int("postgre-port", 5432, "postgre port")
var postgreUsername = flag.String("postgre-username", "localhost", "postgre username")
var postgrePassword = flag.String("postgre-password", "", "postgre password")

func main() {
	r := mux.NewRouter()
	r.Use(accessControlMiddleware)
	handler, err := postgre.NewHandler(*postgreHost, *postgrePort, *postgreUsername, *postgrePassword)
	if err != nil {
		panic(err)
	}
	handler.Handle(r.PathPrefix("/api/etcd").Subrouter())
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:10001",
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	err = srv.ListenAndServe()
	if err != nil {
		logrus.Errorf("web server start failed %v", err)
		panic(err)
	}
}

func accessControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS,PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}
