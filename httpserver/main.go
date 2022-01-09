package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", HttpHandler)
	http.HandleFunc("/healthz", HealthZ)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func HealthZ(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(200)
}
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	ip := ClientIP(r)
	httpcode := 200
	log.Printf("reqLog: clientIP : %+v, httpcode : %+v", ip, httpcode)
	headers := r.Header
	log.Printf("req headers : %+v", headers)
	for k,v := range headers {
		for _,val := range v {
			w.Header().Set(k,val)
		}
	}
	version := os.Getenv("VERSION")
	if version != "" {
		w.Header().Set("VERSION", version)
	}
	w.WriteHeader(httpcode)
	io.WriteString(w, "ok")
}

func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}
