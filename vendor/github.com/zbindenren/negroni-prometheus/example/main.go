package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/zbindenren/negroni-prometheus"
)

func main() {
	n := negroni.New()
	m := negroniprometheus.NewMiddleware("serviceName")
	// if you want to use other buckets than the default (300, 1200, 5000) you can run:
	// m := negroniprometheus.NewMiddleware("serviceName", 400, 1600, 700)

	n.Use(m)

	r := http.NewServeMux()
	r.Handle("/metrics", prometheus.Handler())
	r.HandleFunc(`/ok`, func(w http.ResponseWriter, r *http.Request) {
		sleep := rand.Intn(4999) + 1
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "slept %d milliseconds\n", sleep)
	})
	r.HandleFunc(`/notfound`, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "not found")
	})

	n.UseHandler(r)

	n.Run(":3000")
}
