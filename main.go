package main

import (
	"log"
	"math/rand"
	"net/http"
	"path"
	"time"

	"github.com/go-echarts/go-echarts/charts"
)

const (
	host = "10.10.30.220:8080"
    maxNum = 100
)

type router struct {
	name string
	charts.RouterOpts
}

func getRenderPath(f string) string {
	return path.Join("html", f)
}

func logTracing(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Tracing request for %s\n", r.RequestURI)
		next.ServeHTTP(w, r)
	}
}

var seed = rand.NewSource(time.Now().UnixNano())

func main() {
	// Avoid "404 page not found".
	http.HandleFunc("/unixbench", logTracing(unixbenchHandler))
	http.HandleFunc("/speccpu2006", logTracing(speccpu2006Handler))
	log.Println("Run server at " + host)
	http.ListenAndServe(host, nil)
}
